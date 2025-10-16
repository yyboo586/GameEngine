package game

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	ErrConcurrentUpdate = errors.New("并发操作，游戏信息已变更，更新失败")
)

type StateAction func(ctx context.Context, game *model.Game, data interface{}) (err error)

type Transition struct {
	TargetStatus model.GameStatus
	Action       StateAction
}

// 状态转移定义
var stateTransitionMap = map[model.GameStatus]map[model.GameEvent]*Transition{
	// Init(初始状态)
	model.GameStatusInit: {
		model.SubmitForReview: {
			TargetStatus: model.GameStatusInReview,
			Action:       handleSubmitForReview, // 由 handleSubmitForReview 处理
		},
	},
	// InReview(审核中)
	model.GameStatusInReview: {
		model.Approve: {
			TargetStatus: model.GameStatusApproved,
			Action:       handleApproved, // 由 handleApproved 处理
		},
		model.Reject: {
			TargetStatus: model.GameStatusInit,
			Action:       handleReject, // 由 handleReject 处理
		},
	},
	// Approved(审核通过)
	model.GameStatusApproved: {
		model.PreRegister: {
			TargetStatus: model.GameStatusPreRegister,
			Action:       handlePreRegister, // 由 handlePreRegister 处理
		},
		model.PublishNow: {
			TargetStatus: model.GameStatusPublished,
			Action:       handlePublished, // 由 handlePublished 处理
		},
		model.UpdateInfo: {
			TargetStatus: model.GameStatusInit,
			Action:       handleUpdateInfo, // 修改信息，回到初始状态
		},
	},
	// PreRegister(可预约)
	model.GameStatusPreRegister: {
		model.AutoPublish: {
			TargetStatus: model.GameStatusPublished,
			Action:       handlePublished, // 由 handlePublished 处理
		},
		model.CancelPreRegister: {
			TargetStatus: model.GameStatusApproved,
			Action:       handleCancelPreRegister, // 取消预约发布，回到审核通过状态
		},
		model.UpdateInfo: {
			TargetStatus: model.GameStatusInit,
			Action:       handleUpdateInfo, // 修改信息，回到初始状态
		},
	},
	// Published(已上架)
	model.GameStatusPublished: {
		model.UnpublishNow: {
			TargetStatus: model.GameStatusUnpublished,
			Action:       handleUnpublished, // 由 handleUnpublished 处理
		},
		model.UpdateVersion: {
			TargetStatus: model.GameStatusInit,
			Action:       handleUpdateVersion, // 更新版本，回到初始状态
		},
	},
}

// SubmitForReview 提交审核
func (gg *Game) SubmitForReview(ctx context.Context, id int64) (err error) {
	return gg.HandleGameEvent(ctx, id, model.SubmitForReview, nil)
}

// Approve 审核通过
func (gg *Game) Approve(ctx context.Context, id int64) (err error) {
	return gg.HandleGameEvent(ctx, id, model.Approve, nil)
}

// Reject 审核拒绝
func (gg *Game) Reject(ctx context.Context, id int64) (err error) {
	return gg.HandleGameEvent(ctx, id, model.Reject, nil)
}

// PreRegisterGame 预约发布游戏
func (gg *Game) PreRegisterGame(ctx context.Context, id int64, publishTime *gtime.Time) (err error) {
	data := map[string]interface{}{
		"publish_time": publishTime,
	}
	return gg.HandleGameEvent(ctx, id, model.PreRegister, data)
}

// PublishGameImmediately 立即发布游戏
func (gg *Game) PublishGameImmediately(ctx context.Context, id int64) (err error) {
	return gg.HandleGameEvent(ctx, id, model.PublishNow, nil)
}

// CancelPreRegisterGame 取消预约发布
func (gg *Game) CancelPreRegisterGame(ctx context.Context, gameID int64) error {
	return gg.HandleGameEvent(ctx, gameID, model.CancelPreRegister, nil)
}

