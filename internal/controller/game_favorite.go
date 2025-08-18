package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

// AddFavorite 添加游戏收藏
func (c *gameController) AddFavorite(ctx context.Context, req *v1.AddGameFavoriteReq) (res *v1.AddGameFavoriteRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Game().AddFavorite(ctx, req.GameID, userInfo.ID)
	if err != nil {
		return &v1.AddGameFavoriteRes{}, err
	}

	return &v1.AddGameFavoriteRes{}, nil
}

// RemoveFavorite 取消游戏收藏
func (c *gameController) RemoveFavorite(ctx context.Context, req *v1.RemoveGameFavoriteReq) (res *v1.RemoveGameFavoriteRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Game().RemoveFavorite(ctx, req.GameID, userInfo.ID)
	if err != nil {
		return &v1.RemoveGameFavoriteRes{}, err
	}

	return &v1.RemoveGameFavoriteRes{}, nil
}

// GetGameFavorite 获取用户游戏收藏
func (c *gameController) GetGameFavorite(ctx context.Context, req *v1.GetGameFavoriteReq) (res *v1.GetGameFavoriteRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	outs, pageRes, err := service.Game().GetUserFavorites(ctx, userInfo.ID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetGameFavoriteRes{
		PageRes: pageRes,
	}
	res.List, err = c.getGameDetails(ctx, outs)
	if err != nil {
		return nil, err
	}

	return
}
