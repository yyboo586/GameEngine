package logics

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"GameEngine/internal/dao"
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// 逻辑层
var (
	logicsAsyncTaskOnce     sync.Once
	logicsAsyncTaskInstance *logicsAsyncTask
)

type logicsAsyncTask struct {
	dbPool *sql.DB
	logger *glog.Logger
	ctx    context.Context

	initInterval     time.Duration   // 工作线程初始化间隔
	queryInterval    time.Duration   // 工作线程没有任务时，休眠间隔
	errSleepInterval time.Duration   // 工作线程获取任务失败，休眠间隔
	backoffIntervals []time.Duration // 工作线程执行任务失败，退避间隔

	sigChanMap map[model.AsyncTaskType]chan struct{} // 任务线程唤醒信号通道

	handler map[model.AsyncTaskType]model.AsyncTaskHandler
	mutex   sync.RWMutex
}

func NewAsyncTask() *logicsAsyncTask {
	logicsAsyncTaskOnce.Do(func() {
		logicsAsyncTaskInstance = &logicsAsyncTask{
			ctx:    context.Background(),
			logger: g.Log(),

			initInterval:     10 * time.Second,
			queryInterval:    30 * time.Second,
			errSleepInterval: 3 * time.Second,
			backoffIntervals: []time.Duration{
				2 * time.Second,
				3 * time.Second,
				5 * time.Second,
				10 * time.Second,
				30 * time.Second,
				1 * time.Minute,
				5 * time.Minute,
			},

			sigChanMap: make(map[model.AsyncTaskType]chan struct{}),

			handler: make(map[model.AsyncTaskType]model.AsyncTaskHandler),
		}
	})

	return logicsAsyncTaskInstance
}

func (o *logicsAsyncTask) RegisterHandler(op model.AsyncTaskType, handler model.AsyncTaskHandler) {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if _, ok := o.handler[op]; ok {
		panic(fmt.Sprintf("[AsyncTask]: handler already registered for task: %v", model.GetAsyncTaskType(op)))
	}
	o.handler[op] = handler
}

func (o *logicsAsyncTask) Start() {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	for op, handler := range o.handler {
		o.sigChanMap[op] = make(chan struct{}, 1000)
		go o.pushWorker(op, handler)
	}

	go o.startTimeoutMonitor()
}

func (o *logicsAsyncTask) AddTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte) error {
	return dao.AsyncTask.AddTask(ctx, tx, op, customID, content)
}

func (o *logicsAsyncTask) AddScheduledTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte, scheduledTime *gtime.Time) error {
	return dao.AsyncTask.AddScheduledTask(ctx, tx, op, customID, content, scheduledTime)
}

// 非阻塞方式发送通知，通知推送线程有新消息
func (o *logicsAsyncTask) WakeUp(op model.AsyncTaskType) {
	o.mutex.RLock()
	ch, ok := o.sigChanMap[op]
	o.mutex.RUnlock()

	if !ok {
		o.logger.Errorf(o.ctx, "[AsyncTask]: notify channel for task %v not found", model.GetAsyncTaskType(op))
		return
	}

	select {
	case ch <- struct{}{}:
		// o.logger.Debugf(o.ctx, "[AsyncTask]: notify channel for task %v send signal", model.GetAsyncTaskType(op))
	default:
		o.logger.Debugf(o.ctx, "[AsyncTask]: notify channel for task %v channel is full", model.GetAsyncTaskType(op))
	}
}

