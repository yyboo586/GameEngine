package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameRating 游戏评分
type GameRating struct {
	ID         int64       `json:"id" dc:"评分ID"`
	GameID     int64       `json:"game_id" dc:"游戏ID"`
	UserID     int64       `json:"user_id" dc:"用户ID"`
	Score      int         `json:"score" dc:"评分(1-5)"`
	Comment    string      `json:"comment" dc:"评论"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}

// GameRatingSummary 游戏评分汇总
type GameRatingSummary struct {
	GameID       int64   `json:"game_id" dc:"游戏ID"`
	TotalScore   int64   `json:"total_score" dc:"总评分"`
	TotalCount   int64   `json:"total_count" dc:"评分次数"`
	AverageScore float64 `json:"average_score" dc:"平均评分"`
	Score1Count  int64   `json:"score_1_count" dc:"1分次数"`
	Score2Count  int64   `json:"score_2_count" dc:"2分次数"`
	Score3Count  int64   `json:"score_3_count" dc:"3分次数"`
	Score4Count  int64   `json:"score_4_count" dc:"4分次数"`
	Score5Count  int64   `json:"score_5_count" dc:"5分次数"`
}
