package service

import (
	"GameEngine/internal/model"
	"context"
)

// IReservation 预约服务接口
type IReservation interface {
	// 游戏预约
	ReserveGame(ctx context.Context, userID, gameID int64) error

	// 取消预约
	CancelReservation(ctx context.Context, userID, gameID int64) error

	// 获取用户预约列表
	GetUserReservations(ctx context.Context, userID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 检查用户是否已预约
	IsUserReserved(ctx context.Context, userID, gameID int64) (bool, error)
}

var localReservation IReservation

func Reservation() IReservation {
	if localReservation == nil {
		panic("implement not found for interface IReservation, forgot register?")
	}
	return localReservation
}

func RegisterReservation(i IReservation) {
	localReservation = i
}
