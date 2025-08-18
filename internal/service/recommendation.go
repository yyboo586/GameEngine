package service

import (
	"GameEngine/internal/model"
	"context"
)

// IRecommendation 推荐服务接口
type IRecommendation interface {
	// 今日精选推荐
	GetTodayPicks(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 相似游戏推荐
	GetSimilarGames(ctx context.Context, gameID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 基于分类的推荐
	GetRecommendationsByCategory(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 基于标签的推荐
	GetRecommendationsByTags(ctx context.Context, tagIDs []int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 热门推荐
	GetPopularRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 新游推荐
	GetNewGameRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
}

var localRecommendation IRecommendation

func Recommendation() IRecommendation {
	if localRecommendation == nil {
		panic("implement not found for interface IRecommendation, forgot register?")
	}
	return localRecommendation
}

func RegisterRecommendation(i IRecommendation) {
	localRecommendation = i
}
