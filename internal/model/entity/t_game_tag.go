package entity

import "github.com/gogf/gf/v2/os/gtime"

type GameTag struct {
	ID         int64       `orm:"id"`
	GameID     int64       `orm:"game_id"`
	TagID      int64       `orm:"tag_id"`
	CreateTime *gtime.Time `orm:"create_time"`
	UpdateTime *gtime.Time `orm:"update_time"`
}
