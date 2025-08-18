package ranking

import (
	"GameEngine/internal/model"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// ScoreCalculator 分数计算器
type ScoreCalculator struct {
	// 可配置的权重参数
	downloadWeight float64
	favoriteWeight float64
	ratingWeight   float64
	timeDecayDays  int
	timeDecayBonus float64
}

// NewScoreCalculator 创建分数计算器
func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{
		downloadWeight: 0.4,
		favoriteWeight: 0.3,
		ratingWeight:   0.2,
		timeDecayDays:  7,
		timeDecayBonus: 100,
	}
}

// CalculateHotScore 计算热度分数
func (sc *ScoreCalculator) CalculateHotScore(game *model.Game) float64 {
	// 基础分数：下载量×0.4 + 收藏数×0.3 + 评分×0.2
	baseScore := float64(game.DownloadCount)*sc.downloadWeight +
		float64(game.FavoriteCount)*sc.favoriteWeight +
		float64(game.RatingScore)*sc.ratingWeight

	// 时间衰减因子：最近发布的游戏获得额外加分
	timeBonus := sc.calculateTimeBonus(game.PublishTime)

	return baseScore + timeBonus
}

// CalculateRatingScore 计算评分分数
func (sc *ScoreCalculator) CalculateRatingScore(game *model.Game) float64 {
	if game.RatingCount == 0 {
		return 0
	}

	// 平均评分
	avgRating := float64(game.RatingScore) / float64(game.RatingCount)

	// 评分分数 = 平均评分 × 评分次数权重
	return avgRating * float64(game.RatingCount) * 0.1
}

// CalculateComprehensiveScore 计算综合分数
func (sc *ScoreCalculator) CalculateComprehensiveScore(game *model.Game) float64 {
	// 综合分数 = 热度分数 × 0.6 + 评分分数 × 0.4
	hotScore := sc.CalculateHotScore(game)
	ratingScore := sc.CalculateRatingScore(game)

	return hotScore*0.6 + ratingScore*0.4
}

// calculateTimeBonus 计算时间加分
func (sc *ScoreCalculator) calculateTimeBonus(publishTime *gtime.Time) float64 {
	if publishTime == nil {
		return 0
	}

	daysSincePublish := time.Since(publishTime.Time).Hours() / 24
	if daysSincePublish <= float64(sc.timeDecayDays) {
		// 7天内线性衰减，新游戏获得最高加分
		decayRatio := (float64(sc.timeDecayDays) - daysSincePublish) / float64(sc.timeDecayDays)
		return sc.timeDecayBonus * decayRatio
	}

	return 0
}

// SetWeights 设置权重参数
func (sc *ScoreCalculator) SetWeights(download, favorite, rating float64) {
	sc.downloadWeight = download
	sc.favoriteWeight = favorite
	sc.ratingWeight = rating
}

// SetTimeDecay 设置时间衰减参数
func (sc *ScoreCalculator) SetTimeDecay(days int, bonus float64) {
	sc.timeDecayDays = days
	sc.timeDecayBonus = bonus
}

// GetWeights 获取当前权重配置
func (sc *ScoreCalculator) GetWeights() map[string]float64 {
	return map[string]float64{
		"download": sc.downloadWeight,
		"favorite": sc.favoriteWeight,
		"rating":   sc.ratingWeight,
	}
}
