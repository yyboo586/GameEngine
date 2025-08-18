package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 推荐相关API结构体

// GetTodayPicksReq 获取今日精选请求
type GetTodayPicksReq struct {
	g.Meta `path:"/recommendations/today" method:"get" tags:"游戏推荐" summary:"获取今日精选"`
	model.PageReq
}

// GetTodayPicksRes 获取今日精选响应
type GetTodayPicksRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

// GetSimilarGamesReq 获取相似游戏请求
type GetSimilarGamesReq struct {
	g.Meta `path:"/games/{id}/similar" method:"get" tags:"游戏推荐" summary:"获取相似游戏"`
	ID     int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	model.PageReq
}

// GetSimilarGamesRes 获取相似游戏响应
type GetSimilarGamesRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"相似游戏列表"`
	*model.PageRes
}

// GetPersonalizedRecommendationsReq 获取个性化推荐请求
type GetPersonalizedRecommendationsReq struct {
	g.Meta `path:"/recommendations/personalized" method:"get" tags:"游戏推荐" summary:"获取个性化推荐"`
	UserID int64 `q:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	model.PageReq
}

// GetPersonalizedRecommendationsRes 获取个性化推荐响应
type GetPersonalizedRecommendationsRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

// GetRecommendationsByCategoryReq 基于分类的推荐请求
type GetRecommendationsByCategoryReq struct {
	g.Meta     `path:"/recommendations/category/{category_id}" method:"get" tags:"游戏推荐" summary:"基于分类的推荐"`
	CategoryID int64 `p:"category_id" v:"required#分类ID不能为空" dc:"分类ID"`
	model.PageReq
}

// GetRecommendationsByCategoryRes 基于分类的推荐响应
type GetRecommendationsByCategoryRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

// GetRecommendationsByTagsReq 基于标签的推荐请求
type GetRecommendationsByTagsReq struct {
	g.Meta `path:"/recommendations/tags" method:"get" tags:"游戏推荐" summary:"基于标签的推荐"`
	TagIDs []int64 `q:"tag_ids" v:"required#标签ID不能为空" dc:"标签ID列表"`
	model.PageReq
}

// GetRecommendationsByTagsRes 基于标签的推荐响应
type GetRecommendationsByTagsRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

// GetPopularRecommendationsReq 获取热门推荐请求
type GetPopularRecommendationsReq struct {
	g.Meta `path:"/recommendations/popular" method:"get" tags:"游戏推荐" summary:"获取热门推荐"`
	model.PageReq
}

// GetPopularRecommendationsRes 获取热门推荐响应
type GetPopularRecommendationsRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

// GetNewGameRecommendationsReq 获取新游推荐请求
type GetNewGameRecommendationsReq struct {
	g.Meta `path:"/recommendations/new" method:"get" tags:"游戏推荐" summary:"获取新游推荐"`
	model.PageReq
}

// GetNewGameRecommendationsRes 获取新游推荐响应
type GetNewGameRecommendationsRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"推荐游戏列表"`
	*model.PageRes
}

type TodayPickItem struct {
	GameID         int64   `json:"game_id" dc:"游戏ID"`
	Game           *Game   `json:"game" dc:"游戏"`
	HotScore       float64 `json:"hot_score" dc:"热度分数"`
	IsNewRelease   bool    `json:"is_new_release" dc:"是否为今日新发布"`
	IsEditorChoice bool    `json:"is_editor_choice" dc:"是否为编辑推荐"`
}

type PersonalizedRecommendationItem struct {
	GameID int64 `json:"game_id" dc:"游戏ID"`
	Game   *Game `json:"game" dc:"游戏"`
}
