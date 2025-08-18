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
	games, pageRes, err := service.Ranking().GetHotGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetHotGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetThisMonthNewGames 获取本月新游戏
func (c *rankingController) GetThisMonthNewGames(ctx context.Context, req *v1.GetThisMonthNewGamesReq) (res *v1.GetThisMonthNewGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetThisMonthNewGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetThisMonthNewGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetUpcomingGames 获取即将上新的游戏
func (c *reservationController) GetUpcomingGames(ctx context.Context, req *v1.GetUpcomingGamesReq) (res *v1.GetUpcomingGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetUpcomingGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetUpcomingGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetCategoryRanking 获取分类榜单
func (c *rankingController) GetCategoryRanking(ctx context.Context, req *v1.GetCategoryRankingReq) (res *v1.GetCategoryRankingRes, err error) {
	games, pageRes, err := service.Ranking().GetCategoryRanking(ctx, req.CategoryID, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetCategoryRankingRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetTagRanking 获取标签榜单
func (c *rankingController) GetTagRanking(ctx context.Context, req *v1.GetTagRankingReq) (res *v1.GetTagRankingRes, err error) {
	games, pageRes, err := service.Ranking().GetTagRanking(ctx, req.TagID, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetTagRankingRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetTodayRecommend 获取今日推荐
func (c *rankingController) GetComprehensiveRanking(ctx context.Context, req *v1.GetTodayRecommendReq) (res *v1.GetTodayRecommendRes, err error) {
	games, pageRes, err := service.Ranking().GetComprehensiveRanking(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetTodayRecommendRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetTopRatedGames 获取高分游戏榜单
func (c *rankingController) GetTopRatedGames(ctx context.Context, req *v1.GetTopRatedGamesReq) (res *v1.GetTopRatedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetTopRatedGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetTopRatedGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

/*
// GetMostDownloadedGames 获取下载量榜单
func (c *rankingController) GetMostDownloadedGames(ctx context.Context, req *v1.GetMostDownloadedGamesReq) (res *v1.GetMostDownloadedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetMostDownloadedGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetMostDownloadedGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}

// GetMostFavoritedGames 获取收藏数榜单
func (c *rankingController) GetMostFavoritedGames(ctx context.Context, req *v1.GetMostFavoritedGamesReq) (res *v1.GetMostFavoritedGamesRes, err error) {
	games, pageRes, err := service.Ranking().GetMostFavoritedGames(ctx, &req.PageReq)
	if err != nil {
		return
	}

	res = &v1.GetMostFavoritedGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return
	}
	return
}
*/

// GetRelatedGames 获取相关游戏推荐
func (c *rankingController) GetRelatedGames(ctx context.Context, req *v1.GetRelatedGamesReq) (res *v1.GetRelatedGamesRes, err error) {
	games, pageRes, err := service.Recommendation().GetSimilarGames(ctx, req.GameID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetRelatedGamesRes{
		List:    make([]*v1.Game, 0, len(games)),
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	return
}
