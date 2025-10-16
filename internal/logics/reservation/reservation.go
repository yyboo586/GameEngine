package reservation

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"fmt"

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

	if gameInfo.Status != model.GameStatusPreRegister {
		return fmt.Errorf("该游戏未开放预约")
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
	exists, err := dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		Where(dao.GameReserve.Columns().UserID, userID).
		Exist()
	if err != nil {
		return false, err
	}
	return exists, nil
}

// GetGameReservations 根据游戏ID获取预约用户列表
func (rl *Reservation) GetGameReservations(ctx context.Context, gameID int64) (outs []*model.ReservationUser, err error) {
	// 检查游戏是否存在
	err = service.Game().AssertExists(ctx, gameID)
	if err != nil {
		return nil, err
	}

	// 获取预约用户列表
	var entities []*entity.GameReserve
	err = dao.GameReserve.Ctx(ctx).
		Where(dao.GameReserve.Columns().GameID, gameID).
		OrderDesc(dao.GameReserve.Columns().CreateTime).
		Scan(&entities)
	if err != nil {
		return nil, err
	}

	for _, entity := range entities {
		outs = append(outs, rl.convertReservationEntityToModel(entity))
	}

	return
}

// convertReservationEntityToModel 转换预约实体到模型
func (rl *Reservation) convertReservationEntityToModel(in *entity.GameReserve) *model.ReservationUser {
	return &model.ReservationUser{
		ID:          in.ID,
		UserID:      in.UserID,
		UserName:    fmt.Sprintf("用户%d", in.UserID), // 临时用户名，实际项目中应该关联用户表
		ReserveTime: in.CreateTime.Format("2006-01-02 15:04:05"),
	}
}
