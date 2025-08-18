package entity

import "github.com/gogf/gf/v2/os/gtime"

type UserBehavior struct {
	ID            int64       `orm:"id" dc:"ID"`
	UserID        int64       `orm:"user_id" dc:"用户ID"`
	GameID        int64       `orm:"game_id" dc:"游戏ID"`
	BehaviorType  int         `orm:"behavior_type" dc:"行为类型"`
	SearchKeyword string      `orm:"search_keyword" dc:"搜索关键词"`
	BehaviorTime  *gtime.Time `orm:"behavior_time" dc:"行为时间"`
	IPAddress     string      `orm:"ip_address" dc:"IP地址"`
}
