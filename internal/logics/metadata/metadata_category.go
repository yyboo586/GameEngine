package metadata

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (m *metadata) CreateCategory(ctx context.Context, name string) (id int64, err error) {
	dataInsert := map[string]interface{}{
		dao.Category.Columns().Name: name,
	}

	id, err = dao.Category.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("分类名称已存在, name: %s", name)
		}
		return
	}

	return
}

func (m *metadata) DeleteCategory(ctx context.Context, id int64) (err error) {
	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Delete()

	return
}

func (m *metadata) UpdateCategory(ctx context.Context, id int64, name string) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Category.Columns().Name: name,
	}

	_, err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().ID, id).Data(dataUpdate).Update()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("分类名称已存在, name: %s", name)
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
			return nil, nil
		}
		return
	}

	out = model.ConvertCategoryEntityToModel(&category)
	return
}

func (m *metadata) ListCategory(ctx context.Context) (outs []*model.Category, err error) {
	var categories []*entity.Category
	err = dao.Category.Ctx(ctx).Scan(&categories)
	if err != nil {
		return
	}

	outs = make([]*model.Category, 0, len(categories))
	for _, category := range categories {
		outs = append(outs, model.ConvertCategoryEntityToModel(category))
	}

	return
}
