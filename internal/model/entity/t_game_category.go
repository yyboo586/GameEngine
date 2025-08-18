package entity

import "github.com/gogf/gf/v2/os/gtime"

type GameCategory struct {
	ID         int64       `orm:"id"`
	GameID     int64       `orm:"game_id"`
	CategoryID int64       `orm:"category_id"`
	CreateTime *gtime.Time `orm:"create_time"`
	UpdateTime *gtime.Time `orm:"update_time"`
}
