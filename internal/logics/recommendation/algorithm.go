package recommendation

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// RecommendationAlgorithm 推荐算法核心
type RecommendationAlgorithm struct {
	hotScoreCalculator *HotScoreCalculator
	similarityEngine   *SimilarityEngine
}

// NewRecommendationAlgorithm 创建推荐算法实例
func NewRecommendationAlgorithm() *RecommendationAlgorithm {
	return &RecommendationAlgorithm{
		hotScoreCalculator: NewHotScoreCalculator(),
		similarityEngine:   NewSimilarityEngine(),
	}
}

// GetTodayPicks 获取今日精选
func (ra *RecommendationAlgorithm) GetTodayPicks(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	// 构建今日精选查询
	query := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished)

	total, err := query.Count()
	if err != nil {
		return
	}

	// 获取游戏列表，按热度分数排序
	var entityGames []*entity.Game
	err = query.OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}

	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	/*
		// 构建今日精选列表
		var todayPicks []*v1.TodayPickItem
		for _, game := range games {
			// 计算热度分数
			hotScore := ra.hotScoreCalculator.CalculateHotScore(game)

			// 判断是否为今日新发布
			isNewRelease := ra.isTodayRelease(game.PublishTime)

			// 判断是否为编辑推荐（基于评分和下载量）
			isEditorChoice := ra.isEditorChoice(game)

			// 生成推荐理由
			pickReason := ra.generatePickReason(game, isNewRelease, isEditorChoice)

			// 转换为API响应格式
			gameResponse := ra.convertModelToResponse(game)

			todayPicks = append(todayPicks, &v1.TodayPickItem{
				Game:           gameResponse,
				PickReason:     pickReason,
				HotScore:       hotScore,
				IsNewRelease:   isNewRelease,
				IsEditorChoice: isEditorChoice,
			})
		}

			// 如果有用户ID，进行个性化排序
			if req.UserID > 0 {
				todayPicks = ra.personalizationEngine.PersonalizeTodayPicks(ctx, req.UserID, todayPicks)
			}
	*/
	return
}

// GetSimilarGames 获取相似游戏
func (ra *RecommendationAlgorithm) GetSimilarGames(ctx context.Context, gameID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	// 获取目标游戏信息
	targetGame, err := service.Game().GetGameByID(ctx, gameID)
	if err != nil {
		return
	}

	// 使用相似度引擎计算相似游戏
	outs, err = ra.similarityEngine.FindSimilarGames(ctx, targetGame, pageReq.Size)
	if err != nil {
		return
	}

	pageRes = &model.PageRes{
		CurrentPage: pageReq.Page,
	}
	return
}

// GetRecommendationsByCategory 基于分类的推荐
func (ra *RecommendationAlgorithm) GetRecommendationsByCategory(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	// 获取分类下的游戏，按热度排序
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().Status+" != ?", model.GameStatusUnpublished).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&games)

	if err != nil {
		return
	}

	outs = make([]*model.Game, 0, len(games))
	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		CurrentPage: pageReq.Page,
	}
	return
}

// GetRecommendationsByTags 基于标签的推荐
func (ra *RecommendationAlgorithm) GetRecommendationsByTags(ctx context.Context, tagIDs []int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	// 获取包含指定标签的游戏，按热度排序
	var entityGames []*entity.Game
	query := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().Status+" != ?", model.GameStatusUnpublished)

	// 构建标签查询条件
	if len(tagIDs) > 0 {
		query = query.Where("id IN (SELECT DISTINCT game_id FROM t_game_tag WHERE tag_id IN (?)", tagIDs)
	}

	err = query.OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		CurrentPage: pageReq.Page,
	}
	return
}

// GetPopularRecommendations 获取热门推荐
func (ra *RecommendationAlgorithm) GetPopularRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}
	return ra.getPopularRecommendations(ctx, pageReq.Size)
}

// GetNewGameRecommendations 获取新游推荐
func (ra *RecommendationAlgorithm) GetNewGameRecommendations(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 20
	}

	// 获取最近30天发布的游戏
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", thirtyDaysAgo).
		OrderDesc(dao.Game.Columns().PublishTime).
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return
	}

	outs = make([]*model.Game, 0, len(entityGames))
	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		CurrentPage: pageReq.Page,
	}
	return
}

// 私有方法

// getPopularRecommendations 获取热门推荐
func (ra *RecommendationAlgorithm) getPopularRecommendations(ctx context.Context, limit int) (outs []*model.Game, pageRes *model.PageRes, err error) {
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().Status+" != ?", model.GameStatusUnpublished).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Limit(limit).
		Scan(&games)

	if err != nil {
		return
	}

	outs = make([]*model.Game, 0, len(games))
	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}

	pageRes = &model.PageRes{}
	return
}

// isTodayRelease 判断是否为今日发布
func (ra *RecommendationAlgorithm) isTodayRelease(publishTime *gtime.Time) bool {
	if publishTime == nil {
		return false
	}

	now := time.Now()
	publish := publishTime.Time

	return publish.Year() == now.Year() &&
		publish.YearDay() == now.YearDay()
}

// isEditorChoice 判断是否为编辑推荐
func (ra *RecommendationAlgorithm) isEditorChoice(game *model.Game) bool {
	// 评分≥4.0且下载量≥1000的游戏
	if game.RatingCount > 0 {
		avgRating := float64(game.RatingScore) / float64(game.RatingCount)
		return avgRating >= 4.0 && game.DownloadCount >= 1000
	}
	return false
}

// generatePickReason 生成推荐理由
func (ra *RecommendationAlgorithm) generatePickReason(game *model.Game, isNewRelease, isEditorChoice bool) string {
	var reasons []string

	if isNewRelease {
		reasons = append(reasons, "今日新发布")
	}

	if isEditorChoice {
		reasons = append(reasons, "编辑推荐")
	}

	if game.RatingCount > 0 {
		avgRating := float64(game.RatingScore) / float64(game.RatingCount)
		if avgRating >= 4.5 {
			reasons = append(reasons, "超高评分")
		} else if avgRating >= 4.0 {
			reasons = append(reasons, "高评分")
		}
	}

	if game.DownloadCount >= 10000 {
		reasons = append(reasons, "超热门")
	} else if game.DownloadCount >= 5000 {
		reasons = append(reasons, "热门")
	}

	if game.FavoriteCount >= 1000 {
		reasons = append(reasons, "高收藏")
	}

	if len(reasons) == 0 {
		reasons = append(reasons, "精选推荐")
	}

	return strings.Join(reasons, " · ")
}

// convertModelToResponse 将model.Game转换为v1.Game
func (ra *RecommendationAlgorithm) convertModelToResponse(in *model.Game) *v1.Game {
	return &v1.Game{
		ID:             in.ID,
		Name:           in.Name,
		DistributeType: model.GetGameDistributeTypeText(in.DistributeType),
		Developer:      in.Developer,
		Publisher:      in.Publisher,
		Description:    in.Description,
		Details:        in.Details,
		Status:         model.GetGameStatusText(in.Status),
		PublishTime:    in.PublishTime,
		ReserveCount:   in.ReserveCount,
		RatingScore:    in.RatingScore,
		RatingCount:    in.RatingCount,
		FavoriteCount:  in.FavoriteCount,
		DownloadCount:  in.DownloadCount,
		CreateTime:     in.CreateTime,
		UpdateTime:     in.UpdateTime,
	}
}
