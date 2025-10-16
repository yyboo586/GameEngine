package metadata

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"database/sql"
	"errors"
	"strings"
)

var (
	ErrCategoryExists      = errors.New("分类已存在")
	ErrCategoryNotExists   = errors.New("分类不存在")
	ErrCategoryHasGameByID = errors.New("该分类有关联的游戏，无法删除")
)

func (m *metadata) CreateCategory(ctx context.Context, name string) (id int64, err error) {
	dataInsert := map[string]interface{}{
		dao.Category.Columns().Name: name,
	}

	id, err = dao.Category.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = ErrCategoryExists
		}
		return
	}

	return
}

// 业务逻辑：
// 1、检查分类是否存在
// 2、检查分类是否有关联的游戏
// 2.1 如果有，则返回错误
// 2.2 如果没有，则删除分类
func (m *metadata) DeleteCategory(ctx context.Context, id int64) (err error) {
	if err = m.AssertCategoryExists(ctx, id); err != nil {
		return
	}

	exists, err := m.IsCategoryHasGame(ctx, id)
	if err != nil {
		return
	}
	if exists {
		return ErrCategoryHasGameByID
	}

	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Delete()

	return
}

func (m *metadata) UpdateCategory(ctx context.Context, id int64, name string) (err error) {
	if err = m.AssertCategoryExists(ctx, id); err != nil {
		return
	}

	dataUpdate := map[string]interface{}{
		dao.Category.Columns().Name: name,
	}

	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Data(dataUpdate).Update()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = ErrCategoryExists
		}
		return
	}

	return
}

func (m *metadata) GetCategoryByID(ctx context.Context, id int64) (out *model.Category, err error) {
	var category entity.Category
	err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Scan(&category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotExists
		}
		return
	}

	out = model.ConvertCategoryEntityToModel(&category)
	return
}

// TODO:游戏分类暂时不会太多，所以直接返回所有分类，不做分页
func (m *metadata) SearchCategory(ctx context.Context, name string) (outs []*model.Category, err error) {
	var categories []*entity.Category
	if name == "" {
		err = dao.Category.Ctx(ctx).Scan(&categories)
	} else {
		err = dao.Category.Ctx(ctx).WhereLike(dao.Category.Columns().Name, name+"%").Scan(&categories)
	}
	if err != nil {
		return
	}

	outs = make([]*model.Category, 0, len(categories))
	for _, category := range categories {
		outs = append(outs, model.ConvertCategoryEntityToModel(category))
	}

	return
}

func (m *metadata) AssertCategoryExists(ctx context.Context, id int64) (err error) {
	exists, err := dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Exist()
	if err != nil {
		return
	}
	if !exists {
		return ErrCategoryNotExists
	}
	return
}
