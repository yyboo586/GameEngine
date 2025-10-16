package model

import (
	"GameEngine/internal/model/entity"
	"context"
	"encoding/json"
	"errors"
)

var (
	ErrNoRowsAffected = errors.New("concurrent update error: no rows affected")
)

// 任务类型
type AsyncTaskType int

const (
	_                                    AsyncTaskType = iota
	AsyncTaskTypeGameAutoPublish                       // 游戏预约，到时发布
	AsyncTaskTypeGameNotifyReservedUsers               // 游戏发布后，通知预约用户游戏已上线
)

// 任务执行状态
type AsyncTaskStatus int

const (
	AsyncTaskStatusPending    AsyncTaskStatus = iota // 待执行
	AsyncTaskStatusProcessing                        // 执行中
	AsyncTaskStatusSuccess                           // 执行成功
)

func GetAsyncTaskType(op AsyncTaskType) string {
	switch op {
	case AsyncTaskTypeGameAutoPublish:
		return "GameAutoPublish"
	case AsyncTaskTypeGameNotifyReservedUsers:
		return "GameNotifyReservedUsers"
	default:
		return "Unknown"
	}
}

type AsyncTaskHandler func(ctx context.Context, in *AsyncTask) error

type AsyncTask struct {
	ID            int64           `json:"id"`
	CustomID      string          `json:"custom_id"`
	TaskType      AsyncTaskType   `json:"task_type"`
	Status        AsyncTaskStatus `json:"status"`
	RetryCount    int             `json:"retry_count"`
	Content       interface{}     `json:"content"`
	Version       int             `json:"version"`
	NextRetryTime int64           `json:"next_retry_time"`
	CreateTime    int64           `json:"create_time"`
	UpdateTime    int64           `json:"update_time"`
}

func ConvertAsyncTaskEntityToModel(in *entity.AsyncTask) (out *AsyncTask, err error) {
	out = &AsyncTask{
		ID:            in.ID,
		CustomID:      in.CustomID,
		TaskType:      AsyncTaskType(in.TaskType),
		Status:        AsyncTaskStatus(in.Status),
		RetryCount:    in.RetryCount,
		Content:       in.Content,
		Version:       in.Version,
		NextRetryTime: in.NextRetryTime,
		CreateTime:    in.CreateTime,
		UpdateTime:    in.UpdateTime,
	}
	err = json.Unmarshal([]byte(in.Content), &out.Content)
	if err != nil {
		return nil, err
	}

	return out, nil
}
