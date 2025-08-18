package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 新增的API结构，用于重构后的服务接口
type CreateGameReq struct {
	g.Meta `path:"/games" method:"post" tags:"Game Management" summary:"Create Game"`
	model.Author
	Name           string  `json:"name" v:"required|length:1,30#游戏名称不能为空|游戏名称长度不能超过30个字符" dc:"游戏名称"`
	DistributeType int     `json:"distribute_type" v:"required#游戏分发类型不能为空" dc:"游戏分发类型(1:APK,2:链接)"`
	CategoryID     int64   `json:"category_id" v:"required#游戏分类不能为空" dc:"游戏分类"`
	TagIDs         []int64 `json:"tag_ids" dc:"游戏标签"`
	Developer      string  `json:"developer" v:"required#游戏开发者不能为空" dc:"游戏开发者"`
	Publisher      string  `json:"publisher" v:"required#游戏发行商不能为空" dc:"游戏发行商"`
	Description    string  `json:"description" v:"required#游戏描述不能为空" dc:"游戏基本描述"`
	Details        string  `json:"details" v:"required#游戏详情不能为空" dc:"游戏详情"`
}

type CreateGameRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"游戏ID"`
}

type DeleteGameReq struct {
	g.Meta `path:"/games/{id}" method:"delete" tags:"Game Management" summary:"Delete Game"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type DeleteGameRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateGameReq struct {
	g.Meta `path:"/games/{id}" method:"put" tags:"Game Management" summary:"Update Game"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type UpdateGameRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateGamePublishStatusReq struct {
	g.Meta `path:"/games/{id}/publish_status" method:"put" tags:"Game Management" summary:"Update Game Publish Status"`
	model.Author
	ID     int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	Status int   `p:"status" v:"required|in:1,2,3#游戏状态不能为空|游戏状态必须是1,2,3" dc:"游戏状态(1:未上架 2:已上架 3:已下架)"`
}

type UpdateGamePublishStatusRes struct {
	g.Meta `mime:"application/json"`
}

type GetGameByIDReq struct {
	g.Meta `path:"/games/{id}" method:"get" tags:"Game Management" summary:"Get Game Details"`
	ID     int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type GetGameByIDRes struct {
	g.Meta `mime:"application/json"`
	*Game  `json:"game" dc:"游戏详情"`
}

type ListGameReq struct {
	g.Meta `path:"/games" method:"get" tags:"Game Management" summary:"List Game"`
	model.PageReq
}

type ListGameRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	*model.PageRes
}

type SearchGameByGameNameReq struct {
	g.Meta `path:"/games/search-by-name" method:"get" tags:"Game Management" summary:"Search Game By Name"`
	model.PageReq
	Name string `json:"name" dc:"游戏名称"`
}

type SearchGameByGameNameRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
}

type ListGamesByCategoryNameReq struct {
	g.Meta       `path:"/games/search-by-category-name" method:"get" tags:"Game Management" summary:"List Game By Category Name"`
	CategoryName string `p:"category_name" v:"required#游戏分类名称不能为空" dc:"游戏分类名称"`
	model.PageReq
}

type ListGamesByCategoryNameRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	*model.PageRes
}

type ListGamesByTagNameReq struct {
	g.Meta  `path:"/games/search-by-tag-name" method:"get" tags:"Game Management" summary:"List Game By Tag Name"`
	TagName string `p:"tag_name" v:"required#游戏标签名称不能为空" dc:"游戏标签名称"`
	model.PageReq
}

type ListGamesByTagNameRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	*model.PageRes
}

type Game struct {
	ID             int64         `json:"id" dc:"游戏ID"`
	Name           string        `json:"name" dc:"游戏名称"`
	DistributeType int           `json:"distribute_type" dc:"游戏分发类型"`
	Category       *CategoryInfo `json:"category" dc:"游戏分类"`
	Tags           []*TagInfo    `json:"tags" dc:"游戏标签"`
	Developer      string        `json:"developer" dc:"游戏开发者"`
	Publisher      string        `json:"publisher" dc:"游戏发行商"`
	Description    string        `json:"description" dc:"游戏描述"`
	Details        string        `json:"details" dc:"游戏详情"`

	Status       int         `json:"status" dc:"游戏状态(1:未上架 2:已上架 3:已下架)"`
	PublishTime  *gtime.Time `json:"publish_time" dc:"发布时间"`
	ReserveCount int64       `json:"reserve_count" dc:"预约次数"`

	AverageRating float64 `json:"average_rating" dc:"游戏平均评分"`
	RatingScore   int64   `json:"rating_score" dc:"游戏评分总分"`
	RatingCount   int64   `json:"rating_count" dc:"游戏评分次数"`
	FavoriteCount int64   `json:"favorite_count" dc:"游戏收藏次数"`
	DownloadCount int64   `json:"download_count" dc:"游戏下载次数"`

	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`

	MediaInfos []*GameMediaInfo `json:"media_infos" dc:"游戏媒体信息"`
}
