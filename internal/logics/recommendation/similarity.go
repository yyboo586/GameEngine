package recommendation

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"context"
	"sort"
)

// SimilarityEngine 相似度引擎
type SimilarityEngine struct{}

// NewSimilarityEngine 创建相似度引擎实例
func NewSimilarityEngine() *SimilarityEngine {
	return &SimilarityEngine{}
}

// FindSimilarGames 查找相似游戏
func (se *SimilarityEngine) FindSimilarGames(ctx context.Context, targetGame *model.Game, limit int) ([]*model.Game, error) {
	// 获取目标游戏的标签
	targetTags, err := se.getGameTags(ctx, targetGame.ID)
	if err != nil {
		return nil, err
	}

	// 获取目标游戏的分类
	targetCategories, err := se.getGameCategories(ctx, targetGame.ID)
	if err != nil {
		return nil, err
	}

	// 基于标签相似度查找游戏
	tagSimilarGames, err := se.findGamesByTagSimilarity(ctx, targetTags, targetGame.ID, limit*2)
	if err != nil {
		return nil, err
	}

	// 基于分类相似度查找游戏
	categorySimilarGames, err := se.findGamesByCategorySimilarity(ctx, targetCategories, targetGame.ID, limit*2)
	if err != nil {
		return nil, err
	}

	// 合并并去重
	allGames := se.mergeAndDeduplicate(tagSimilarGames, categorySimilarGames)

	// 计算相似度分数
	scoredGames := se.calculateSimilarityScores(allGames, targetGame, targetTags, targetCategories)

	// 按相似度分数排序
	sort.Slice(scoredGames, func(i, j int) bool {
		return scoredGames[i].SimilarityScore > scoredGames[j].SimilarityScore
	})

	// 返回前limit个
	if len(scoredGames) > limit {
		scoredGames = scoredGames[:limit]
	}

	// 转换为Game模型
	var result []*model.Game
	for _, scoredGame := range scoredGames {
		result = append(result, scoredGame.Game)
	}

	return result, nil
}

// ScoredGame 带分数的游戏
type ScoredGame struct {
	Game            *model.Game
	SimilarityScore float64
}

// getGameTags 获取游戏标签
func (se *SimilarityEngine) getGameTags(ctx context.Context, gameID int64) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := dao.GameTag.Ctx(ctx).
		Fields("t_tag.*").
		LeftJoin("t_tag", "t_tag.id = t_game_tag.tag_id").
		Where("t_game_tag.game_id", gameID).
		Scan(&tags)
	return tags, err
}

// getGameCategories 获取游戏分类
func (se *SimilarityEngine) getGameCategories(ctx context.Context, gameID int64) ([]*model.Category, error) {
	var categories []*model.Category
	err := dao.GameCategory.Ctx(ctx).
		Fields("t_category.*").
		LeftJoin("t_category", "t_category.id = t_game_category.category_id").
		Where("t_game_category.game_id", gameID).
		Scan(&categories)
	return categories, err
}

// findGamesByTagSimilarity 基于标签相似度查找游戏
func (se *SimilarityEngine) findGamesByTagSimilarity(ctx context.Context, targetTags []*model.Tag, excludeGameID int64, limit int) ([]*model.Game, error) {
	if len(targetTags) == 0 {
		return nil, nil
	}

	var tagIDs []int64
	for _, tag := range targetTags {
		tagIDs = append(tagIDs, tag.ID)
	}

	var games []*model.Game
	err := dao.Game.Ctx(ctx).
		Fields("t_game.*").
		LeftJoin("t_game_tag", "t_game_tag.game_id = t_game.id").
		Where("t_game_tag.tag_id IN (?)", tagIDs).
		Where("t_game.id != ?", excludeGameID).
		Where("t_game.status", model.GameStatusPublished).
		Group("t_game.id").
		OrderDesc("COUNT(t_game_tag.tag_id)").
		Limit(limit).
		Scan(&games)

	return games, err
}

// findGamesByCategorySimilarity 基于分类相似度查找游戏
func (se *SimilarityEngine) findGamesByCategorySimilarity(ctx context.Context, targetCategories []*model.Category, excludeGameID int64, limit int) ([]*model.Game, error) {
	if len(targetCategories) == 0 {
		return nil, nil
	}

	var categoryIDs []int64
	for _, category := range targetCategories {
		categoryIDs = append(categoryIDs, category.ID)
	}

	var games []*model.Game
	err := dao.Game.Ctx(ctx).
		Fields("t_game.*").
		LeftJoin("t_game_category", "t_game_category.game_id = t_game.id").
		Where("t_game_category.category_id IN (?)", categoryIDs).
		Where("t_game.id != ?", excludeGameID).
		Where("t_game.status", model.GameStatusPublished).
		Group("t_game.id").
		OrderDesc("COUNT(t_game_category.category_id)").
		Limit(limit).
		Scan(&games)

	return games, err
}

