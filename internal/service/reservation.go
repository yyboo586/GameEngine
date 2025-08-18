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

	// 获取游戏预约数量
	GetGameReservationCount(ctx context.Context, gameID int64) (int64, error)

	// 获取即将上新的游戏
	GetUpcomingGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 检查用户是否已预约
	IsUserReserved(ctx context.Context, userID, gameID int64) (bool, error)

	// 获取本月新游戏
	GetThisMonthNewGames(ctx context.Context, page, size int) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 批量获取游戏预约状态
	GetBatchReservationStatus(ctx context.Context, userID int64, gameIDs []int64) (map[int64]bool, error)

	// 获取热门预约游戏
	GetPopularReservationGames(ctx context.Context, limit int) (outs []*model.Game, err error)
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
