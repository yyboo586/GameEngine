package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/service"
	"context"
)

var RankingController = &rankingController{}

// rankingController 榜单控制器
type rankingController struct{}

// GetHotGames 获取热门游戏榜单
func (c *rankingController) GetHotGames(ctx context.Context, req *v1.GetHotGamesReq) (res *v1.GetHotGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetHotGames(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetHotGamesRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetNewGames 获取新游榜单
func (c *rankingController) GetNewGames(ctx context.Context, req *v1.GetNewGamesReq) (res *v1.GetNewGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetNewGames(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetNewGamesRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

/*
// GetTopRatedGames 获取高分游戏榜单
func (c *rankingController) GetTopRatedGames(ctx context.Context, req *v1.GetTopRatedGamesReq) (res *v1.GetTopRatedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetTopRatedGames(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetTopRatedGamesRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetMostDownloadedGames 获取下载量榜单
func (c *rankingController) GetMostDownloadedGames(ctx context.Context, req *v1.GetMostDownloadedGamesReq) (res *v1.GetMostDownloadedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetMostDownloadedGames(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetMostDownloadedGamesRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetMostFavoritedGames 获取收藏数榜单
func (c *rankingController) GetMostFavoritedGames(ctx context.Context, req *v1.GetMostFavoritedGamesReq) (res *v1.GetMostFavoritedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetMostFavoritedGames(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetMostFavoritedGamesRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetCategoryRanking 获取分类榜单
func (c *rankingController) GetCategoryRanking(ctx context.Context, req *v1.GetCategoryRankingReq) (res *v1.GetCategoryRankingRes, err error) {
	games, pageRes, err := service.Ranking().GetCategoryRanking(ctx, req.CategoryID, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetCategoryRankingRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetTagRanking 获取标签榜单
func (c *rankingController) GetTagRanking(ctx context.Context, req *v1.GetTagRankingReq) (res *v1.GetTagRankingRes, err error) {
	games, pageRes, err := service.Ranking().GetTagRanking(ctx, req.TagID, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetTagRankingRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}

// GetComprehensiveRanking 获取综合评分榜单
func (c *rankingController) GetComprehensiveRanking(ctx context.Context, req *v1.GetComprehensiveRankingReq) (res *v1.GetComprehensiveRankingRes, err error) {
	games, pageRes, err := service.Ranking().GetComprehensiveRanking(ctx, req.Page, req.Size)
	if err != nil {
		return
	}

	res = &v1.GetComprehensiveRankingRes{
		List:    games,
		PageRes: pageRes,
	}
	return
}
*/
