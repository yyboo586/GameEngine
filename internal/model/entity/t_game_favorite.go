package entity

import "github.com/gogf/gf/v2/os/gtime"

// GameFavorite 游戏收藏实体
type GameFavorite struct {
	ID         int64       `json:"id" dc:"主键"`
	GameID     int64       `json:"game_id" dc:"游戏ID"`
	UserID     int64       `json:"user_id" dc:"用户ID"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}
