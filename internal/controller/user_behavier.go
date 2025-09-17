package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
	"fmt"
)

var UserBehavierController = &userBehavierController{}

type userBehavierController struct{}

func (c *userBehavierController) GetSearchHistory(ctx context.Context, req *v1.GetSearchHistoryReq) (res *v1.GetSearchHistoryRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	outs, pageRes, err := service.UserBehavior().GetSearchHistory(ctx, userInfo.ID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetSearchHistoryRes{
		List:    make([]*v1.SearchHistoryItem, 0, len(outs)),
		PageRes: pageRes,
	}
	for _, out := range outs {
		res.List = append(res.List, c.convertModelToResponse(out))
	}
	return
}

func (c *userBehavierController) ClearSearchHistory(ctx context.Context, req *v1.ClearSearchHistoryReq) (res *v1.ClearSearchHistoryRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.UserBehavior().ClearSearchHistory(ctx, userInfo.ID)
	if err != nil {
		return nil, err
	}

	res = &v1.ClearSearchHistoryRes{}
	return
}

func (c *userBehavierController) convertModelToResponse(in *model.UserBehavior) (out *v1.SearchHistoryItem) {
	out = &v1.SearchHistoryItem{
		ID:            in.ID,
		UserID:        in.UserID,
		GameID:        in.GameID,
		BehaviorType:  model.GetBehaviorTypeString(in.BehaviorType),
		SearchKeyword: in.SearchKeyword,
		SearchTime:    in.BehaviorTime,
	}
	return
}

// PlayGame 玩游戏
func (c *userBehavierController) PlayGame(ctx context.Context, req *v1.PlayGameReq) (res *v1.PlayGameRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	// 检查游戏是否存在
	gameInfo, err := service.Game().GetGameByID(ctx, req.GameID)
	if err != nil {
		return nil, err
	}

	if gameInfo == nil {
		return nil, fmt.Errorf("游戏不存在")
	}

	if gameInfo.Status != model.GameStatusPublished {
		return nil, fmt.Errorf("游戏未发布")
	}

	// 记录玩游戏行为
	err = service.UserBehavior().RecordBehavior(ctx, userInfo.ID, req.GameID, model.BehaviorPlay, "", gameInfo.Name)
	if err != nil {
		return nil, err
	}

	res = &v1.PlayGameRes{}
	return
}

// GetPlayHistory 获取玩过游戏历史记录
func (c *userBehavierController) GetPlayHistory(ctx context.Context, req *v1.GetPlayHistoryReq) (res *v1.GetPlayHistoryRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	behaviors, pageRes, err := service.UserBehavior().GetPlayHistory(ctx, userInfo.ID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetPlayHistoryRes{
		List:    make([]*v1.PlayHistoryItem, 0, len(behaviors)),
		PageRes: pageRes,
	}

	// 获取游戏详情并转换为响应格式
	for _, behavior := range behaviors {
		game, err := service.Game().GetGameByID(ctx, behavior.GameID)
		if err != nil {
			return nil, err
		}
		if game == nil {
			continue // 游戏不存在，跳过
		}

		item := &v1.PlayHistoryItem{
			ID:          behavior.ID,
			GameID:      behavior.GameID,
			GameName:    game.Name,
			Developer:   game.Developer,
			Publisher:   game.Publisher,
			Description: game.Description,
			PlayTime:    behavior.BehaviorTime,
			IPAddress:   behavior.IPAddress,
		}

		// 检查是否已收藏
		isFavorited, err := service.Game().IsUserFavorited(ctx, behavior.GameID, userInfo.ID)
		if err != nil {
			return nil, err
		}
		item.IsFavorite = isFavorited

		// 检查是否已预约
		isReserved, err := service.Reservation().IsUserReserved(ctx, userInfo.ID, behavior.GameID)
		if err != nil {
			return nil, err
		}
		item.IsReserve = isReserved

		res.List = append(res.List, item)
	}

	return
}
