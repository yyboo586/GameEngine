package service

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"context"
)

// IGame 游戏服务接口
type IGame interface {
	// 游戏基础管理
	CreateGame(ctx context.Context, in *v1.CreateGameReq) (id int64, err error)
	DeleteGame(ctx context.Context, id int64) (err error)
	UpdateGame(ctx context.Context, in *v1.UpdateGameReq) (err error)
	UpdateGamePublishStatus(ctx context.Context, id int64, status int) (err error)
	GetGameByID(ctx context.Context, id int64) (out *model.Game, err error)
	GetGamesByIDs(ctx context.Context, ids []int64) (out []*model.Game, err error)
	ListGame(ctx context.Context, in *v1.ListGameReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 游戏搜索
	SearchGameByGameName(ctx context.Context, name string, page, size int) (out []*model.Game, pageRes *model.PageRes, err error)

	// 游戏媒体管理
	AddMediaInfo(ctx context.Context, mediaInfo *model.GameMediaInfo) (err error)
	GetMediaInfo(ctx context.Context, gameID int64) (out []*model.GameMediaInfo, err error)
	UpdateMediaInfoByGameID(ctx context.Context, gameID int64, mediaInfos []*model.GameMediaInfo) (err error)
	UpdateMediaInfoStatusByFileID(ctx context.Context, fileID string, status model.GameMediaStatus) (err error)

	// 游戏收藏
	AddFavorite(ctx context.Context, gameID, userID int64) error
	RemoveFavorite(ctx context.Context, gameID, userID int64) error
	GetUserFavorites(ctx context.Context, userID int64, pageReq *model.PageReq) (out []*model.Game, pageRes *model.PageRes, err error)

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
