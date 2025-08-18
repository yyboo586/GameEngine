package service

import (
	"GameEngine/internal/model"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

// IMetadata 元数据管理服务接口
type IMetadata interface {
	// 分类管理
	CreateCategory(ctx context.Context, name string) (id int64, err error)
	UpdateCategory(ctx context.Context, id int64, name string) (err error)
	DeleteCategory(ctx context.Context, id int64) (err error)
	GetCategoryByID(ctx context.Context, id int64) (out *model.Category, err error)
	ListCategory(ctx context.Context) (outs []*model.Category, err error)

	// 标签管理
	CreateTag(ctx context.Context, name string) (id int64, err error)
	UpdateTag(ctx context.Context, id int64, name string) (err error)
	DeleteTag(ctx context.Context, id int64) (err error)
	GetTagByID(ctx context.Context, id int64) (out *model.Tag, err error)
	ListTag(ctx context.Context) (outs []*model.Tag, err error)

	// 游戏分类关联
	AddGameCategory(ctx context.Context, tx gdb.TX, gameID, categoryID int64) error
	RemoveGameCategory(ctx context.Context, tx gdb.TX, gameID int64) error
	GetGameIDsByCategoryName(ctx context.Context, categoryName string) (out []int64, err error)
	GetCategoryByGameID(ctx context.Context, gameID int64) (out *model.Category, err error)
	//GetGameCategories(ctx context.Context, gameID int64) ([]*model.Category, error)
	//GetCategoryGames(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)

	// 游戏标签关联
	AddGameTags(ctx context.Context, tx gdb.TX, gameID int64, tagIDs []int64) error
	RemoveGameTags(ctx context.Context, tx gdb.TX, gameID int64) error
	GetGameIDsByTagName(ctx context.Context, tagName string) (out []int64, err error)
	GetTagsByGameID(ctx context.Context, gameID int64) (outs []*model.Tag, err error)
	// GetTagGames(ctx context.Context, tagID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error)
}

var localMetadata IMetadata

func Metadata() IMetadata {
	if localMetadata == nil {
		panic("implement not found for interface IMetadata, forgot register?")
	}
	return localMetadata
}

func RegisterMetadata(i IMetadata) {
	localMetadata = i
}
