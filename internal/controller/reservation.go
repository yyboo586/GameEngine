package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

var ReservationController = &reservationController{}

// reservation 预约控制器
type reservationController struct{}

// ReserveGame 游戏预约
func (c *reservationController) ReserveGame(ctx context.Context, req *v1.ReserveGameReq) (res *v1.ReserveGameRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Reservation().ReserveGame(ctx, userInfo.ID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.ReserveGameRes{}
	return
}

// CancelReservation 取消预约
func (c *reservationController) CancelReservation(ctx context.Context, req *v1.CancelReservationReq) (res *v1.CancelReservationRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Reservation().CancelReservation(ctx, userInfo.ID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.CancelReservationRes{}
	return
}

// GetUserReservations 获取用户预约列表
func (c *reservationController) GetUserReservations(ctx context.Context, req *v1.GetUserReservationsReq) (res *v1.GetUserReservationsRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	games, pageRes, err := service.Reservation().GetUserReservations(ctx, userInfo.ID, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetUserReservationsRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// IsUserReserved 检查用户是否已预约
func (c *reservationController) IsUserReserved(ctx context.Context, req *v1.IsUserReservedReq) (res *v1.IsUserReservedRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	isReserved, err := service.Reservation().IsUserReserved(ctx, userInfo.ID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.IsUserReservedRes{
		IsReserved: isReserved,
	}
	return
}
