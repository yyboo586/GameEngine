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
	userInfo := ctx.Value(model.UserInfoKey).(model.User)
	err = service.Reservation().ReserveGame(ctx, userInfo.ID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.ReserveGameRes{}
	return
}

// CancelReservation 取消预约
func (c *reservationController) CancelReservation(ctx context.Context, req *v1.CancelReservationReq) (res *v1.CancelReservationRes, err error) {
	userInfo := ctx.Value(model.UserInfoKey).(model.User)
	err = service.Reservation().CancelReservation(ctx, userInfo.ID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.CancelReservationRes{}
	return
}

// GetUserReservations 获取用户预约列表
func (c *reservationController) GetUserReservations(ctx context.Context, req *v1.GetUserReservationsReq) (res *v1.GetUserReservationsRes, err error) {
	userInfo := ctx.Value(model.UserInfoKey).(model.User)
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

// GetThisMonthNewGames 获取本月新游戏
func (c *reservationController) GetThisMonthNewGames(ctx context.Context, req *v1.GetThisMonthNewGamesReq) (res *v1.GetThisMonthNewGamesRes, err error) {
	games, pageRes, err := service.Reservation().GetThisMonthNewGames(ctx, req.PageReq.Page, req.PageReq.Size)
	if err != nil {
		return
	}

	res = &v1.GetThisMonthNewGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetUpcomingGames 获取即将上新的游戏
func (c *reservationController) GetUpcomingGames(ctx context.Context, req *v1.GetUpcomingGamesReq) (res *v1.GetUpcomingGamesRes, err error) {
	games, pageRes, err := service.Reservation().GetUpcomingGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetUpcomingGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

/*
// GetGameReservationCount 获取游戏预约数量
func (c *reservationController) GetGameReservationCount(ctx context.Context, req *v1.GetGameReservationCountReq) (res *v1.GetGameReservationCountRes, err error) {
	count, err := service.Reservation().GetGameReservationCount(ctx, req.GameID)
	if err != nil {
		return
	}

	res = &v1.GetGameReservationCountRes{
		Count: count,
	}
	return
}



// IsUserReserved 检查用户是否已预约
func (c *reservationController) IsUserReserved(ctx context.Context, req *v1.IsUserReservedReq) (res *v1.IsUserReservedRes, err error) {
	isReserved, err := service.Reservation().IsUserReserved(ctx, req.UserID, req.GameID)
	if err != nil {
		return
	}

	res = &v1.IsUserReservedRes{
		IsReserved: isReserved,
	}
	return
}

*/

/*
// GetBatchReservationStatus 批量获取游戏预约状态
func (c *reservationController) GetBatchReservationStatus(ctx context.Context, req *v1.GetBatchReservationStatusReq) (res *v1.GetBatchReservationStatusRes, err error) {
	statusMap, err := service.Reservation().GetBatchReservationStatus(ctx, req.UserID, req.GameIDs)
	if err != nil {
		return
	}

	res = &v1.GetBatchReservationStatusRes{
		StatusMap: statusMap,
	}
	return
}

// GetPopularReservationGames 获取热门预约游戏
func (c *reservationController) GetPopularReservationGames(ctx context.Context, req *v1.GetPopularReservationGamesReq) (res *v1.GetPopularReservationGamesRes, err error) {
	games, err := service.Reservation().GetPopularReservationGames(ctx, req.PageReq.Size)
	if err != nil {
		return
	}

	res = &v1.GetPopularReservationGamesRes{
		Total: int64(len(games)),
	}
	for _, game := range games {
		res.List = append(res.List, GameController.convertModelToResponse(game))
	}
	return
}
*/