// mergeAndDeduplicate 合并并去重
func (se *SimilarityEngine) mergeAndDeduplicate(games1, games2 []*model.Game) []*model.Game {
	gameMap := make(map[int64]*model.Game)

	// 添加第一组游戏
	for _, game := range games1 {
		gameMap[game.ID] = game
	}

	// 添加第二组游戏（如果有重复，会覆盖）
	for _, game := range games2 {
		gameMap[game.ID] = game
	}

	// 转换为切片
	var result []*model.Game
	for _, game := range gameMap {
		result = append(result, game)
	}

	return result
}

// calculateSimilarityScores 计算相似度分数
func (se *SimilarityEngine) calculateSimilarityScores(games []*model.Game, targetGame *model.Game, targetTags []*model.Tag, targetCategories []*model.Category) []*ScoredGame {
	var scoredGames []*ScoredGame

	for _, game := range games {
		score := se.calculateGameSimilarity(game, targetGame, targetTags, targetCategories)
		scoredGames = append(scoredGames, &ScoredGame{
			Game:            game,
			SimilarityScore: score,
		})
	}

	return scoredGames
}

// calculateGameSimilarity 计算单个游戏的相似度
func (se *SimilarityEngine) calculateGameSimilarity(game *model.Game, targetGame *model.Game, targetTags []*model.Tag, targetCategories []*model.Category) float64 {
	score := 0.0

	// 标签相似度（权重0.4）
	tagScore := se.calculateTagSimilarity(game.ID, targetTags)
	score += tagScore * 0.4

	// 分类相似度（权重0.3）
	categoryScore := se.calculateCategorySimilarity(game.ID, targetCategories)
	score += categoryScore * 0.3

	// 热度相似度（权重0.2）
	hotScore := se.calculateHotSimilarity(game, targetGame)
	score += hotScore * 0.2

	// 时间相似度（权重0.1）
	timeScore := se.calculateTimeSimilarity(game, targetGame)
	score += timeScore * 0.1

	return score
}

// calculateTagSimilarity 计算标签相似度
func (se *SimilarityEngine) calculateTagSimilarity(gameID int64, targetTags []*model.Tag) float64 {
	// 简化实现，实际应该查询数据库
	// 这里返回一个基础分数
	return 0.5
}

// calculateCategorySimilarity 计算分类相似度
func (se *SimilarityEngine) calculateCategorySimilarity(gameID int64, targetCategories []*model.Category) float64 {
	// 简化实现，实际应该查询数据库
	// 这里返回一个基础分数
	return 0.5
}

// calculateHotSimilarity 计算热度相似度
func (se *SimilarityEngine) calculateHotSimilarity(game *model.Game, targetGame *model.Game) float64 {
	// 基于下载量和收藏数的相似度
	gameHotness := float64(game.DownloadCount)*0.6 + float64(game.FavoriteCount)*0.4
	targetHotness := float64(targetGame.DownloadCount)*0.6 + float64(targetGame.FavoriteCount)*0.4

	if targetHotness == 0 {
		return 0.5
	}

	// 计算相似度，值越接近1越相似
	diff := se.abs(gameHotness - targetHotness)
	maxHotness := se.max(gameHotness, targetHotness)

	return 1.0 - (diff / maxHotness)
}

// calculateTimeSimilarity 计算时间相似度
func (se *SimilarityEngine) calculateTimeSimilarity(game *model.Game, targetGame *model.Game) float64 {
	// 基于发布时间的相似度
	if game.PublishTime == nil || targetGame.PublishTime == nil {
		return 0.5
	}

	gameTime := game.PublishTime.Time
	targetTime := targetGame.PublishTime.Time

	// 计算时间差的绝对值（小时）
	diffHours := se.abs(gameTime.Sub(targetTime).Hours())

	// 如果时间差小于24小时，相似度较高
	if diffHours < 24 {
		return 0.9
	} else if diffHours < 168 { // 7天
		return 0.7
	} else if diffHours < 720 { // 30天
		return 0.5
	} else {
		return 0.3
	}
}

// 辅助函数
func (se *SimilarityEngine) abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (se *SimilarityEngine) max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
