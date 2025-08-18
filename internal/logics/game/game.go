package game

import (
	"GameEngine/internal/dao"
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

func (gg *Game) AssertGameExists(ctx context.Context, gameID int64) (err error) {
	exists, err := dao.Game.Ctx(ctx).Where(dao.Game.Columns().ID, gameID).Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("游戏不存在")
	}
	return
}