func (o *logicsAsyncTask) pushWorker(taskType model.AsyncTaskType, handler model.AsyncTaskHandler) {
	defer func() {
		if r := recover(); r != nil {
			o.logger.Errorf(o.ctx, "[AsyncTask]: push worker for task %v panic: %v", model.GetAsyncTaskType(taskType), r)
		}
	}()

	time.Sleep(o.initInterval)

	o.logger.Infof(o.ctx, "[AsyncTask]: push worker for task %v start", model.GetAsyncTaskType(taskType))

	// 状态变量
	var (
		nextFetchTime time.Time = time.Now() // pushWorker启动时，立即查询一次。
	)

	for {
		// 优先检查上下文取消
		select {
		case <-o.ctx.Done():
			o.logger.Infof(o.ctx, "[AsyncTask]: push worker for task %v received exit signal", model.GetAsyncTaskType(taskType))
			return
		default:
		}

		// 阻塞等待信号触发或者定时器触发，直到有任务可处理
		select {
		case <-o.sigChanMap[taskType]: // 信号触发，立即执行
			// o.logger.Debugf(o.ctx, "[AsyncTask]: push goroutine for task %v notify by signal", model.GetAsyncTaskType(taskType))
		case <-time.After(time.Until(nextFetchTime)): // 正常等待计时器触发
			// o.logger.Debugf(o.ctx, "[AsyncTask]: push goroutine for task %v notify by timer", model.GetAsyncTaskType(taskType))
		}

		// 获取并处理消息
		taskEntity, err := dao.AsyncTask.FetchPendingTask(o.ctx, taskType)
		if err != nil {
			nextFetchTime = time.Now().Add(o.errSleepInterval)
			continue
		}

		// 没有待处理任务可执行，获取下次处理时间最小的任务
		if taskEntity == nil {
			taskEntity, err := dao.AsyncTask.GetMinNextRetryTime(o.ctx, taskType)
			if err != nil {
				nextFetchTime = time.Now().Add(o.errSleepInterval)
				continue
			}
			// 无任务时使用查询间隔
			if taskEntity == nil {
				nextFetchTime = time.Now().Add(o.queryInterval)
				continue
			}
			// 有未来任务时精准等待
			nextFetchTime = time.Unix(taskEntity.NextRetryTime, 0)
			o.logger.Debugf(o.ctx, "[AsyncTask]: push goroutine for task %v no task to process, nextFetchTime: %+v", model.GetAsyncTaskType(taskType), nextFetchTime.Local())
			continue
		}

		// g.Log().Infof(o.ctx, "taskID: %d, now(): %+v, nextRetryTime: %+v, createTime: %+v", taskEntity.ID, gtime.Now().Unix(), taskEntity.NextRetryTime, taskEntity.CreateTime)

		// 不管消息处理成功还是失败，立即准备下次获取，避免消息堆积
		nextFetchTime = time.Now()

		taskInfo, err := model.ConvertAsyncTaskEntityToModel(taskEntity)
		if err != nil {
			o.logger.Errorf(o.ctx, "[AsyncTask]: push goroutine for task %v convert task entity to model error: %v", model.GetAsyncTaskType(taskType), err)
			continue
		}
		if err = o.handle(o.ctx, taskInfo, handler); err != nil {
			o.logger.Errorf(o.ctx, "[AsyncTask]: push goroutine for task %v handle task error: %v", model.GetAsyncTaskType(taskType), err)
			continue
		}
	}
}

func (o *logicsAsyncTask) handle(ctx context.Context, taskInfo *model.AsyncTask, handler model.AsyncTaskHandler) (err error) {
	var dataUpdate map[string]interface{}
	if err = handler(ctx, taskInfo); err != nil {
		dataUpdate = map[string]interface{}{
			dao.AsyncTask.Columns().Status:        model.AsyncTaskStatusPending,
			dao.AsyncTask.Columns().RetryCount:    taskInfo.RetryCount + 1,
			dao.AsyncTask.Columns().NextRetryTime: o.calculateNextRetryTime(taskInfo),
			dao.AsyncTask.Columns().Version:       taskInfo.Version,
		}
	} else {
		dataUpdate = map[string]interface{}{
			dao.AsyncTask.Columns().Status:  model.AsyncTaskStatusSuccess,
			dao.AsyncTask.Columns().Version: taskInfo.Version + 1,
		}
	}

	err = dao.AsyncTask.Update(ctx, taskInfo.ID, taskInfo.Version, dataUpdate)

	return
}

func (o *logicsAsyncTask) calculateNextRetryTime(taskInfo *model.AsyncTask) (nextRetryTimeStamp int64) {
	if taskInfo.RetryCount >= len(o.backoffIntervals) {
		return time.Now().Add(o.backoffIntervals[len(o.backoffIntervals)-1]).Unix()
	}
	return time.Now().Add(o.backoffIntervals[taskInfo.RetryCount]).Unix()
}

func (o *logicsAsyncTask) startTimeoutMonitor() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	time.Sleep(o.queryInterval)

	for {
		select {
		case <-o.ctx.Done():
			return
		case <-ticker.C:
			rowsAffected, err := dao.AsyncTask.ResetTimeoutTasks(o.ctx)
			if err != nil {
				o.logger.Errorf(o.ctx, "[AsyncTask]: reset timeout tasks error: %v", err)
				continue
			}
			if rowsAffected > 0 {
				o.logger.Infof(o.ctx, "[AsyncTask]: reset %d timeout tasks", rowsAffected)
			}
		}
	}
}
