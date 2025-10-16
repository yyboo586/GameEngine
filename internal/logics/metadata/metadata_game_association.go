package metadata

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"database/sql"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AddGameCategory 添加游戏分类关联
func (m *metadata) AddGameCategory(ctx context.Context, tx gdb.TX, gameID, categoryID int64) (err error) {
	err = m.AssertCategoryExists(ctx, categoryID)
	if err != nil {
		return
	}

	_, err = dao.GameCategory.Ctx(ctx).TX(tx).Insert(map[string]interface{}{
		dao.GameCategory.Columns().GameID:     gameID,
		dao.GameCategory.Columns().CategoryID: categoryID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			g.Log().Warningf(ctx, "游戏分类关联已存在, gameID: %d, categoryID: %d", gameID, categoryID)
			return nil
		}
		return
	}

	return
}

// RemoveGameCategory 移除游戏分类关联
func (m *metadata) RemoveGameCategory(ctx context.Context, tx gdb.TX, gameID int64) (err error) {
	_, err = dao.GameCategory.Ctx(ctx).TX(tx).
		Where(dao.GameCategory.Columns().GameID, gameID).
		Delete()

	return
}

func (m *metadata) GetCategoryByGameID(ctx context.Context, gameID int64) (out *model.Category, err error) {
	var entityCategory entity.Category
	err = dao.Category.Ctx(ctx).
		As("c").
		Fields("c.*").
		LeftJoin(dao.GameCategory.Table()+" as gc", "c.id = gc.category_id").
		Where("gc.game_id", gameID).
		Scan(&entityCategory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}

	out = model.ConvertCategoryEntityToModel(&entityCategory)
	return
}

func (m *metadata) GetGameIDsByCategoryName(ctx context.Context, categoryName string) (out []int64, err error) {
	var entityGameCategory []*entity.GameCategory
	err = dao.GameCategory.Ctx(ctx).
		As("gc").
		Fields("gc.game_id").
		LeftJoin(dao.Category.Table()+" as c", "c.id = gc.category_id").
		Where("c.name", categoryName).
		Scan(&entityGameCategory)

	for _, gameCategory := range entityGameCategory {
		out = append(out, gameCategory.GameID)
	}

	return
}

func (m *metadata) IsCategoryHasGame(ctx context.Context, categoryID int64) (bool, error) {
	exists, err := dao.GameCategory.Ctx(ctx).Where(dao.GameCategory.Columns().CategoryID, categoryID).Exist()
	if err != nil {
		return false, err
	}
	return exists, nil
}

/*
// GetCategoryGames 获取分类下的游戏列表
func (m *metadata) GetCategoryGames(ctx context.Context, categoryID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	query := dao.Game.Ctx(ctx).
		Fields("t_game.*").
		LeftJoin("t_game_category", "t_game_category.game_id = t_game.id").
		Where("t_game_category.category_id", categoryID)

	// 获取总数
	total, err := query.Count()
	if err != nil {
		return
	}

	var games []*entity.Game
	err = query.Page(pageReq.Page, pageReq.Size).Scan(&games)
	if err != nil {
		return
	}

	// 构建分页信息
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}
*/

// AddGameTags 添加游戏标签关联
func (m *metadata) AddGameTags(ctx context.Context, tx gdb.TX, gameID int64, tagIDs []int64) (err error) {
	for _, tagID := range tagIDs {
		err = m.AssertTagExists(ctx, tagID)
		if err != nil {
			return
		}
	}

	if len(tagIDs) == 0 {
		return
	}

	dataInserts := make([]map[string]interface{}, 0, len(tagIDs))
	for _, tagID := range tagIDs {
		dataInserts = append(dataInserts, map[string]interface{}{
			dao.GameTag.Columns().GameID: gameID,
			dao.GameTag.Columns().TagID:  tagID,
		})
	}
	_, err = dao.GameTag.Ctx(ctx).TX(tx).Insert(dataInserts)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			g.Log().Warningf(ctx, "游戏标签关联已存在, gameID: %d, tagIDs: %v", gameID, tagIDs)
			return nil
		}
		return
	}

	return
}

// RemoveGameTags 移除游戏标签关联
func (m *metadata) RemoveGameTags(ctx context.Context, tx gdb.TX, gameID int64) (err error) {
	_, err = dao.GameTag.Ctx(ctx).TX(tx).
		Where(dao.GameTag.Columns().GameID, gameID).
		Delete()

	return
}

// GetTagsByGameID 获取游戏标签列表
func (m *metadata) GetTagsByGameID(ctx context.Context, gameID int64) (outs []*model.Tag, err error) {
	var entityTags []*entity.Tag
	err = dao.Tag.Ctx(ctx).
		As("t").
		Fields("t.*").
		LeftJoin(dao.GameTag.Table()+" as gt", "t.id = gt.tag_id").
		Where("gt.game_id", gameID).
		Scan(&entityTags)
	if err != nil {
		return
	}

	outs = make([]*model.Tag, 0, len(entityTags))
	for _, tag := range entityTags {
		outs = append(outs, model.ConvertTagEntityToModel(tag))
	}
	return
}

func (m *metadata) GetGameIDsByTagName(ctx context.Context, tagName string) (out []int64, err error) {
	var entityGameTag []*entity.GameTag
	err = dao.GameTag.Ctx(ctx).
		As("gt").
		Fields("gt.game_id").
		LeftJoin(dao.Tag.Table()+" as t", "t.id = gt.tag_id").
		Where("t.name", tagName).
		Scan(&entityGameTag)
	if err != nil {
		return
	}

	for _, gameTag := range entityGameTag {
		out = append(out, gameTag.GameID)
	}
	return
}

/*
// GetTagGames 获取标签下的游戏列表
func (m *metadata) GetTagGames(ctx context.Context, tagID int64, pageReq *model.PageReq) (outs []*model.Game, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	query := dao.Game.Ctx(ctx).
		Fields("t_game.*").
		LeftJoin("t_game_tag", "t_game_tag.game_id = t_game.id").
		Where("t_game_tag.tag_id", tagID)

	// 获取总数
	total, err := query.Count()
	if err != nil {
		return
	}

	var games []*entity.Game
	err = query.Page(pageReq.Page, pageReq.Size).Scan(&games)
	if err != nil {
		return
	}

	// 构建分页信息
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}
*/

func (m *metadata) IsTagHasGame(ctx context.Context, tagID int64) (bool, error) {
	exists, err := dao.GameTag.Ctx(ctx).Where(dao.GameTag.Columns().TagID, tagID).Exist()
	if err != nil {
		return false, err
	}
	return exists, nil
}
