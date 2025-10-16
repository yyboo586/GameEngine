package entity

import "github.com/gogf/gf/v2/os/gtime"

type Game struct {
	ID             int64  `orm:"id" dc:"ID"`
	Name           string `orm:"name" dc:"名称"`
	DistributeType int    `orm:"distribute_type" dc:"类型"`
	Developer      string `orm:"developer" dc:"开发商"`
	Publisher      string `orm:"publisher" dc:"发行商"`
	Description    string `orm:"description" dc:"描述"`
	Details        string `orm:"details" dc:"详情"`

	Status       int         `orm:"status" dc:"状态"`
	PublishTime  *gtime.Time `orm:"publish_time" dc:"发布时间"`
	ReserveCount int64       `orm:"reserve_count" dc:"预约次数"`

	RatingScore   int64 `orm:"rating_score" dc:"评分总分"`
	RatingCount   int64 `orm:"rating_count" dc:"评分次数"`
	FavoriteCount int64 `orm:"favorite_count" dc:"收藏次数"`
	DownloadCount int64 `orm:"download_count" dc:"下载次数"`

	Version    int         `orm:"version" dc:"版本"`
	CreateTime *gtime.Time `orm:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `orm:"update_time" dc:"更新时间"`
}
