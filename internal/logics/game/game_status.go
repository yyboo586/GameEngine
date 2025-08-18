package game

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (gg *Game) SubmitForReview(ctx context.Context, id int64) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}
	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusInReview) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusInReview))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusInReview)
}

func (gg *Game) Approve(ctx context.Context, id int64) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}
	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusApproved) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusApproved))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusApproved)
}

func (gg *Game) Reject(ctx context.Context, id int64) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}
	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusInit) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusInit))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusInit)
}

func (gg *Game) PublishGameImmediately(ctx context.Context, id int64) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}

	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusPublished) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusPublished))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusPublished)
}

// TODO: 定时发布的问题
func (gg *Game) PreRegisterGame(ctx context.Context, id int64, publishTime *gtime.Time) (err error) {
	if publishTime.Before(gtime.Now()) {
		return fmt.Errorf("发布时间不能小于当前时间")
	}

	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}

	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusPreRegister) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusPreRegister))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusPreRegister)
}

func (gg *Game) UnpublishGame(ctx context.Context, id int64, unpublishReason string) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}

	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatusUnpublished) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatusUnpublished))
	}

	return gg.updateGameStatus(ctx, id, model.GameStatusUnpublished)
}

func (gg *Game) ListInReview(ctx context.Context, pageReq *model.PageReq) (out []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	query := dao.Game.Ctx(ctx).Where(dao.Game.Columns().Status, int(model.GameStatusInReview))

	total, err := query.Count()
	if err != nil {
		return
	}

	var entities []*entity.Game
	err = query.Page(pageReq.Page, pageReq.Size).OrderDesc(dao.Game.Columns().CreateTime).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		out = append(out, model.ConvertGameEntityToModel(entity))
	}

	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// 游戏状态管理 - 完整的状态流转逻辑
func (gg *Game) updateGameStatus(ctx context.Context, id int64, status model.GameStatus) (err error) {
	currentGame, err := gg.GetGameByID(ctx, id)
	if err != nil {
		return err
	}
	if currentGame == nil {
		return fmt.Errorf("游戏不存在")
	}

	// 验证状态流转是否合法
	if !gg.isValidStatusTransition(currentGame.Status, model.GameStatus(status)) {
		return fmt.Errorf("无效的状态转换: %s -> %s",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameStatusText(model.GameStatus(status)))
	}

	// 执行状态流转
	err = gg.executeStatusTransition(ctx, currentGame, currentGame.Status, model.GameStatus(status))
	return
}

// 状态流转验证
func (gg *Game) isValidStatusTransition(from, to model.GameStatus) bool {
	switch from {
	case model.GameStatusInit:
		// 初始状态 -> 审核中
		return to == model.GameStatusInReview

	case model.GameStatusInReview:
		// 审核中 -> 审核通过 或 初始状态(审核不通过)
		return to == model.GameStatusApproved || to == model.GameStatusInit

	case model.GameStatusApproved:
		// 审核通过 -> 可预约 或 已上架
		return to == model.GameStatusPreRegister || to == model.GameStatusPublished

	case model.GameStatusPreRegister:
		// 可预约 -> 已上架
		return to == model.GameStatusPublished

	case model.GameStatusPublished:
		// 已上架 -> 已下架
		return to == model.GameStatusUnpublished

	case model.GameStatusUnpublished:
		// 已下架 -> 初始状态(重新编辑)
		return to == model.GameStatusInit

	default:
		return false
	}
}

// 执行状态流转
func (gg *Game) executeStatusTransition(ctx context.Context, gameInfo *model.Game, from, to model.GameStatus) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		dataUpdate := map[string]interface{}{
			dao.Game.Columns().Status: int(to),
		}

		// 根据不同状态流转执行相应的业务逻辑
		switch to {
		case model.GameStatusInReview:
			// 提交审核
			err := gg.handleSubmitForReview(ctx, tx, gameInfo)
			if err != nil {
				return err
			}

		case model.GameStatusApproved:
			// 审核通过
			err := gg.handleApproved(ctx, tx, gameInfo)
			if err != nil {
				return err
			}

		case model.GameStatusPreRegister:
			// 设置为可预约状态
			dataUpdate[dao.Game.Columns().PublishTime] = gameInfo.PublishTime
			err := gg.handlePreRegister(ctx, tx, gameInfo)
			if err != nil {
				return err
			}

		case model.GameStatusPublished:
			// 游戏上架
			dataUpdate[dao.Game.Columns().PublishTime] = time.Now()
			err := gg.handlePublished(ctx, tx, gameInfo)
			if err != nil {
				return err
			}

		case model.GameStatusUnpublished:
			// 游戏下架
			err := gg.handleUnpublished(ctx, tx, gameInfo)
			if err != nil {
				return err
			}

		case model.GameStatusInit:
			// 重新编辑，清空发布时间
			dataUpdate[dao.Game.Columns().PublishTime] = nil
		}

		// 更新游戏状态
		_, err := dao.Game.Ctx(ctx).TX(tx).
			Where(dao.Game.Columns().ID, gameInfo.ID).
			Data(dataUpdate).
			Update()
		if err != nil {
			return err
		}

		return nil
	})
}

