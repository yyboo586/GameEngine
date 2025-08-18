package v1

import "github.com/gogf/gf/v2/frame/g"

// GetSearchHistoryReq 获取搜索历史请求
type GetSearchHistoryReq struct {
	g.Meta `path:"/search-history" method:"get" tags:"搜索历史" summary:"获取搜索历史"`
	UserID int64 `json:"user_id" v:"required" dc:"用户ID"`
	Limit  int   `json:"limit" dc:"返回数量限制，默认20"`
}

// GetSearchHistoryRes 获取搜索历史响应
type GetSearchHistoryRes struct {
	List  []*SearchHistoryItem `json:"list" dc:"搜索历史列表"`
	Total int                  `json:"total" dc:"总数"`
}

// SearchHistoryItem 搜索历史项
type SearchHistoryItem struct {
	ID            int64  `json:"id" dc:"ID"`
	SearchKeyword string `json:"search_keyword" dc:"搜索关键词"`
	SearchTime    string `json:"search_time" dc:"搜索时间"`
	ResultCount   int    `json:"result_count" dc:"搜索结果数量"`
}

// ClearSearchHistoryReq 清空搜索历史请求
type ClearSearchHistoryReq struct {
	g.Meta `path:"/search-history" method:"delete" tags:"搜索历史" summary:"清空搜索历史"`
	UserID int64 `json:"user_id" v:"required" dc:"用户ID"`
}

// ClearSearchHistoryRes 清空搜索历史响应
type ClearSearchHistoryRes struct {
	Success bool   `json:"success" dc:"是否成功"`
	Message string `json:"message" dc:"响应消息"`
}

// GetPopularKeywordsReq 获取热门搜索关键词请求
type GetPopularKeywordsReq struct {
	g.Meta `path:"/popular-keywords" method:"get" tags:"搜索历史" summary:"获取热门搜索关键词"`
	Limit  int `json:"limit" dc:"返回数量限制，默认10"`
}

// GetPopularKeywordsRes 获取热门搜索关键词响应
type GetPopularKeywordsRes struct {
	Keywords []string `json:"keywords" dc:"热门搜索关键词列表"`
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

// RecordGameBehaviorReq 记录游戏行为请求
type RecordGameBehaviorReq struct {
	g.Meta       `path:"/game-behavior" method:"post" tags:"游戏行为" summary:"记录游戏行为"`
	UserID       int64  `json:"user_id" v:"required" dc:"用户ID"`
	GameID       int64  `json:"game_id" v:"required" dc:"游戏ID"`
	BehaviorType int    `json:"behavior_type" v:"required" dc:"行为类型(1:查看 2:下载 3:收藏 4:评分)"`
	IPAddress    string `json:"ip_address" dc:"IP地址"`
}

// RecordGameBehaviorRes 记录游戏行为响应
type RecordGameBehaviorRes struct {
	Success bool   `json:"success" dc:"是否成功"`
	Message string `json:"message" dc:"响应消息"`
}
