package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

// AddRating 添加游戏评分
func (c *gameController) AddRating(ctx context.Context, req *v1.AddGameRatingReq) (res *v1.AddGameRatingRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Game().AddRating(ctx, req.GameID, userInfo.ID, req.Rating)
	if err != nil {
		return &v1.AddGameRatingRes{}, err
	}

	return &v1.AddGameRatingRes{}, nil
}
