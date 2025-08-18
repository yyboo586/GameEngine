package model

import (
	"GameEngine/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// Category 游戏分类模型
type Category struct {
	ID         int64       `json:"id" dc:"分类ID"`
	Name       string      `json:"name" dc:"分类名称"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}

type Tag struct {
	ID         int64       `json:"id" dc:"标签ID"`
	Name       string      `json:"name" dc:"标签名称"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}

func ConvertCategoryEntityToModel(in *entity.Category) (out *Category) {
	out = &Category{
		ID:         in.ID,
		Name:       in.Name,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}
	return
}

func ConvertTagEntityToModel(in *entity.Tag) (out *Tag) {
	out = &Tag{
		ID:         in.ID,
		Name:       in.Name,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}
	return
}
