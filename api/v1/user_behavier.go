package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetSearchHistoryReq 获取搜索历史请求
type GetSearchHistoryReq struct {
	g.Meta `path:"/games/search-history" method:"get" tags:"Game Management/User Behavior" summary:"获取搜索历史"`
	model.PageReq
}

// GetSearchHistoryRes 获取搜索历史响应
type GetSearchHistoryRes struct {
	g.Meta `mime:"application/json"`
	List   []*SearchHistoryItem `json:"list" dc:"搜索历史列表"`
	*model.PageRes
}

// SearchHistoryItem 搜索历史项
type SearchHistoryItem struct {
	ID            int64       `json:"id" dc:"ID"`
	UserID        int64       `json:"user_id" dc:"用户ID"`
	GameID        int64       `json:"game_id" dc:"游戏ID"`
	BehaviorType  string      `json:"behavior_type" dc:"行为类型"`
	SearchKeyword string      `json:"search_keyword" dc:"搜索关键词"`
	SearchTime    *gtime.Time `json:"search_time" dc:"搜索时间"`
}

// ClearSearchHistoryReq 清空搜索历史请求
type ClearSearchHistoryReq struct {
	g.Meta `path:"/games/search-history" method:"delete" tags:"Game Management/User Behavior" summary:"清空搜索历史"`
}

// ClearSearchHistoryRes 清空搜索历史响应
type ClearSearchHistoryRes struct {
	g.Meta `mime:"application/json"`
}

// GetGameHistoryReq 获取游戏历史请求
type GetGameHistoryReq struct {
	g.Meta `path:"/game-history" method:"get" tags:"游戏历史" summary:"获取游戏历史"`
	UserID int64 `json:"user_id" v:"required" dc:"用户ID"`
	Limit  int   `json:"limit" dc:"返回数量限制，默认50"`
}

// GetGameHistoryRes 获取游戏历史响应
type GetGameHistoryRes struct {
	List  []*GameHistoryItem `json:"list" dc:"游戏历史列表"`
	Total int                `json:"total" dc:"总数"`
}

// GameHistoryItem 游戏历史项
type GameHistoryItem struct {
	ID           int64  `json:"id" dc:"ID"`
	GameID       int64  `json:"game_id" dc:"游戏ID"`
	GameName     string `json:"game_name" dc:"游戏名称"`
	BehaviorType int    `json:"behavior_type" dc:"行为类型"`
	BehaviorTime string `json:"behavior_time" dc:"行为时间"`
	IPAddress    string `json:"ip_address" dc:"IP地址"`
}