// 提交审核处理
func (gg *Game) handleSubmitForReview(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 验证必要信息
	if gameInfo.Name == "" || gameInfo.Developer == "" || gameInfo.Publisher == "" {
		return fmt.Errorf("游戏基本信息不完整，无法提交审核")
	}

	// 检查媒体文件是否上传
	mediaInfos, err := service.Game().GetMediaInfo(ctx, gameInfo.ID)
	if err != nil {
		return err
	}
	if len(mediaInfos) == 0 {
		return fmt.Errorf("请先上传游戏媒体文件")
	}

	// 可以在这里添加其他审核前的验证逻辑
	// 比如检查游戏描述长度、标签数量等

	return nil
}

// 审核通过处理
func (gg *Game) handleApproved(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 审核通过后的处理逻辑
	// 比如发送通知给开发者、更新审核时间等

	// 这里可以添加审核通过后的业务逻辑
	return nil
}

// 可预约状态处理
func (gg *Game) handlePreRegister(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 设置为可预约状态
	// 可以在这里添加预约相关的初始化逻辑

	// 比如设置预约开始时间、预约截止时间等
	return nil
}

// 游戏上架处理
func (gg *Game) handlePublished(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 游戏上架后的处理逻辑

	// 1. 通知所有预约用户
	err := gg.notifyReservedUsers(ctx, tx, gameInfo)
	if err != nil {
		return err
	}

	// 2. 更新游戏索引（如果有搜索功能）
	// err = service.Search().UpdateGameIndex(ctx, gameID)

	// 3. 清理预约数据（可选，根据业务需求决定）
	// err = service.Reservation().ClearReservations(ctx, tx, gameID)

	return nil
}

// 游戏下架处理
func (gg *Game) handleUnpublished(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 游戏下架后的处理逻辑

	// 1. 清理游戏缓存
	// err = service.Cache().ClearGameCache(ctx, gameID)

	// 2. 更新搜索索引
	// err = service.Search().RemoveGameIndex(ctx, gameID)

	// 3. 记录下架原因等

	return nil
}

// 通知预约用户
func (gg *Game) notifyReservedUsers(ctx context.Context, tx gdb.TX, gameInfo *model.Game) error {
	// 获取所有预约用户
	var reservations []*entity.GameReserve
	err := dao.GameReserve.Ctx(ctx).TX(tx).
		Where(dao.GameReserve.Columns().GameID, gameInfo.ID).
		Scan(&reservations)
	if err != nil {
		return err
	}

	// 这里可以添加通知逻辑
	// 比如发送推送通知、邮件、短信等

	return nil
}

// 批量更新游戏状态（用于定时任务等）
func (gg *Game) BatchUpdateGameStatus(ctx context.Context) error {
	// 处理预约到期的游戏自动上架
	err := gg.handleScheduledPublish(ctx)
	if err != nil {
		return err
	}

	return nil
}

// 处理定时发布
func (gg *Game) handleScheduledPublish(ctx context.Context) error {
	// 查找需要自动上架的游戏
	// 这里需要根据具体的业务需求来实现
	// 比如根据预约时间、发布时间等条件

	// 示例：查找所有可预约状态且到达预约时间的游戏
	// var games []*entity.Game
	// err := dao.Game.Ctx(ctx).
	//     Where(dao.Game.Columns().Status, int(model.GameStatusPreRegister)).
	//     Where("reserve_time <= ?", time.Now()).
	//     Scan(&games)

	// for _, game := range games {
	//     err = gg.updateGameStatus(ctx, game.ID, int(model.GameStatusPublished))
	//     if err != nil {
	//         g.Log().Errorf(ctx, "自动上架游戏失败: gameID=%d, error=%v", game.ID, err)
	//     }
	// }

	return nil
}
