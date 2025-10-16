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
	GameMediaTypeH5Link                   // 链接
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
	GameStatusInit        GameStatus = iota // 初始状态
	GameStatusInReview                      // 审核中
	GameStatusApproved                      // 审核通过
	GameStatusPreRegister                   // 可预约
	GameStatusPublished                     // 已上架
	GameStatusUnpublished                   // 已下架
)

type GameEvent int

const (
	SubmitForReview   GameEvent = iota // 提交审核
	Approve                            // 审核通过
	Reject                             // 审核失败
	PreRegister                        // 预约发布
	PublishNow                         // 立即发布
	UpdateInfo                         // 更新游戏信息
	CancelPreRegister                  // 取消预约发布
	UnpublishNow                       // 立即下架
	UpdateVersion                      // 更新游戏版本
	AutoPublish                        // 自动发布(预约时间到达)
)

// 获取事件名称
func GetGameEventText(event GameEvent) string {
	switch event {
	case SubmitForReview:
		return "提交审核"
	case Approve:
		return "审核通过"
	case Reject:
		return "审核失败"
	case PreRegister:
		return "预约发布"
	case PublishNow:
		return "立即发布"
	case UpdateInfo:
		return "更新游戏信息"
	case CancelPreRegister:
		return "取消预约发布"
	case UnpublishNow:
		return "立即下架"
	case UpdateVersion:
		return "更新游戏版本"
	case AutoPublish:
		return "自动发布"
	default:
		return "未知事件"
	}
}

// 获取状态名称
func GetGameStatusText(status GameStatus) string {
	switch status {
	case GameStatusInit:
		return "初始状态"
	case GameStatusInReview:
		return "审核中"
	case GameStatusApproved:
		return "审核通过"
	case GameStatusPreRegister:
		return "可预约"
	case GameStatusPublished:
		return "已上架"
	case GameStatusUnpublished:
		return "已下架"
	default:
		return "未知状态"
	}
}

func GetGameDistributeTypeText(distributeType GameDistributeType) string {
	switch distributeType {
	case GameDistributeTypeAPK:
		return "APK"
	case GameDistributeTypeLink:
		return "H5"
	}
	return "未知类型"
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

	Version    int         `json:"version" dc:"版本"`
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

		Version:    in.Version,
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
