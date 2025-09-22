package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GetSearchHistoryReq 获取搜索历史请求
type GetSearchHistoryReq struct {
	g.Meta `path:"/games/search-history" method:"get" tags:"Game Management/User Behavior" summary:"Game Search History"`
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
	g.Meta `path:"/games/search-history" method:"delete" tags:"Game Management/User Behavior" summary:"Clear Game Search History"`
}

// ClearSearchHistoryRes 清空搜索历史响应
type ClearSearchHistoryRes struct {
	g.Meta `mime:"application/json"`
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

// PlayGameReq 玩游戏请求
type PlayGameReq struct {
	g.Meta `path:"/games/{game_id}/play" method:"post" tags:"Game Management/User Behavior" summary:"Play Game"`
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// PlayGameRes 玩游戏响应
type PlayGameRes struct {
	g.Meta `mime:"application/json"`
}

// GetPlayHistoryReq 获取玩过游戏历史记录请求
type GetPlayHistoryReq struct {
	g.Meta `path:"/games/play-history" method:"get" tags:"Game Management/User Behavior" summary:"Get Play History"`
	model.PageReq
}

// GetPlayHistoryRes 获取玩过游戏历史记录响应
type GetPlayHistoryRes struct {
	g.Meta  `mime:"application/json"`
	List    []*PlayHistoryItem `json:"list" dc:"玩过游戏历史列表"`
	PageRes *model.PageRes     `json:"page_res" dc:"分页信息"`
}

// PlayHistoryItem 玩过游戏历史项
type PlayHistoryItem struct {
	*Game
	PlayTime  *gtime.Time `json:"play_time" dc:"游玩时间"`
	IPAddress string      `json:"ip_address" dc:"IP地址"`
}
