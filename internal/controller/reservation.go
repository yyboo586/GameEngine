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

// setUserGameStatus 为登录用户设置游戏状态（预约和收藏）
func (c *reservationController) setUserGameStatus(ctx context.Context, games []*v1.Game) error {
	if value := ctx.Value(model.UserInfoKey); value != nil {
		userID := value.(model.User).ID
		for _, g := range games {
			// 检查是否已预约
			var isReserved bool
			var err error
			isReserved, err = service.Reservation().IsUserReserved(ctx, userID, g.ID)
			if err != nil {
				return err
			}
			g.IsReserve = isReserved

			// 检查是否已收藏
			var isFavorited bool
			isFavorited, err = service.Game().IsUserFavorited(ctx, g.ID, userID)
			if err != nil {
				return err
			}
			g.IsFavorite = isFavorited
		}
	}
	return nil
}

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

	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
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

/*
// GetGameReservations 根据游戏ID获取预约用户列表
func (c *reservationController) GetGameReservations(ctx context.Context, req *v1.GetGameReservationsReq) (res *v1.GetGameReservationsRes, err error) {
	users, pageRes, err := service.Reservation().GetGameReservations(ctx, req.GameID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetGameReservationsRes{
		List:    make([]*v1.ReservationUser, 0, len(users)),
		PageRes: pageRes,
	}

	for _, user := range users {
		res.List = append(res.List, c.convertReservationUserModelToResponse(user))
	}

	return
}

// convertReservationUserModelToResponse 转换预约用户模型到响应
func (c *reservationController) convertReservationUserModelToResponse(in *model.ReservationUser) *v1.ReservationUser {
	return &v1.ReservationUser{
		ID:          in.ID,
		UserID:      in.UserID,
		UserName:    in.UserName,
		ReserveTime: in.ReserveTime,
	}
}
*/
