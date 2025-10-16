package dao

import (
	"GameEngine/internal/dao/internal"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"database/sql"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
)

// asyncTaskDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type asyncTaskDao struct {
	*internal.AsyncTaskDao
}

var (
	// AsyncTask is globally public accessible object for table tools_gen_table operations.
	AsyncTask = asyncTaskDao{
		internal.NewAsyncTaskDao(),
	}
)

func (asyncTask *asyncTaskDao) AddTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte) error {
	data := map[string]interface{}{
		asyncTask.Columns().CustomID:      customID,
		asyncTask.Columns().TaskType:      op,
		asyncTask.Columns().Content:       string(content),
		asyncTask.Columns().NextRetryTime: gtime.Now().UTC().Unix(),
		asyncTask.Columns().CreateTime:    gtime.Now().UTC().Unix(),
		asyncTask.Columns().UpdateTime:    gtime.Now().UTC().Unix(),
	}
	_, err := asyncTask.Ctx(ctx).TX(tx).Data(data).Insert()
	if err != nil {
		return err
	}

	return nil
}

// AddScheduledTask 添加定时任务，指定执行时间
func (asyncTask *asyncTaskDao) AddScheduledTask(ctx context.Context, tx gdb.TX, op model.AsyncTaskType, customID string, content []byte, scheduledTime *gtime.Time) error {
	data := map[string]interface{}{
		asyncTask.Columns().CustomID:      customID,
		asyncTask.Columns().TaskType:      op,
		asyncTask.Columns().Content:       string(content),
		asyncTask.Columns().NextRetryTime: scheduledTime.UTC().Unix(),
		asyncTask.Columns().CreateTime:    gtime.Now().UTC().Unix(),
		asyncTask.Columns().UpdateTime:    gtime.Now().UTC().Unix(),
	}
	_, err := asyncTask.Ctx(ctx).TX(tx).Data(data).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (asyncTask *asyncTaskDao) FetchPendingTask(ctx context.Context, op model.AsyncTaskType) (out *entity.AsyncTask, err error) {
	out = &entity.AsyncTask{}

	err = asyncTask.Ctx(ctx).
		Where(asyncTask.Columns().TaskType, op).
		Where(asyncTask.Columns().Status, model.AsyncTaskStatusPending).
		WhereLT(asyncTask.Columns().NextRetryTime, gtime.Now().UTC().Unix()).
		OrderAsc(asyncTask.Columns().NextRetryTime).
		Scan(&out)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	result, err := asyncTask.Ctx(ctx).
		Where(asyncTask.Columns().ID, out.ID).
		Where(asyncTask.Columns().Version, out.Version).
		Data(map[string]interface{}{
			asyncTask.Columns().Status:     model.AsyncTaskStatusProcessing,
			asyncTask.Columns().Version:    out.Version + 1,
			asyncTask.Columns().UpdateTime: gtime.Now().UTC().Unix(),
		}).
		Update()
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, model.ErrNoRowsAffected
	}

	out.Version += 1
	return out, nil
}

func (asyncTask *asyncTaskDao) GetMinNextRetryTime(ctx context.Context, op model.AsyncTaskType) (out *entity.AsyncTask, err error) {
	out = &entity.AsyncTask{}
	err = asyncTask.Ctx(ctx).
		Where(asyncTask.Columns().TaskType, op).
		Where(asyncTask.Columns().Status, model.AsyncTaskStatusPending).
		OrderAsc(asyncTask.Columns().NextRetryTime).
		Limit(1).
		Scan(out)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return out, nil
}

func (asyncTask *asyncTaskDao) Update(ctx context.Context, id int64, version int, data map[string]interface{}) (err error) {
	data[asyncTask.Columns().UpdateTime] = gtime.Now().UTC().Unix()
	result, err := asyncTask.Ctx(ctx).
		Where(asyncTask.Columns().ID, id).
		Where(asyncTask.Columns().Version, version).
		Data(data).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return model.ErrNoRowsAffected
	}

	return nil
}

func (asyncTask *asyncTaskDao) ResetTimeoutTasks(ctx context.Context) (rowsAffected int64, err error) {
	result, err := asyncTask.Ctx(ctx).
		Where(asyncTask.Columns().Status, model.AsyncTaskStatusProcessing).
		WhereGT(asyncTask.Columns().UpdateTime, gtime.Now().Add(-24*time.Hour).Unix()).
		Data(map[string]interface{}{
			asyncTask.Columns().Status:     model.AsyncTaskStatusPending,
			asyncTask.Columns().Version:    gdb.Raw("version + 1"),
			asyncTask.Columns().UpdateTime: gtime.Now().Unix(),
		}).
		Update()
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}
