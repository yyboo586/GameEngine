package model

import (
	"GameEngine/internal/model/entity"
	"math"

	"github.com/gogf/gf/v2/os/gtime"
)

type GameDistributeType int

const (
	_                      GameDistributeType = iota
	GameDistributeTypeAPK                     // APK文件
	GameDistributeTypeLink                    // 链接
)

type GameMediaType int

const (
	_                       GameMediaType = iota
	GameMediaTypeIcon                     // 图标
	GameMediaTypeScreenshot               // 截图
	GameMediaTypeVideo                    // 视频
	GameMediaTypeApkFile                  // APK文件
)

type GameMediaStatus int

const (
	_                      GameMediaStatus = iota
	GameMediaStatusInit                    // 初始化
	GameMediaStatusSuccess                 // 成功
	GameMediaStatusFailed                  // 失败
)

type GameStatus int

const (
	_                     GameStatus = iota
	GameStatusUnpublished            // 未上架
	GameStatusPublished              // 已上架
	GameStatusDeleted                // 已下架
)

type BehaviorType int

const (
	BehaviorTypeView     BehaviorType = 1 // 查看
	BehaviorTypeDownload BehaviorType = 2 // 下载
	BehaviorTypeFavorite BehaviorType = 3 // 收藏
	BehaviorTypeRating   BehaviorType = 4 // 评分
)

type SearchHistory struct {
	ID            int64       `json:"id" dc:"ID"`
	UserID        int64       `json:"user_id" dc:"用户ID"`
	SearchKeyword string      `json:"search_keyword" dc:"搜索关键词"`
	SearchTime    *gtime.Time `json:"search_time" dc:"搜索时间"`
	ResultCount   int         `json:"result_count" dc:"搜索结果数量"`
}

type GameBehavior struct {
	ID           int64        `json:"id" dc:"ID"`
	UserID       int64        `json:"user_id" dc:"用户ID"`
	GameID       int64        `json:"game_id" dc:"游戏ID"`
	BehaviorType BehaviorType `json:"behavior_type" dc:"行为类型"`
	BehaviorTime *gtime.Time  `json:"behavior_time" dc:"行为时间"`
	IPAddress    string       `json:"ip_address" dc:"IP地址"`
}

type Game struct {
	ID             int64              `json:"id" dc:"ID"`
	Name           string             `json:"name" dc:"名称"`
	DistributeType GameDistributeType `json:"distribute_type" dc:"分发类型"`
	Developer      string             `json:"developer" dc:"开发商"`
	Publisher      string             `json:"publisher" dc:"发行商"`
	Description    string             `json:"description" dc:"描述"`
	Details        string             `json:"details" dc:"详情"`

	Status       GameStatus  `json:"status" dc:"状态"`
	PublishTime  *gtime.Time `json:"publish_time" dc:"发布时间"`
	ReserveCount int64       `json:"reserve_count" dc:"预约次数"`

	RatingCount   int64   `json:"rating_count" dc:"评分次数"`
	RatingScore   int64   `json:"rating_score" dc:"评分总分"`
	AverageRating float64 `json:"average_rating" dc:"平均评分"`
	FavoriteCount int64   `json:"favorite_count" dc:"收藏次数"`
	DownloadCount int64   `json:"download_count" dc:"下载次数"`

	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}

type GameMediaInfo struct {
	ID         int64           `json:"id" dc:"ID"`
	GameID     int64           `json:"game_id" dc:"游戏ID"`
	FileID     string          `json:"file_id" dc:"文件ID"`
	Status     GameMediaStatus `json:"status" dc:"状态"`
	MediaType  GameMediaType   `json:"media_type" dc:"媒体类型"`
	MediaUrl   string          `json:"media_url" dc:"媒体URL"`
	CreateTime *gtime.Time     `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time     `json:"update_time" dc:"更新时间"`
}

func ConvertGameEntityToModel(in *entity.Game) (out *Game) {
	out = &Game{
		ID:             in.ID,
		Name:           in.Name,
		DistributeType: GameDistributeType(in.DistributeType),
		Developer:      in.Developer,
		Publisher:      in.Publisher,
		Description:    in.Description,
		Details:        in.Details,

		Status:       GameStatus(in.Status),
		PublishTime:  in.PublishTime,
		ReserveCount: in.ReserveCount,

		RatingScore:   in.RatingScore,
		RatingCount:   in.RatingCount,
		FavoriteCount: in.FavoriteCount,
		DownloadCount: in.DownloadCount,

		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}

	out.AverageRating = CalcRating(in.RatingScore, in.RatingCount)
	return
}

func CalcRating(totalRatingScore, totalRatingCount int64) (averageRating float64) {
	if totalRatingCount == 0 {
		return 0
	}
	averageRating = float64(totalRatingScore) / float64(totalRatingCount)
	averageRating = math.Round(averageRating*10) / 10
	return
}
