package entity

import "github.com/gogf/gf/v2/os/gtime"

type GameBehavior struct {
	ID           int64       `orm:"id" dc:"ID"`
	UserID       int64       `orm:"user_id" dc:"用户ID"`
	GameID       int64       `orm:"game_id" dc:"游戏ID"`
	BehaviorType int         `orm:"behavior_type" dc:"行为类型(1:查看 2:下载 3:收藏 4:评分)"`
	BehaviorTime *gtime.Time `orm:"behavior_time" dc:"行为时间"`
	IPAddress    string      `orm:"ip_address" dc:"IP地址"`
}