// UnpublishGame 下架游戏
func (gg *Game) UnpublishGame(ctx context.Context, id int64, unpublishReason string) (err error) {
	// TODO: 可以在这里记录下架原因到数据库
	return gg.HandleGameEvent(ctx, id, model.UnpublishNow, nil)
}

// UpdateGameInfo 更新游戏信息（需要回到初始状态）
func (gg *Game) UpdateGameInfo(ctx context.Context, gameID int64) error {
	return gg.HandleGameEvent(ctx, gameID, model.UpdateInfo, nil)
}

// UpdateGameVersion 更新游戏版本（需要回到初始状态）
func (gg *Game) UpdateGameVersion(ctx context.Context, gameID int64) error {
	return gg.HandleGameEvent(ctx, gameID, model.UpdateVersion, nil)
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
	err = query.Page(pageReq.Page, pageReq.Size).OrderAsc(dao.Game.Columns().CreateTime).Scan(&entities)
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

// 提交审核处理
func handleSubmitForReview(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	// 验证必要信息
	if gameInfo.Name == "" || gameInfo.Developer == "" || gameInfo.Publisher == "" {
		return fmt.Errorf("游戏基本信息不完整，无法提交审核")
	}

	// 检查媒体文件是否上传
	err := service.Game().CheckMediaInfo(ctx, gameInfo)
	if err != nil {
		return err
	}

	// 可以在这里添加其他审核前的验证逻辑
	// 比如检查游戏描述长度、标签数量等
	updateData := map[string]interface{}{
		dao.Game.Columns().Status: int(model.GameStatusInReview),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

	return nil
}

// 审核通过处理
func handleApproved(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	updateData := map[string]interface{}{
		dao.Game.Columns().Version: gameInfo.Version + 1,
		dao.Game.Columns().Status:  int(model.GameStatusApproved),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

	return nil
}

// 审核拒绝处理
func handleReject(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	updateData := map[string]interface{}{
		dao.Game.Columns().Version: gameInfo.Version + 1,
		dao.Game.Columns().Status:  int(model.GameStatusInit),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

	return nil
}

// 预约发布
func handlePreRegister(ctx context.Context, gameInfo *model.Game, data interface{}) (err error) {
	publishTime := data.(map[string]interface{})["publish_time"].(*gtime.Time)

	// 确保时间比较时使用相同的时区
	now := gtime.Now()
	if publishTime.Before(now) {
		return fmt.Errorf("发布时间不能小于当前时间")
	}

	updateData := map[string]interface{}{
		dao.Game.Columns().Status:      int(model.GameStatusPreRegister),
		dao.Game.Columns().PublishTime: publishTime,
		dao.Game.Columns().Version:     gameInfo.Version + 1,
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.Game.Ctx(ctx).TX(tx).
			Where(dao.Game.Columns().ID, gameInfo.ID).
			Where(dao.Game.Columns().Version, gameInfo.Version).
			Data(updateData).
			Update()
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ErrConcurrentUpdate
		}

		// 任务内容
		taskContent := map[string]interface{}{
			"game_id":      gameInfo.ID,
			"publish_time": publishTime,
		}
		contentBytes, err := json.Marshal(taskContent)
		if err != nil {
			return fmt.Errorf("序列化任务内容失败: %v", err)
		}

		// 添加定时任务，设置next_retry_time为发布时间
		g.Log().Infof(ctx, "创建定时发布任务: gameID=%d, publishTime=%s (local), publishTimeUTC=%s",
			gameInfo.ID, publishTime.Format("2006-01-02 15:04:05"), publishTime.UTC().Format("2006-01-02 15:04:05"))

		err = service.AsyncTask().AddScheduledTask(ctx, tx, model.AsyncTaskTypeGameAutoPublish, "", contentBytes, publishTime)
		if err != nil {
			return fmt.Errorf("添加自动发布任务失败: %v", err)
		}

		g.Log().Infof(ctx, "游戏预约发布成功: gameID=%d, publishTime=%s", gameInfo.ID, publishTime.Format("2006-01-02 15:04:05"))
		return nil
	})

	service.AsyncTask().WakeUp(model.AsyncTaskTypeGameAutoPublish)

	return err
}

// 游戏上架处理
func handlePublished(ctx context.Context, gameInfo *model.Game, data interface{}) (err error) {
	updateData := map[string]interface{}{
		dao.Game.Columns().Version:     gameInfo.Version + 1,
		dao.Game.Columns().Status:      int(model.GameStatusPublished),
		dao.Game.Columns().PublishTime: gtime.Now(),
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.Game.Ctx(ctx).
			Where(dao.Game.Columns().ID, gameInfo.ID).
			Where(dao.Game.Columns().Version, gameInfo.Version).
			Data(updateData).
			Update()
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ErrConcurrentUpdate
		}

		// 添加通知预约用户任务
		taskContent := map[string]interface{}{
			"game_id": gameInfo.ID,
		}
		contentBytes, err := json.Marshal(taskContent)
		if err != nil {
			return fmt.Errorf("序列化任务内容失败: %v", err)
		}

		err = service.AsyncTask().AddTask(ctx, tx, model.AsyncTaskTypeGameNotifyReservedUsers, "", contentBytes)
		if err != nil {
			return fmt.Errorf("添加通知预约用户任务失败: %v", err)
		}

		return nil
	})

	service.AsyncTask().WakeUp(model.AsyncTaskTypeGameNotifyReservedUsers)

	return err
}

