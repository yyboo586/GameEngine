package game

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func (gg *Game) CreateGame(ctx context.Context, in *v1.CreateGameReq) (id int64, err error) {
	dataGameInsert := map[string]interface{}{
		dao.Game.Columns().Name:           in.Name,
		dao.Game.Columns().DistributeType: in.DistributeType,
		dao.Game.Columns().Developer:      in.Developer,
		dao.Game.Columns().Publisher:      in.Publisher,
		dao.Game.Columns().Description:    in.Description,
		dao.Game.Columns().Details:        in.Details,
		dao.Game.Columns().Status:         model.GameStatusUnpublished,
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err = dao.Game.Ctx(ctx).TX(tx).Data(dataGameInsert).InsertAndGetId()
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				return fmt.Errorf("游戏名称已存在")
			}
			return err
		}

		err = service.Metadata().AddGameCategory(ctx, tx, id, in.CategoryID)
		if err != nil {
			return err
		}

		if len(in.TagIDs) > 0 {
			err = service.Metadata().AddGameTags(ctx, tx, id, in.TagIDs)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return
	}

	return
}

func (gg *Game) DeleteGame(ctx context.Context, id int64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Game.Ctx(ctx).TX(tx).Where(dao.Game.Columns().ID, id).Delete()
		if err != nil {
			return err
		}

		err = service.Metadata().RemoveGameCategory(ctx, tx, id)
		if err != nil {
			return err
		}

		err = service.Metadata().RemoveGameTags(ctx, tx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// TODO: 这个更新逻辑该怎么设计呢
func (gg *Game) UpdateGame(ctx context.Context, in *v1.UpdateGameReq) (err error) {
	return
}

func (gg *Game) GetGameByID(ctx context.Context, id int64) (out *model.Game, err error) {
	var entity entity.Game
	err = dao.Game.Ctx(ctx).Where(dao.Game.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}

	out = model.ConvertGameEntityToModel(&entity)
	return
}

func (gg *Game) GetGamesByIDs(ctx context.Context, ids []int64) (out []*model.Game, err error) {
	var entities []*entity.Game
	err = dao.Game.Ctx(ctx).WhereIn(dao.Game.Columns().ID, ids).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		out = append(out, model.ConvertGameEntityToModel(entity))
	}
	return
}

func (gg *Game) ListGame(ctx context.Context, in *v1.ListGameReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if in.PageReq.Page == 0 {
		in.PageReq.Page = 1
	}
	if in.PageReq.Size == 0 {
		in.PageReq.Size = 10
	}

	query := dao.Game.Ctx(ctx)

	total, err := query.Count()
	if err != nil {
		return
	}

	var entities []*entity.Game
	err = query.Page(in.PageReq.Page, in.PageReq.Size).OrderDesc(dao.Game.Columns().CreateTime).Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		outs = append(outs, model.ConvertGameEntityToModel(entity))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: in.PageReq.Page,
	}
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

// 游戏状态管理

func (gg *Game) UpdateGamePublishStatus(ctx context.Context, id int64, status int) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Game.Columns().Status: status,
	}
	if status == int(model.GameStatusPublished) {
		dataUpdate[dao.Game.Columns().PublishTime] = time.Now()
	}
	_, err = dao.Game.Ctx(ctx).Where(dao.Game.Columns().ID, id).Data(dataUpdate).Update()
	return
}
