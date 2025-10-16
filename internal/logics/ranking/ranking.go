package ranking

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"sort"
	"time"
)

// RankingLogic 榜单逻辑实现
type Ranking struct{}

// NewRanking 创建榜单逻辑实例
func NewRanking() service.IRanking {
	return &Ranking{}
}

// GetHotGames 获取热门游戏榜单
func (rl *Ranking) GetHotGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按热度分数排序
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc("(download_count * 0.5 + favorite_count * 0.3 + rating_score * 0.2)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)
	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetThisMonthNewGames 获取本月新游戏
func (rl *Ranking) GetThisMonthNewGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取本月开始时间
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", monthStart).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", monthStart).
		OrderDesc(dao.Game.Columns().PublishTime).
		Page(pageReq.Page, pageReq.Size).
		Scan(&games)
	if err != nil {
		return nil, nil, err
	}

	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetUpcomingGames 获取即将上新的游戏
func (rl *Ranking) GetUpcomingGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPreRegister).
		Count()
	if err != nil {
		return nil, nil, err
	}

	var games []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPreRegister).
		OrderDesc(dao.Game.Columns().ReserveCount).
		Page(pageReq.Page, pageReq.Size).
		Scan(&games)
	if err != nil {
		return nil, nil, err
	}

	for _, game := range games {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetCategoryRanking 获取分类榜单
func (rl *Ranking) GetCategoryRanking(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_category WHERE category_id = ?)", categoryID).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_category WHERE category_id = ?)", categoryID).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetTagRanking 获取标签榜单
func (rl *Ranking) GetTagRanking(ctx context.Context, tagID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_tag WHERE tag_id = ?)", tagID).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_tag WHERE tag_id = ?)", tagID).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetComprehensiveRanking 获取综合评分榜单
func (rl *Ranking) GetComprehensiveRanking(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按综合评分排序
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount+" > 0").
		OrderDesc("(download_count * 0.3 + favorite_count * 0.2 + (rating_score / rating_count) * 0.5)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetTopRatedGames 获取高分游戏榜单
func (rl *Ranking) GetTopRatedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按平均评分排序
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount+" > 0").
		OrderDesc("(rating_score / rating_count)").
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetMostDownloadedGames 获取下载量榜单
func (rl *Ranking) GetMostDownloadedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按下载量排序
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc(dao.Game.Columns().DownloadCount).
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetMostFavoritedGames 获取收藏数榜单
func (rl *Ranking) GetMostFavoritedGames(ctx context.Context, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按收藏数排序
	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc(dao.Game.Columns().FavoriteCount).
		Page(pageReq.Page, pageReq.Size).
		Scan(&entityGames)

	if err != nil {
		return nil, nil, err
	}

	for _, game := range entityGames {
		outs = append(outs, model.ConvertGameEntityToModel(game))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// GetRelatedGames 获取相关游戏推荐
func (rl *Ranking) GetRelatedGames(ctx context.Context, gameID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 获取目标游戏信息
	targetGame, err := service.Game().GetGameByID(ctx, gameID)
	if err != nil {
		return
	}

	// 获取目标游戏的标签
	targetTags, err := service.Metadata().GetTagsByGameID(ctx, gameID)
	if err != nil {
		return
	}

	// 获取目标游戏的分类
	targetCategory, err := service.Metadata().GetCategoryByGameID(ctx, gameID)
	if err != nil {
		return
	}

	// 基于标签相似度查找游戏
	tagSimilarGames, err := rl.findGamesByTagSimilarity(ctx, targetTags, gameID, pageReq.Size/2)
	if err != nil {
		return
	}

	// 基于分类相似度查找游戏
	categorySimilarGames, err := rl.findGamesByCategorySimilarity(ctx, targetCategory, gameID, pageReq.Size/2)
	if err != nil {
		return
	}

	// 合并并去重
	allGames := rl.mergeAndDeduplicateGames(tagSimilarGames, categorySimilarGames)

	// 计算相似度分数并排序
	scoredGames := rl.calculateSimilarityScores(allGames, targetGame, targetTags, targetCategory)

	// 按相似度分数排序
	sort.Slice(scoredGames, func(i, j int) bool {
		return scoredGames[i].SimilarityScore > scoredGames[j].SimilarityScore
	})

	// 返回前limit个
	if len(scoredGames) > pageReq.Size {
		scoredGames = scoredGames[:pageReq.Size]
	}

	// 转换为Game模型
	for _, scoredGame := range scoredGames {
		outs = append(outs, scoredGame.Game)
	}

	return
}

// ScoredGame 带分数的游戏
type ScoredGame struct {
	Game            *model.Game
	SimilarityScore float64
}

// findGamesByTagSimilarity 基于标签相似度查找游戏
func (rl *Ranking) findGamesByTagSimilarity(ctx context.Context, targetTags []*model.Tag, excludeGameID int64, limit int) ([]*model.Game, error) {
	if len(targetTags) == 0 {
		return nil, nil
	}

	var tagIDs []int64
	for _, tag := range targetTags {
		tagIDs = append(tagIDs, tag.ID)
	}

	var entityGames []*entity.Game
	err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id != ?", excludeGameID).
		Where("id IN (SELECT game_id FROM t_game_tag WHERE tag_id IN (?)", tagIDs).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Limit(limit).
		Scan(&entityGames)

	if err != nil {
		return nil, err
	}

	var games []*model.Game
	for _, game := range entityGames {
		games = append(games, model.ConvertGameEntityToModel(game))
	}
	return games, nil
}

// findGamesByCategorySimilarity 基于分类相似度查找游戏
func (rl *Ranking) findGamesByCategorySimilarity(ctx context.Context, targetCategory *model.Category, excludeGameID int64, limit int) ([]*model.Game, error) {
	if targetCategory == nil {
		return nil, nil
	}

	var categoryIDs []int64
	categoryIDs = append(categoryIDs, targetCategory.ID)

	var entityGames []*entity.Game
	err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id != ?", excludeGameID).
		Where("id IN (SELECT game_id FROM t_game_category WHERE category_id IN (?)", categoryIDs).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Limit(limit).
		Scan(&entityGames)

	if err != nil {
		return nil, err
	}

	var games []*model.Game
	for _, game := range entityGames {
		games = append(games, model.ConvertGameEntityToModel(game))
	}
	return games, nil
}

// mergeAndDeduplicateGames 合并并去重游戏列表
func (rl *Ranking) mergeAndDeduplicateGames(gameLists ...[]*model.Game) []*model.Game {
	gameMap := make(map[int64]*model.Game)

	for _, games := range gameLists {
		for _, game := range games {
			if _, exists := gameMap[game.ID]; !exists {
				gameMap[game.ID] = game
			}
		}
	}

	var result []*model.Game
	for _, game := range gameMap {
		result = append(result, game)
	}
	return result
}

// calculateSimilarityScores 计算相似度分数
func (rl *Ranking) calculateSimilarityScores(games []*model.Game, targetGame *model.Game, targetTags []*model.Tag, targetCategory *model.Category) []*ScoredGame {
	var scoredGames []*ScoredGame

	for _, game := range games {
		score := rl.calculateGameSimilarityScore(game, targetGame, targetTags, targetCategory)
		scoredGames = append(scoredGames, &ScoredGame{
			Game:            game,
			SimilarityScore: score,
		})
	}

	return scoredGames
}

// calculateGameSimilarityScore 计算单个游戏的相似度分数
func (rl *Ranking) calculateGameSimilarityScore(game *model.Game, targetGame *model.Game, targetTags []*model.Tag, targetCategory *model.Category) float64 {
	score := 0.0

	// 开发商相似度 (权重: 0.3)
	if game.Developer == targetGame.Developer {
		score += 0.3
	}

	// 标签相似度 (权重: 0.4)
	tagScore := rl.calculateTagSimilarity(game.ID, targetTags)
	score += tagScore * 0.4

	// 分类相似度 (权重: 0.3)
	categoryScore := rl.calculateCategorySimilarity(game.ID, targetCategory)
	score += categoryScore * 0.3

	// 热度加成 (权重: 0.1)
	popularityScore := rl.calculatePopularityScore(game)
	score += popularityScore * 0.1

	return score
}

// calculateTagSimilarity 计算标签相似度
func (rl *Ranking) calculateTagSimilarity(gameID int64, targetTags []*model.Tag) float64 {
	if len(targetTags) == 0 {
		return 0
	}

	// 获取当前游戏的标签
	currentTags, err := service.Metadata().GetTagsByGameID(context.Background(), gameID)
	if err != nil {
		return 0.5 // 出错时返回基础分数
	}

	if len(currentTags) == 0 {
		return 0
	}

	// 计算Jaccard相似度
	commonTags := 0
	targetTagMap := make(map[int64]bool)
	currentTagMap := make(map[int64]bool)

	for _, tag := range targetTags {
		targetTagMap[tag.ID] = true
	}
	for _, tag := range currentTags {
		currentTagMap[tag.ID] = true
		if targetTagMap[tag.ID] {
			commonTags++
		}
	}

	unionSize := len(targetTags) + len(currentTags) - commonTags
	if unionSize == 0 {
		return 0
	}

	return float64(commonTags) / float64(unionSize)
}

// calculateCategorySimilarity 计算分类相似度
func (rl *Ranking) calculateCategorySimilarity(gameID int64, targetCategory *model.Category) float64 {
	if targetCategory == nil {
		return 0
	}

	// 获取当前游戏的分类
	currentCategory, err := service.Metadata().GetCategoryByGameID(context.Background(), gameID)
	if err != nil {
		return 0.5 // 出错时返回基础分数
	}
	if currentCategory == nil {
		return 0
	}

	// 计算分类相似度（分类通常较少，权重可以更高）
	commonCategories := 0
	targetCategoryMap := make(map[int64]bool)
	currentCategoryMap := make(map[int64]bool)

	targetCategoryMap[targetCategory.ID] = true

	currentCategoryMap[currentCategory.ID] = true
	if targetCategoryMap[currentCategory.ID] {
		commonCategories++
	}

	unionSize := 2 - commonCategories
	if unionSize == 0 {
		return 0
	}

	return float64(commonCategories) / float64(unionSize)
}

// calculatePopularityScore 计算热度分数
func (rl *Ranking) calculatePopularityScore(game *model.Game) float64 {
	// 基于下载量、收藏数、评分计算热度分数
	downloadScore := float64(game.DownloadCount) / 1000.0 // 每1000次下载得1分
	favoriteScore := float64(game.FavoriteCount) / 100.0  // 每100次收藏得1分
	ratingScore := game.AverageRating / 5.0               // 评分归一化

	// 限制最高分数
	totalScore := downloadScore + favoriteScore + ratingScore
	if totalScore > 1.0 {
		totalScore = 1.0
	}

	return totalScore
}
