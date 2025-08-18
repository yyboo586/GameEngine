package entity

import "github.com/gogf/gf/v2/os/gtime"

type GameReserve struct {
	ID         int64       `orm:"id" dc:"ID"`
	GameID     int64       `orm:"game_id" dc:"游戏ID"`
	UserID     int64       `orm:"user_id" dc:"用户ID"`
	CreateTime *gtime.Time `orm:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `orm:"update_time" dc:"更新时间"`
}
