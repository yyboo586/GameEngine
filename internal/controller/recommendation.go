package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

var RecommendationController = &recommendationController{}

// recommendationController 推荐控制器
type recommendationController struct{}

// setUserGameStatus 为登录用户设置游戏状态（预约和收藏）
func (c *recommendationController) setUserGameStatus(ctx context.Context, games []*v1.Game) error {
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

// GetTodayPicks 获取今日精选
func (c *recommendationController) GetTodayPicks(ctx context.Context, req *v1.GetTodayPicksReq) (res *v1.GetTodayPicksRes, err error) {
	games, pageRes, err := service.Recommendation().GetTodayPicks(ctx, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetTodayPicksRes{
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

// GetSimilarGames 获取相似游戏
func (c *recommendationController) GetSimilarGames(ctx context.Context, req *v1.GetSimilarGamesReq) (res *v1.GetSimilarGamesRes, err error) {
	games, pageRes, err := service.Recommendation().GetSimilarGames(ctx, req.ID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetSimilarGamesRes{
		PageRes: pageRes,
	}

	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

// GetRecommendationsByCategory 基于分类的推荐
func (c *recommendationController) GetRecommendationsByCategory(ctx context.Context, req *v1.GetRecommendationsByCategoryReq) (res *v1.GetRecommendationsByCategoryRes, err error) {
	games, pageRes, err := service.Recommendation().GetRecommendationsByCategory(ctx, req.CategoryID, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetRecommendationsByCategoryRes{
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

// GetRecommendationsByTags 基于标签的推荐
func (c *recommendationController) GetRecommendationsByTags(ctx context.Context, req *v1.GetRecommendationsByTagsReq) (res *v1.GetRecommendationsByTagsRes, err error) {
	games, pageRes, err := service.Recommendation().GetRecommendationsByTags(ctx, req.TagIDs, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetRecommendationsByTagsRes{
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

// GetPopularRecommendations 获取热门推荐
func (c *recommendationController) GetPopularRecommendations(ctx context.Context, req *v1.GetPopularRecommendationsReq) (res *v1.GetPopularRecommendationsRes, err error) {
	games, pageRes, err := service.Recommendation().GetPopularRecommendations(ctx, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetPopularRecommendationsRes{
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

// GetNewGameRecommendations 获取新游推荐
func (c *recommendationController) GetNewGameRecommendations(ctx context.Context, req *v1.GetNewGameRecommendationsReq) (res *v1.GetNewGameRecommendationsRes, err error) {
	games, pageRes, err := service.Recommendation().GetNewGameRecommendations(ctx, &req.PageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetNewGameRecommendationsRes{
		PageRes: pageRes,
	}
	res.List, err = GameController.getGameDetails(ctx, games)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}
