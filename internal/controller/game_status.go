package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/service"
	"context"
)

// 提交审核
func (c *gameController) SubmitGameForReview(ctx context.Context, req *v1.SubmitGameForReviewReq) (res *v1.SubmitGameForReviewRes, err error) {
	err = service.Game().SubmitForReview(ctx, req.ID)
	return
}

// 审核游戏
func (c *gameController) ApproveGame(ctx context.Context, req *v1.ApproveGameReq) (res *v1.ApproveGameRes, err error) {
	err = service.Game().Approve(ctx, req.ID)
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
