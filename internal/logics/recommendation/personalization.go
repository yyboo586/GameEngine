package recommendation

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"context"
	"sort"
)

// PersonalizationEngine 个性化推荐引擎
type PersonalizationEngine struct{}

// NewPersonalizationEngine 创建个性化推荐引擎实例
func NewPersonalizationEngine() *PersonalizationEngine {
	return &PersonalizationEngine{}
}

// PersonalizeTodayPicks 个性化今日精选排序
func (pe *PersonalizationEngine) PersonalizeTodayPicks(ctx context.Context, userID int64, picks []*v1.TodayPickItem) []*v1.TodayPickItem {
	// 获取用户偏好
	userPreferences, err := pe.GetUserPreferences(ctx, userID)
	if err != nil {
		// 如果获取失败，返回原始排序
		return picks
	}

	// 根据用户偏好重新排序
	sort.Slice(picks, func(i, j int) bool {
		scoreI := pe.calculatePersonalizedScore(picks[i], userPreferences)
		scoreJ := pe.calculatePersonalizedScore(picks[j], userPreferences)
		return scoreI > scoreJ
	})

	return picks
}

// GetUserPreferences 获取用户偏好
func (pe *PersonalizationEngine) GetUserPreferences(ctx context.Context, userID int64) (*model.UserPreferences, error) {
	// 获取用户历史行为
	userBehaviors, err := pe.getUserGameHistory(ctx, userID, 100)
	if err != nil {
		return nil, err
	}

	// 分析用户偏好
	preferences := pe.analyzeUserPreferences(userBehaviors)
	return preferences, nil
}

// GeneratePersonalizedRecommendations 生成个性化推荐
func (pe *PersonalizationEngine) GeneratePersonalizedRecommendations(ctx context.Context, userPreferences *model.UserPreferences, limit int) ([]*model.Game, error) {
	// 基于用户偏好生成推荐
	var games []*model.Game
	query := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().Status+" != ?", model.GameStatusDeleted)

	// 如果有偏好分类，优先推荐
	if userPreferences.FavoriteCategory != "" {
		query = query.Where("id IN (SELECT DISTINCT game_id FROM t_game_category WHERE category_id IN (SELECT id FROM t_category WHERE name = ?))", userPreferences.FavoriteCategory)
	}

	err := query.OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Limit(limit).
		Scan(&games)

	if err != nil {
		return nil, err
	}

	return games, nil
}

// 私有方法

// getUserGameHistory 获取用户游戏历史
func (pe *PersonalizationEngine) getUserGameHistory(ctx context.Context, userID int64, limit int) ([]*model.GameBehavior, error) {
	var behaviors []*model.GameBehavior
	err := dao.GameBehavior.Ctx(ctx).
		Where(dao.GameBehavior.Columns().UserID, userID).
		OrderDesc(dao.GameBehavior.Columns().BehaviorTime).
		Limit(limit).
		Scan(&behaviors)
	return behaviors, err
}

// analyzeUserPreferences 分析用户偏好
func (pe *PersonalizationEngine) analyzeUserPreferences(behaviors []*model.GameBehavior) *model.UserPreferences {
	preferences := &model.UserPreferences{
		UserID:        behaviors[0].UserID, // 假设至少有一个行为
		FavoriteTags:  make([]string, 0),
		GamePlayStyle: "balanced", // 默认平衡型
	}

	// 统计用户偏好
	// tagPreferences := make(map[string]float64)
	// categoryPreferences := make(map[string]float64)

	for _, behavior := range behaviors {
		// 根据行为类型给予不同权重
		// weight := 1.0
		switch behavior.BehaviorType {
		case model.BehaviorTypeFavorite:
			// weight = 3.0 // 收藏权重最高
		case model.BehaviorTypeDownload:
			// weight = 2.0 // 下载次之
		case model.BehaviorTypeRating:
			// weight = 1.5 // 评分再次
		case model.BehaviorTypeView:
			// weight = 1.0 // 查看最低
		}

		// 这里可以进一步分析游戏分类和标签偏好
		// 暂时简化处理
	}

	// 设置偏好分类（简化实现）
	preferences.FavoriteCategory = "动作枪战" // 默认分类

	return preferences
}

// calculatePersonalizedScore 计算个性化分数
func (pe *PersonalizationEngine) calculatePersonalizedScore(pick *v1.TodayPickItem, userPreferences *model.UserPreferences) float64 {
	// 基础热度分数
	baseScore := pick.HotScore

	// 个性化加分
	personalBonus := 0.0

	// 检查是否为今日新发布（用户偏好新鲜内容）
	if pick.IsNewRelease {
		personalBonus += 50
	}

	// 检查是否为编辑推荐（用户偏好高质量内容）
	if pick.IsEditorChoice {
		personalBonus += 30
	}

	// 这里可以进一步根据用户历史偏好计算个性化分数
	// 比如检查游戏分类、标签是否匹配用户偏好

	return baseScore + personalBonus
}
