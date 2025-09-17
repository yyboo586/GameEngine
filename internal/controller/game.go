package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

var GameController = &gameController{}

type gameController struct{}

// setUserGameStatus 为登录用户设置游戏状态（预约和收藏）
func (c *gameController) setUserGameStatus(ctx context.Context, games []*v1.Game) error {
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
