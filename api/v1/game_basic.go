package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CreateGameReq struct {
	g.Meta `path:"/games" method:"post" tags:"Game Management" summary:"Create Game"`
	model.AuthorRequired
	Name           string  `json:"name" v:"required|length:1,30#游戏名称不能为空|游戏名称长度不能超过30个字符" dc:"游戏名称"`
	DistributeType int     `json:"distribute_type" v:"required#游戏分发类型不能为空" dc:"游戏分发类型(1:APK,2:H5)"`
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
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type DeleteGameRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateGameReq struct {
	g.Meta `path:"/games/{id}" method:"put" tags:"Game Management" summary:"Update Game"`
	model.AuthorRequired
	ID             int64   `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	Name           string  `json:"name" v:"length:1,30#游戏名称长度不能超过30个字符" dc:"游戏名称"`
	DistributeType int     `json:"distribute_type" v:"in:1,2#游戏分发类型必须是1,2" dc:"游戏分发类型(1:APK,2:H5)"`
	CategoryID     int64   `json:"category_id" dc:"游戏分类"`
	TagIDs         []int64 `json:"tag_ids" dc:"游戏标签"`
	Developer      string  `json:"developer" v:"length:1,30#游戏开发者长度不能超过30个字符" dc:"游戏开发者"`
	Publisher      string  `json:"publisher" v:"length:1,30#游戏发行商长度不能超过30个字符" dc:"游戏发行商"`
	Description    string  `json:"description" v:"length:1,200#游戏描述长度不能超过200个字符" dc:"游戏基本描述"`
	Details        string  `json:"details" v:"length:1,500#游戏详情长度不能超过500个字符" dc:"游戏详情"`
}

type UpdateGameRes struct {
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
	Name string `json:"name" dc:"游戏名称"`
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

type Game struct {
	ID             int64         `json:"id" dc:"游戏ID"`
	Name           string        `json:"name" dc:"游戏名称"`
	DistributeType string        `json:"distribute_type" dc:"游戏分发类型"`
	Category       *CategoryInfo `json:"category" dc:"游戏分类"`
	Tags           []*TagInfo    `json:"tags" dc:"游戏标签"`
	Developer      string        `json:"developer" dc:"游戏开发者"`
	Publisher      string        `json:"publisher" dc:"游戏发行商"`
	Description    string        `json:"description" dc:"游戏描述"`
	Details        string        `json:"details" dc:"游戏详情"`

	Status       string      `json:"status" dc:"游戏状态"`
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

	IsFavorite bool `json:"is_favorite" dc:"是否收藏"`
	IsReserve  bool `json:"is_reserve" dc:"是否预约"`
}
