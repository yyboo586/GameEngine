package entity

import "github.com/gogf/gf/v2/os/gtime"

type SearchHistory struct {
	ID            int64       `orm:"id" dc:"ID"`
	UserID        int64       `orm:"user_id" dc:"用户ID"`
	SearchKeyword string      `orm:"search_keyword" dc:"搜索关键词"`
	SearchTime    *gtime.Time `orm:"search_time" dc:"搜索时间"`
	ResultCount   int         `orm:"result_count" dc:"搜索结果数量"`
}
