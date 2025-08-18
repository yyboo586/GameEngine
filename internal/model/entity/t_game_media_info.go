package entity

import "github.com/gogf/gf/v2/os/gtime"

// GameMediaInfo 游戏媒体信息实体
type GameMediaInfo struct {
	ID         int64       `orm:"id" dc:"主键"`
	GameID     int64       `orm:"game_id" dc:"游戏ID"`
	FileID     string      `orm:"file_id" dc:"文件ID"`
	MediaType  int         `orm:"media_type" dc:"媒体类型"`
	MediaUrl   string      `orm:"media_url" dc:"媒体URL"`
	Status     int         `orm:"status" dc:"状态"`
	CreateTime *gtime.Time `orm:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `orm:"update_time" dc:"更新时间"`
}
