package game

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// 游戏收藏相关方法
func (gg *Game) AddFavorite(ctx context.Context, gameID, userID int64) (err error) {
	err = gg.AssertGameExists(ctx, gameID)
	if err != nil {
		return
	}

	gameInfo, err := gg.getGameFavoriteByGameIDAndUserID(ctx, gameID, userID)
	if err != nil {
		return
	}
	if gameInfo != nil {
		return // 已经收藏过了
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GameFavorite.Ctx(ctx).TX(tx).Data(map[string]interface{}{
			dao.GameFavorite.Columns().GameID: gameID,
			dao.GameFavorite.Columns().UserID: userID,
		}).Insert()
		if err != nil {
			return err
		}

		_, err = dao.Game.Ctx(ctx).TX(tx).Where(dao.Game.Columns().ID, gameID).Increment(dao.Game.Columns().FavoriteCount, 1)
		if err != nil {
			return err
		}
		return nil
	})

	return
}

func (gg *Game) RemoveFavorite(ctx context.Context, gameID, userID int64) (err error) {
	err = gg.AssertGameExists(ctx, gameID)
	if err != nil {
		return
	}

	gameInfo, err := gg.getGameFavoriteByGameIDAndUserID(ctx, gameID, userID)
	if err != nil {
		return
	}
	if gameInfo == nil {
		return // 没有收藏过
	}

	// 删除收藏记录
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GameFavorite.Ctx(ctx).TX(tx).
			Where(dao.GameFavorite.Columns().GameID, gameID).
			Where(dao.GameFavorite.Columns().UserID, userID).Delete()
		if err != nil {
			return err
		}

		_, err = dao.Game.Ctx(ctx).TX(tx).Where(dao.Game.Columns().ID, gameID).Decrement(dao.Game.Columns().FavoriteCount, 1)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (gg *Game) GetUserFavorites(ctx context.Context, userID int64, pageReq *model.PageReq) (out []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	var entityGames []*entity.Game
	err = dao.Game.Ctx(ctx).
		As("g").
		Fields("g.*").
		LeftJoin(dao.GameFavorite.Table()+" as gf", "g.id = gf.game_id").
		Where("gf.user_id", userID).
		Scan(&entityGames)
	if err != nil {
		return
	}

	for _, entity := range entityGames {
		out = append(out, model.ConvertGameEntityToModel(entity))
	}
	return
}

func (gg *Game) getGameFavoriteByGameIDAndUserID(ctx context.Context, gameID, userID int64) (out *model.Game, err error) {
	var entityGames entity.Game
	err = dao.Game.Ctx(ctx).
		As("g").Fields("g.*").
		LeftJoin(dao.GameFavorite.Table()+" as gf", "g.id = gf.game_id").
		Where("gf.game_id", gameID).
		Where("gf.user_id", userID).
		Scan(&entityGames)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}

	out = model.ConvertGameEntityToModel(&entityGames)
	return
}
