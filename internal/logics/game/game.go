package game

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"fmt"
	"sync"
)

var (
	gameOnce     sync.Once
	gameInstance *Game
)

type Game struct {
}

func NewGame() service.IGame {
	gameOnce.Do(func() {
		gameInstance = &Game{}
	})
	return gameInstance
}

func (gg *Game) Download(ctx context.Context, gameID, userID int64) (err error) {

	return
}

// 游戏搜索相关方法
func (gg *Game) SearchGameByGameName(ctx context.Context, name string, page, size int) (out []*model.Game, pageRes *model.PageRes, err error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 20 // 搜索接口最多返回20条
	}
	pageRes = &model.PageRes{
		CurrentPage: page,
	}

	// 使用LIKE进行模糊搜索
	query := dao.Game.Ctx(ctx).WhereLike(dao.Game.Columns().Name, "%"+name+"%")

	var entities []*entity.Game
	err = query.Page(page, size).OrderDesc(dao.Game.Columns().CreateTime).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		out = append(out, model.ConvertGameEntityToModel(entity))
	}
	return
}

func (gg *Game) AssertExists(ctx context.Context, gameID int64) (err error) {
	exists, err := dao.Game.Ctx(ctx).Where(dao.Game.Columns().ID, gameID).Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("游戏不存在")
	}
	return
}
