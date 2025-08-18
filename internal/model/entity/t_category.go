package entity

import "github.com/gogf/gf/v2/os/gtime"

type Category struct {
	ID         int64       `orm:"id"`
	Name       string      `orm:"name"`
	CreateTime *gtime.Time `orm:"create_time"`
	UpdateTime *gtime.Time `orm:"update_time"`
}
