package service

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

// IGame 游戏服务接口
type IGame interface {
	// 游戏基础管理
	CreateGame(ctx context.Context, in *v1.CreateGameReq) (id int64, err error)
	DeleteGame(ctx context.Context, id int64) (err error)
	UpdateGame(ctx context.Context, in *v1.UpdateGameReq) (err error)
	GetGameByID(ctx context.Context, id int64) (out *model.Game, err error)
	GetGamesByIDs(ctx context.Context, ids []int64) (out []*model.Game, err error)
	ListGame(ctx context.Context, in *v1.ListGameReq) (outs []*model.Game, pageRes *model.PageRes, err error)
	AssertExists(ctx context.Context, id int64) (err error)

	// 游戏状态流转
	// 提交审核
	SubmitForReview(ctx context.Context, id int64) (err error)
	// 获取待审核的游戏
	ListInReview(ctx context.Context, pageReq *model.PageReq) (out []*model.Game, pageRes *model.PageRes, err error)
	// 游戏审核
	// 审核中 -> 审核通过/审核不通过
	Approve(ctx context.Context, id int64) (err error)
	Reject(ctx context.Context, id int64) (err error)
	// 发布游戏/游戏预约发布
	// 审核通过 -> 可预约/已发布
	PublishGameImmediately(ctx context.Context, id int64) (err error)
	PreRegisterGame(ctx context.Context, id int64, publishTime *gtime.Time) (err error)
	// 下架游戏
	UnpublishGame(ctx context.Context, id int64, unpublishReason string) (err error)

	// 事件驱动的状态流转
	HandleGameEvent(ctx context.Context, gameID int64, event model.GameEvent, data interface{}) error
	CancelPreRegisterGame(ctx context.Context, gameID int64) error
	UpdateGameInfo(ctx context.Context, gameID int64) error
	UpdateGameVersion(ctx context.Context, gameID int64) error

	HandleGameAutoPublish(ctx context.Context, task *model.AsyncTask) error
	NotifyReservedUsers(ctx context.Context, task *model.AsyncTask) error

	// 批量状态更新（定时任务用）
	BatchUpdateGameStatus(ctx context.Context) error

	// 游戏搜索
	SearchGameByGameName(ctx context.Context, name string, page, size int) (out []*model.Game, pageRes *model.PageRes, err error)

	// 游戏媒体管理
	AddMediaInfo(ctx context.Context, mediaInfo *model.GameMediaInfo) (err error)
	GetMediaInfo(ctx context.Context, gameID int64) (out []*model.GameMediaInfo, err error)
	UpdateMediaInfoByGameID(ctx context.Context, gameID int64, mediaInfos []*model.GameMediaInfo) (err error)
	UpdateMediaInfoStatusByFileID(ctx context.Context, fileID string, status model.GameMediaStatus) (err error)
	// 设置H5链接
	SetH5Link(ctx context.Context, gameID int64, link string) (err error)
	// 游戏提交审核时，检查必要的媒体文件是否上传
	CheckMediaInfo(ctx context.Context, gameInfo *model.Game) (err error)

	// 游戏收藏
	AddFavorite(ctx context.Context, gameID, userID int64) error
	RemoveFavorite(ctx context.Context, gameID, userID int64) error
	GetUserFavorites(ctx context.Context, userID int64, pageReq *model.PageReq) (out []*model.Game, pageRes *model.PageRes, err error)
	IsUserFavorited(ctx context.Context, gameID, userID int64) (bool, error)

	// 游戏评分
	AddRating(ctx context.Context, gameID, userID int64, rating int) error

	// 游戏下载
	Download(ctx context.Context, gameID, userID int64) error
}

var localGame IGame

func Game() IGame {
	if localGame == nil {
		panic("implement not found for interface IGame, forgot register?")
	}
	return localGame
}

func RegisterGame(i IGame) {
	localGame = i
}
