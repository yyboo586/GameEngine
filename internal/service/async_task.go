package service

import (
	"GameEngine/internal/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

// IAsyncTask 异步任务接口
type IAsyncTask interface {
	// 注册任务处理函数
	RegisterHandler(taskType model.AsyncTaskType, handler model.AsyncTaskHandler)

	// 添加任务
	AddTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte) error

	// 添加定时任务，指定执行时间
	AddScheduledTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte, scheduledTime *gtime.Time) error

	// 唤醒任务
	WakeUp(taskType model.AsyncTaskType)

	// 启动异步任务处理线程
	Start()
}

var (
	localAsyncTask IAsyncTask
)

func AsyncTask() IAsyncTask {
	if localAsyncTask == nil {
		panic("implement not found for interface IAsyncTask, forgot register?")
	}
	return localAsyncTask
}

func RegisterAsyncTask(i IAsyncTask) {
	localAsyncTask = i
}