// 修改游戏信息
func handleUpdateInfo(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	updateData := map[string]interface{}{
		dao.Game.Columns().Status: int(model.GameStatusInit),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

	return nil
}

func handleCancelPreRegister(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		updateData := map[string]interface{}{
			dao.Game.Columns().Status: int(model.GameStatusApproved),
		}
		result, err := dao.Game.Ctx(ctx).TX(tx).
			Where(dao.Game.Columns().ID, gameInfo.ID).
			Where(dao.Game.Columns().Version, gameInfo.Version).
			Data(updateData).
			Update()
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return ErrConcurrentUpdate
		}

		// 删除对应的自动发布任务
		customID := fmt.Sprintf("game_auto_publish_%d", gameInfo.ID)
		_, err = dao.AsyncTask.Ctx(ctx).TX(tx).
			Where(dao.AsyncTask.Columns().CustomID, customID).
			Where(dao.AsyncTask.Columns().TaskType, model.AsyncTaskTypeGameAutoPublish).
			Where(dao.AsyncTask.Columns().Status, model.AsyncTaskStatusPending).
			Delete()
		if err != nil {
			g.Log().Errorf(ctx, "删除自动发布任务失败: gameID=%d, customID=%s, error=%v",
				gameInfo.ID, customID, err)
		}

		g.Log().Infof(ctx, "取消预约发布成功: gameID=%d", gameInfo.ID)
		return nil
	})
}

func handleUnpublished(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	updateData := map[string]interface{}{
		dao.Game.Columns().Status: int(model.GameStatusUnpublished),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

	return nil
}

func handleUpdateVersion(ctx context.Context, gameInfo *model.Game, data interface{}) error {
	updateData := map[string]interface{}{
		dao.Game.Columns().Status: int(model.GameStatusInit),
	}
	result, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().ID, gameInfo.ID).
		Where(dao.Game.Columns().Version, gameInfo.Version).
		Data(updateData).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrConcurrentUpdate
	}

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
	var games []*entity.Game
	err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, int(model.GameStatusPreRegister)).
		Where("publish_time <= ?", time.Now()).
		Scan(&games)

	if err != nil {
		return err
	}

	for _, game := range games {
		err = gg.HandleGameEvent(ctx, game.ID, model.AutoPublish, nil)
		if err != nil {
			g.Log().Errorf(ctx, "自动上架游戏失败: gameID=%d, error=%v", game.ID, err)
		}
	}

	return nil
}

