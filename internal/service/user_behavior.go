package service

import (
	"GameEngine/internal/model"
	"context"
)

// IUserBehavior 用户行为服务接口
type IUserBehavior interface {
	RecordBehavior(ctx context.Context, userID, gameID int64, behaviorType model.BehaviorType, ipAddress string) error

	// 搜索历史管理
	GetSearchHistory(ctx context.Context, userID int64, pageReq *model.PageReq) ([]*model.UserBehavior, *model.PageRes, error)
	ClearSearchHistory(ctx context.Context, userID int64) error
}

var localUserBehavior IUserBehavior

func UserBehavior() IUserBehavior {
	if localUserBehavior == nil {
		panic("implement not found for interface IUserBehavior, forgot register?")
	}
	return localUserBehavior
}

func RegisterUserBehavior(i IUserBehavior) {
	localUserBehavior = i
}
