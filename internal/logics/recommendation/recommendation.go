package recommendation

import (
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

// Recommendation 推荐逻辑实现
type Recommendation struct {
	algorithm *RecommendationAlgorithm
}

// NewRecommendation 创建推荐逻辑实例
func NewRecommendation() service.IRecommendation {
	return &Recommendation{
		algorithm: NewRecommendationAlgorithm(),
	}
}

// GetTodayPicks 获取今日精选
func (rl *Recommendation) GetTodayPicks(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetTodayPicks(ctx, pageReq)
}

// GetSimilarGames 获取相似游戏
func (rl *Recommendation) GetSimilarGames(ctx context.Context, gameID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetSimilarGames(ctx, gameID, pageReq)
}

// GetRecommendationsByCategory 基于分类的推荐
func (rl *Recommendation) GetRecommendationsByCategory(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetRecommendationsByCategory(ctx, categoryID, pageReq)
}

// GetRecommendationsByTags 基于标签的推荐
func (rl *Recommendation) GetRecommendationsByTags(ctx context.Context, tagIDs []int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetRecommendationsByTags(ctx, tagIDs, pageReq)
}

// GetPopularRecommendations 获取热门推荐
func (rl *Recommendation) GetPopularRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetPopularRecommendations(ctx, pageReq)
}

// GetNewGameRecommendations 获取新游推荐
func (rl *Recommendation) GetNewGameRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	return rl.algorithm.GetNewGameRecommendations(ctx, pageReq)
}
