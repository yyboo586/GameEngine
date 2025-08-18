package reservation

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"fmt"
	"time"

	"GameEngine/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReservationLogic 预约逻辑实现
type Reservation struct{}

// NewReservation 创建预约逻辑实例
func NewReservation() service.IReservation {
	return &Reservation{}
}

// ReserveGame 游戏预约
func (rl *Reservation) ReserveGame(ctx context.Context, userID, gameID int64) error {
	// 检查游戏是否存在且未上架
	gameInfo, err := service.Game().GetGameByID(ctx, gameID)
	if err != nil {
		return err
	}
	if gameInfo == nil {
		return fmt.Errorf("游戏不存在")
	}

	if gameInfo.Status != model.GameStatusUnpublished {
		return fmt.Errorf("只能预约未上架的游戏")
	}

	// 检查是否已经预约
	exists, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		Where(dao.GameReserve.Columns().UserID, userID).
		Exist()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	// 创建预约记录
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		dataInsert := map[string]interface{}{
			dao.GameReserve.Columns().GameID: gameID,
			dao.GameReserve.Columns().UserID: userID,
		}
		_, err = dao.GameReserve.Ctx(ctx).TX(tx).Data(dataInsert).Insert()
		if err != nil {
			return err
		}

		_, err = dao.Game.Ctx(ctx).TX(tx).
			Where(dao.Game.Columns().ID, gameID).
			Increment(dao.Game.Columns().ReserveCount, 1)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

// CancelReservation 取消预约
func (rl *Reservation) CancelReservation(ctx context.Context, userID, gameID int64) error {
	// 检查是否已经预约
	exists, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		Where(dao.GameReserve.Columns().UserID, userID).
		Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("未预约该游戏")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GameReserve.Ctx(ctx).TX(tx).
			Where(dao.GameReserve.Columns().GameID, gameID).
			Where(dao.GameReserve.Columns().UserID, userID).
			Delete()
		if err != nil {
			return err
		}

		_, err = dao.Game.Ctx(ctx).TX(tx).
			Where(dao.Game.Columns().ID, gameID).
			Decrement(dao.Game.Columns().ReserveCount, 1)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetUserReservations 获取用户预约列表
func (rl *Reservation) GetUserReservations(ctx context.Context, userID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().UserID, userID).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取预约的游戏列表
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Fields("t_game.*").
		LeftJoin("t_game_reserve", "t_game_reserve.game_id = t_game.id").
		Where("t_game_reserve.user_id", userID).
		OrderDesc("t_game_reserve.create_time").
		Page(pageReq.Page, pageReq.Size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}

	// 构建分页信息
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetGameReservationCount 获取游戏预约数量
func (rl *Reservation) GetGameReservationCount(ctx context.Context, gameID int64) (int64, error) {
	count, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		Count()
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

// GetUpcomingGames 获取即将上新的游戏
func (rl *Reservation) GetUpcomingGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusUnpublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusUnpublished).
		OrderDesc(dao.Game.Columns().CreateTime).
		Page(pageReq.Page, pageReq.Size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// IsUserReserved 检查用户是否已预约
func (rl *Reservation) IsUserReserved(ctx context.Context, userID, gameID int64) (bool, error) {
	count, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		Where(dao.GameReserve.Columns().UserID, userID).
		Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetThisMonthNewGames 获取本月新游戏
func (rl *Reservation) GetThisMonthNewGames(ctx context.Context, page, size int) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取本月开始时间
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", monthStart).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", monthStart).
		OrderDesc(dao.Game.Columns().PublishTime).
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}
	return
}

// GetBatchReservationStatus 批量获取游戏预约状态
func (rl *Reservation) GetBatchReservationStatus(ctx context.Context, userID int64, gameIDs []int64) (map[int64]bool, error) {
	if len(gameIDs) == 0 {
		return make(map[int64]bool), nil
	}

	// 查询用户预约的游戏
	var reservations []*entity.GameReserve
	err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().UserID, userID).
		Where(dao.GameReserve.Columns().GameID+" IN (?)", gameIDs).
		Scan(&reservations)
	if err != nil {
		return nil, err
	}

	// 构建预约状态映射
	statusMap := make(map[int64]bool)
	for _, gameID := range gameIDs {
		statusMap[gameID] = false
	}
	for _, reservation := range reservations {
		statusMap[reservation.GameID] = true
	}

	return statusMap, nil
}

// GetPopularReservationGames 获取热门预约游戏
func (rl *Reservation) GetPopularReservationGames(ctx context.Context, limit int) (outs []*model.Game, err error) {
	if limit == 0 {
		limit = 10
	}

	// 获取预约数量最多的游戏
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusUnpublished).
		OrderDesc(dao.Game.Columns().ReserveCount).
		Limit(limit).
		Scan(&games)
	if err != nil {
		return nil, err
	}

	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	return
}

// convertModelToResponse 将model.Game转换为v1.Game
func (rl *Reservation) convertModelToResponse(in *model.Game) *v1.Game {
	return &v1.Game{
		ID:             in.ID,
		Name:           in.Name,
		DistributeType: int(in.DistributeType),
		Developer:      in.Developer,
		Publisher:      in.Publisher,
		Description:    in.Description,
		Details:        in.Details,
		Status:         int(in.Status),
		PublishTime:    in.PublishTime,
		ReserveCount:   in.ReserveCount,
		RatingScore:    in.RatingScore,
		RatingCount:    in.RatingCount,
		FavoriteCount:  in.FavoriteCount,
		DownloadCount:  in.DownloadCount,
		CreateTime:     in.CreateTime,
		UpdateTime:     in.UpdateTime,
	}
}
