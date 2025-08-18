package service

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"context"
)

// IRanking 榜单服务接口
type IRanking interface {
	// 热门游戏榜单
	GetHotGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 新游榜单
	GetNewGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 高分游戏榜单
	GetTopRatedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 下载量榜单
	GetMostDownloadedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 收藏数榜单
	GetMostFavoritedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 分类榜单
	GetCategoryRanking(ctx context.Context, categoryID int64, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 标签榜单
	GetTagRanking(ctx context.Context, tagID int64, page, size int) ([]*v1.Game, *model.PageRes, error)

	// 综合评分榜单
	GetComprehensiveRanking(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error)
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
