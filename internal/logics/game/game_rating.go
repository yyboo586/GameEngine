package game

import (
	"GameEngine/internal/dao"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (gg *Game) AddRating(ctx context.Context, gameID, userID int64, score int) (err error) {
	if score < 1 || score > 5 {
		err = fmt.Errorf("评分必须在1-5之间")
		return
	}

	err = gg.assertRatingNotExists(ctx, gameID, userID)
	if err != nil {
		return
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		dataInsert := map[string]interface{}{
			dao.GameRating.Columns().GameID: gameID,
			dao.GameRating.Columns().UserID: userID,
			dao.GameRating.Columns().Score:  score,
		}
		_, err = dao.GameRating.Ctx(ctx).TX(tx).Data(dataInsert).Insert()
		if err != nil {
			return err
		}
		_, err = dao.Game.Ctx(ctx).TX(tx).Where(dao.Game.Columns().ID, gameID).Increment(dao.Game.Columns().RatingScore, score)
		if err != nil {
			return err
		}
		_, err = dao.Game.Ctx(ctx).TX(tx).Where(dao.Game.Columns().ID, gameID).Increment(dao.Game.Columns().RatingCount, 1)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (gg *Game) assertRatingNotExists(ctx context.Context, gameID, userID int64) (err error) {
	exists, err := dao.GameRating.Ctx(ctx).
		Where(dao.GameRating.Columns().GameID, gameID).
		Where(dao.GameRating.Columns().UserID, userID).Exist()
	if err != nil {
		return err
	}
	if exists {
		err = fmt.Errorf("已经评分过")
		return
	}
	return nil
}
