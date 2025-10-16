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
	ErrTagExists      = errors.New("标签已存在")
	ErrTagNotExists   = errors.New("标签不存在")
	ErrTagHasGameByID = errors.New("该标签有关联的游戏，无法删除")
)

func (m *metadata) CreateTag(ctx context.Context, name string) (id int64, err error) {
	dataInsert := map[string]interface{}{
		dao.Tag.Columns().Name: name,
	}

	id, err = dao.Tag.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = ErrTagExists
		}
		return
	}

	return
}

func (m *metadata) DeleteTag(ctx context.Context, id int64) (err error) {
	if err = m.AssertTagExists(ctx, id); err != nil {
		return
	}

	exists, err := m.IsTagHasGame(ctx, id)
	if err != nil {
		return
	}
	if exists {
		return ErrTagHasGameByID
	}

	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Delete()

	return
}

func (m *metadata) UpdateTag(ctx context.Context, id int64, name string) (err error) {
	if err = m.AssertTagExists(ctx, id); err != nil {
		return
	}

	dataUpdate := map[string]interface{}{
		dao.Tag.Columns().Name: name,
	}

	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Data(dataUpdate).Update()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = ErrTagExists
		}
		return
	}

	return
}

func (m *metadata) GetTagByID(ctx context.Context, id int64) (out *model.Tag, err error) {
	var tag entity.Tag
	err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Scan(&tag)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTagNotExists
		}
		return
	}

	out = model.ConvertTagEntityToModel(&tag)
	return
}

func (m *metadata) SearchTag(ctx context.Context, name string) (outs []*model.Tag, err error) {
	var tags []*entity.Tag
	if name == "" {
		err = dao.Tag.Ctx(ctx).Scan(&tags)
	} else {
		err = dao.Tag.Ctx(ctx).WhereLike(dao.Tag.Columns().Name, name+"%").Scan(&tags)
	}
	if err != nil {
		return
	}

	outs = make([]*model.Tag, 0, len(tags))
	for _, tag := range tags {
		outs = append(outs, model.ConvertTagEntityToModel(tag))
	}
	return
}

func (m *metadata) AssertTagExists(ctx context.Context, id int64) (err error) {
	exists, err := dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Exist()
	if err != nil {
		return
	}
	if !exists {
		return ErrTagNotExists
	}
	return
}
