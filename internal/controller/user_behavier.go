package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
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
