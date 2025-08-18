package recommendation

import (
	"GameEngine/internal/model"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
)

// HotScoreCalculator 热度分数计算器
type HotScoreCalculator struct{}

// NewHotScoreCalculator 创建热度分数计算器实例
func NewHotScoreCalculator() *HotScoreCalculator {
	return &HotScoreCalculator{}
}

// CalculateHotScore 计算热度分数
func (hsc *HotScoreCalculator) CalculateHotScore(game *model.Game) float64 {
	// 基础分数：下载量×0.4 + 收藏数×0.3 + 评分×0.2 + 评分次数×0.1
	baseScore := float64(game.DownloadCount)*0.4 +
		float64(game.FavoriteCount)*0.3 +
		float64(game.RatingScore)*0.2 +
		float64(game.RatingCount)*0.1

	// 时间衰减因子：最近发布的游戏获得额外加分
	timeBonus := hsc.calculateTimeBonus(game.PublishTime)

	return baseScore + timeBonus
}

// calculateTimeBonus 计算时间加分
func (hsc *HotScoreCalculator) calculateTimeBonus(publishTime *gtime.Time) float64 {
	if publishTime == nil {
		return 0.0
	}

	daysSincePublish := time.Since(publishTime.Time).Hours() / 24

	// 7天内线性衰减
	if daysSincePublish <= 7 {
		return 100 * (7 - daysSincePublish) / 7
	}

	// 30天内缓慢衰减
	if daysSincePublish <= 30 {
		return 20 * (30 - daysSincePublish) / 23 // 23 = 30-7
	}

	// 超过30天无加分
	return 0.0
}

// CalculateWeightedScore 计算加权分数
func (hsc *HotScoreCalculator) CalculateWeightedScore(game *model.Game, weights *ScoreWeights) float64 {
	if weights == nil {
		weights = &ScoreWeights{
			DownloadWeight: 0.4,
			FavoriteWeight: 0.3,
			RatingWeight:   0.2,
			CountWeight:    0.1,
		}
	}

	baseScore := float64(game.DownloadCount)*weights.DownloadWeight +
		float64(game.FavoriteCount)*weights.FavoriteWeight +
		float64(game.RatingScore)*weights.RatingWeight +
		float64(game.RatingCount)*weights.CountWeight

	timeBonus := hsc.calculateTimeBonus(game.PublishTime)

	return baseScore + timeBonus
}

// ScoreWeights 分数权重配置
type ScoreWeights struct {
	DownloadWeight float64 `json:"download_weight" dc:"下载量权重"`
	FavoriteWeight float64 `json:"favorite_weight" dc:"收藏数权重"`
	RatingWeight   float64 `json:"rating_weight" dc:"评分权重"`
	CountWeight    float64 `json:"count_weight" dc:"评分次数权重"`
}

// GetDefaultWeights 获取默认权重
func (hsc *HotScoreCalculator) GetDefaultWeights() *ScoreWeights {
	return &ScoreWeights{
		DownloadWeight: 0.4,
		FavoriteWeight: 0.3,
		RatingWeight:   0.2,
		CountWeight:    0.1,
	}
}

// GetHotGameWeights 获取热门游戏权重（更重视下载量）
func (hsc *HotScoreCalculator) GetHotGameWeights() *ScoreWeights {
	return &ScoreWeights{
		DownloadWeight: 0.6,
		FavoriteWeight: 0.2,
		RatingWeight:   0.15,
		CountWeight:    0.05,
	}
}

// GetQualityGameWeights 获取高质量游戏权重（更重视评分）
func (hsc *HotScoreCalculator) GetQualityGameWeights() *ScoreWeights {
	return &ScoreWeights{
		DownloadWeight: 0.2,
		FavoriteWeight: 0.3,
		RatingWeight:   0.4,
		CountWeight:    0.1,
	}
}
