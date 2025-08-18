package service

import (
	"GameEngine/internal/model"
	"context"
)

// IRanking 榜单服务接口
type IRanking interface {
	// 热门游戏榜单
	GetHotGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	// 获取本月新游戏
	GetThisMonthNewGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	// 获取即将上新的游戏
	GetUpcomingGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 分类榜单
	GetCategoryRanking(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	// 标签榜单
	GetTagRanking(ctx context.Context, tagID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 综合评分榜单
	GetComprehensiveRanking(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	// 高分游戏榜单
	GetTopRatedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 下载量榜单
	GetMostDownloadedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	// 收藏数榜单
	GetMostFavoritedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 相关游戏推荐
	GetRelatedGames(ctx context.Context, gameID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
}

var localRanking IRanking

func Ranking() IRanking {
	if localRanking == nil {
		panic("implement not found for interface IRanking, forgot register?")
	}
	return localRanking
}

func RegisterRanking(i IRanking) {
	localRanking = i
}