// HandleGameEvent 基于事件驱动的状态转换统一入口
func (gg *Game) HandleGameEvent(ctx context.Context, gameID int64, event model.GameEvent, data interface{}) error {
	// 获取当前游戏信息
	currentGame, err := gg.GetGameByID(ctx, gameID)
	if err != nil {
		return err
	}

	// 获取当前状态的事件映射
	eventMap, exists := stateTransitionMap[currentGame.Status]
	if !exists {
		return fmt.Errorf("当前状态[%s]不支持任何事件",
			model.GetGameStatusText(currentGame.Status))
	}

	// 获取事件对应的转换配置
	transition, exists := eventMap[event]
	if !exists {
		return fmt.Errorf("当前状态[%s]不支持事件[%s]",
			model.GetGameStatusText(currentGame.Status),
			model.GetGameEventText(event))
	}

	// 执行状态转换
	return gg.executeEventTransition(ctx, currentGame, event, transition, data)
}

// executeEventTransition 执行基于事件的状态转换
func (gg *Game) executeEventTransition(ctx context.Context, gameInfo *model.Game, event model.GameEvent, transition *Transition, data interface{}) error {
	// 执行自定义的转换动作
	if transition.Action != nil {
		err := transition.Action(ctx, gameInfo, data)
		if err != nil {
			return err
		}
	}

	// 记录状态变更日志
	g.Log().Infof(ctx, "游戏状态变更: gameID=%d, event=%s, %s -> %s",
		gameInfo.ID,
		model.GetGameEventText(event),
		model.GetGameStatusText(gameInfo.Status),
		model.GetGameStatusText(transition.TargetStatus))

	return nil
}

func (gg *Game) HandleGameAutoPublish(ctx context.Context, task *model.AsyncTask) (err error) {
	// 解析任务内容
	taskContent, ok := task.Content.(map[string]interface{})
	if !ok {
		return fmt.Errorf("任务内容格式错误")
	}

	gameID, ok := taskContent["game_id"].(float64)
	if !ok {
		return fmt.Errorf("游戏ID格式错误")
	}

	return gg.HandleGameEvent(ctx, int64(gameID), model.AutoPublish, nil)
}

// NotifyReservedUsers
func (gg *Game) NotifyReservedUsers(ctx context.Context, task *model.AsyncTask) (err error) {
	// 解析任务内容
	taskContent, ok := task.Content.(map[string]interface{})
	if !ok {
		return fmt.Errorf("任务内容格式错误")
	}

	gameID, ok := taskContent["game_id"].(float64)
	if !ok {
		return fmt.Errorf("游戏ID格式错误")
	}

	// 获取当前游戏信息
	gameInfo, err := gg.GetGameByID(ctx, int64(gameID))
	if err != nil {
		return fmt.Errorf("获取游戏信息失败: %v", err)
	}

	// 获取所有预约用户
	users, err := service.Reservation().GetGameReservations(ctx, gameInfo.ID)
	if err != nil {
		return
	}
	if len(users) == 0 {
		return
	}

	var body map[string]interface{} = make(map[string]interface{})
	var userIDs []string = make([]string, 0, len(users))
	for _, reservation := range users {
		userIDs = append(userIDs, fmt.Sprintf("%d", reservation.UserID))
	}
	body["user_ids"] = userIDs
	body["content"] = map[string]interface{}{
		"title":     "游戏已发布",
		"game_id":   gameInfo.ID,
		"game_name": gameInfo.Name,
		"message":   "游戏已发布，请登录游戏引擎查看",
	}

	err = service.MQ().Publish(ctx, "core.push.users", body)
	if err != nil {
		return
	}
	return nil
}
