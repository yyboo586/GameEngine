package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// GetHotGamesReq 获取热门游戏榜单请求
type GetHotGamesReq struct {
	g.Meta `path:"/games/ranking/hot" method:"get" tags:"Game Management/Ranking" summary:"Get Hot Games"`
	model.PageReq
}

// GetHotGamesRes 获取热门游戏榜单响应
type GetHotGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetThisMonthNewGamesReq 获取本月新游戏请求
type GetThisMonthNewGamesReq struct {
	g.Meta `path:"/games/ranking/this-month-new" method:"get" tags:"Game Management/Ranking" summary:"Get This Month New Games"`
	model.PageReq
}

// GetThisMonthNewGamesRes 获取本月新游戏响应
type GetThisMonthNewGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// GetUpcomingGamesReq 获取即将上新游戏请求
type GetUpcomingGamesReq struct {
	g.Meta `path:"/games/ranking/upcoming" method:"get" tags:"Game Management/Ranking" summary:"Get Upcoming Games"`
	model.PageReq
}

// GetUpcomingGamesRes 获取即将上新游戏响应
type GetUpcomingGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// GetCategoryRankingReq 获取分类榜单请求
type GetCategoryRankingReq struct {
	g.Meta     `path:"/games/ranking/category/{category_id}" method:"get" tags:"Game Management/Ranking" summary:"Get Category Ranking"`
	CategoryID int64 `p:"category_id" v:"required#分类ID不能为空" dc:"分类ID"`
	model.PageReq
}

// GetCategoryRankingRes 获取分类榜单响应
type GetCategoryRankingRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetTagRankingReq 获取标签榜单请求
type GetTagRankingReq struct {
	g.Meta `path:"/games/ranking/tag/{tag_id}" method:"get" tags:"Game Management/Ranking" summary:"Get Tag Ranking"`
	TagID  int64 `p:"tag_id" v:"required#标签ID不能为空" dc:"标签ID"`
	model.PageReq
}

// GetTagRankingRes 获取标签榜单响应
type GetTagRankingRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetTodayRecommendReq 获取今日推荐请求
type GetTodayRecommendReq struct {
	g.Meta `path:"/games/ranking/today-recommend" method:"get" tags:"Game Management/Ranking" summary:"Get Today Recommend"`
	model.PageReq
}

// GetTodayRecommendRes 获取今日推荐响应
type GetTodayRecommendRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetTopRatedGamesReq 获取高分游戏榜单请求
type GetTopRatedGamesReq struct {
	g.Meta `path:"/games/ranking/top-rated" method:"get" tags:"Game Management/Ranking" summary:"Get Top Rated Games"`
	model.PageReq
}

// GetTopRatedGamesRes 获取高分游戏榜单响应
type GetTopRatedGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetMostDownloadedGamesReq 获取下载量榜单请求
type GetMostDownloadedGamesReq struct {
	g.Meta `path:"/games/ranking/most-downloaded" method:"get" tags:"Game Management/Ranking" summary:"Get Most Downloaded Games"`
	model.PageReq
}

// GetMostDownloadedGamesRes 获取下载量榜单响应
type GetMostDownloadedGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetMostFavoritedGamesReq 获取收藏数榜单请求
type GetMostFavoritedGamesReq struct {
	g.Meta `path:"/games/ranking/most-favorited" method:"get" tags:"Game Management/Ranking" summary:"Get Most Favorited Games"`
	model.PageReq
}

// GetMostFavoritedGamesRes 获取收藏数榜单响应
type GetMostFavoritedGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes
}

// GetRelatedGamesReq 获取相关游戏推荐请求
type GetRelatedGamesReq struct {
	g.Meta `path:"/games/ranking/related" method:"get" tags:"Game Management/Ranking" summary:"Get Related Games Recommendation"`
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	model.PageReq
}

// GetRelatedGamesRes 获取相关游戏推荐响应
type GetRelatedGamesRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"相关游戏列表"`
	*model.PageRes
}
