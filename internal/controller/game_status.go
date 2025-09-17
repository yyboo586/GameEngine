package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
)

// 提交审核
func (c *gameController) SubmitGameForReview(ctx context.Context, req *v1.SubmitGameForReviewReq) (res *v1.SubmitGameForReviewRes, err error) {
	err = service.Game().SubmitForReview(ctx, req.ID)
	return
}

// 审核通过
func (c *gameController) ApproveGame(ctx context.Context, req *v1.ApproveGameReq) (res *v1.ApproveGameRes, err error) {
	err = service.Game().Approve(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return
}

// 拒绝游戏
func (c *gameController) RejectGame(ctx context.Context, req *v1.RejectGameReq) (res *v1.RejectGameRes, err error) {
	err = service.Game().Reject(ctx, req.ID)
	return
}

// 立即上架游戏
func (c *gameController) PublishGameImmediately(ctx context.Context, req *v1.PublishGameImmediatelyReq) (res *v1.PublishGameImmediatelyRes, err error) {
	err = service.Game().PublishGameImmediately(ctx, req.ID)
	// 通知游戏预约者，游戏发布了。
	go func() {
		gameInfo, err := service.Game().GetGameByID(context.Background(), req.ID)
		if err != nil {
			g.Log().Errorf(ctx, "get game by id error: %v", err)
			return
		}
		if gameInfo == nil {
			g.Log().Errorf(ctx, "game not found")
			return
		}
		users, err := service.Reservation().GetGameReservations(context.Background(), req.ID)
		if err != nil {
			g.Log().Errorf(ctx, "get game reservations error: %v", err)
			return
		}
		if len(users) == 0 {
			g.Log().Errorf(ctx, "no reservations found")
			return
		}

		var body map[string]interface{} = make(map[string]interface{})
		var userIDs []string = make([]string, 0, len(users))
		for _, reservation := range users {
			userIDs = append(userIDs, fmt.Sprintf("%d", reservation.UserID))
		}
		body["user_ids"] = userIDs
		body["content"] = map[string]string{
			"title":     "游戏已发布",
			"game_name": gameInfo.Name,
			"message":   "游戏已发布，请登录游戏引擎查看",
		}

		err = service.MQ().Publish(context.Background(), "core.push.users", body)
		if err != nil {
			g.Log().Errorf(ctx, "publish game published error: %v", err)
			return
		}
	}()
	return
}

// 预约发布游戏
func (c *gameController) SetGamePreRegister(ctx context.Context, req *v1.SetGamePreRegisterReq) (res *v1.SetGamePreRegisterRes, err error) {
	err = service.Game().PreRegisterGame(ctx, req.ID, req.PublishTime)
	return
}

// 下架游戏
func (c *gameController) UnpublishGame(ctx context.Context, req *v1.UnpublishGameReq) (res *v1.UnpublishGameRes, err error) {
	err = service.Game().UnpublishGame(ctx, req.ID, req.UnpublishReason)
	return
}

// 获取待审核的游戏
func (c *gameController) ListInReview(ctx context.Context, req *v1.ListInReviewReq) (res *v1.ListInReviewRes, err error) {
	games, pageRes, err := service.Game().ListInReview(ctx, &req.PageReq)
	if err != nil {
		return nil, err
	}
	res = &v1.ListInReviewRes{
		PageRes: *pageRes,
	}

	res.List, err = c.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}

	return
}
