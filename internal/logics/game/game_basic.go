package game

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	ErrGameNotFound = errors.New("游戏不存在")
)

func (gg *Game) CreateGame(ctx context.Context, in *v1.CreateGameReq) (id int64, err error) {
	dataGameInsert := map[string]interface{}{
		dao.Game.Columns().Name:           in.Name,
		dao.Game.Columns().DistributeType: in.DistributeType,
		dao.Game.Columns().Developer:      in.Developer,
		dao.Game.Columns().Publisher:      in.Publisher,
		dao.Game.Columns().Description:    in.Description,
		dao.Game.Columns().Details:        in.Details,
		dao.Game.Columns().Status:         model.GameStatusInit,
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

// 1、删除游戏
// 2、删除游戏分类关联
// 3、删除游戏标签关联
// 4、删除游戏评分
// 5、删除游戏收藏
// 6、删除游戏预约
// 4、删除游戏媒体关联（TODO: 基于异步任务队列）
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

func (gg *Game) UpdateGame(ctx context.Context, in *v1.UpdateGameReq) (err error) {
	// 检查游戏是否存在
	err = gg.AssertExists(ctx, in.ID)
	if err != nil {
		return err
	}

	// 构建更新数据
	updateData := make(map[string]interface{})
	if in.Name != "" {
		updateData[dao.Game.Columns().Name] = in.Name
	}
	if in.DistributeType > 0 {
		updateData[dao.Game.Columns().DistributeType] = in.DistributeType
	}
	if in.Developer != "" {
		updateData[dao.Game.Columns().Developer] = in.Developer
	}
	if in.Publisher != "" {
		updateData[dao.Game.Columns().Publisher] = in.Publisher
	}
	if in.Description != "" {
		updateData[dao.Game.Columns().Description] = in.Description
	}
	if in.Details != "" {
		updateData[dao.Game.Columns().Details] = in.Details
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if len(updateData) > 0 {
			_, err = dao.Game.Ctx(ctx).TX(tx).
				Where(dao.Game.Columns().ID, in.ID).
				Data(updateData).
				Update()
			if err != nil {
				return err
			}
		}

		// 处理分类更新
		if in.CategoryID > 0 {
			// 先移除旧的分类关联
			err = service.Metadata().RemoveGameCategory(ctx, tx, in.ID)
			if err != nil {
				return err
			}
			// 添加新的分类关联
			err = service.Metadata().AddGameCategory(ctx, tx, in.ID, in.CategoryID)
			if err != nil {
				return err
			}
		}

		// 处理标签更新
		if len(in.TagIDs) > 0 {
			// 先移除旧的标签关联
			err = service.Metadata().RemoveGameTags(ctx, tx, in.ID)
			if err != nil {
				return err
			}
			// 添加新的标签关联
			if len(in.TagIDs) > 0 {
				err = service.Metadata().AddGameTags(ctx, tx, in.ID, in.TagIDs)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return
}

func (gg *Game) GetGameByID(ctx context.Context, id int64) (out *model.Game, err error) {
	var entity entity.Game
	err = dao.Game.Ctx(ctx).Where(dao.Game.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrGameNotFound
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

	if in.Name != "" {
		query = query.WhereLike(dao.Game.Columns().Name, in.Name+"%")
	}

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
