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

func (m *metadata) CreateTag(ctx context.Context, name string) (id int64, err error) {
	dataInsert := map[string]interface{}{
		dao.Tag.Columns().Name: name,
	}

	id, err = dao.Tag.Ctx(ctx).Data(dataInsert).InsertAndGetId()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("标签名称已存在, name: %s", name)
		}
		return
	}

	return
}

func (m *metadata) DeleteTag(ctx context.Context, id int64) (err error) {
	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Delete()

	return
}

func (m *metadata) UpdateTag(ctx context.Context, id int64, name string) (err error) {
	dataUpdate := map[string]interface{}{
		dao.Tag.Columns().Name: name,
	}

	_, err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().ID, id).Data(dataUpdate).Update()
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = fmt.Errorf("标签名称已存在, name: %s", name)
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
			return nil, nil
		}
		return
	}

	out = model.ConvertTagEntityToModel(&tag)
	return
}

func (m *metadata) ListTag(ctx context.Context) (outs []*model.Tag, err error) {
	var tags []*entity.Tag
	err = dao.Tag.Ctx(ctx).Scan(&tags)
	if err != nil {
		return
	}

	outs = make([]*model.Tag, 0, len(tags))
	for _, tag := range tags {
		outs = append(outs, model.ConvertTagEntityToModel(tag))
	}
	return
}
