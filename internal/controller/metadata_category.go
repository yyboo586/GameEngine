package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

// CreateCategory 创建分类
func (c *metadata) CreateCategory(ctx context.Context, req *v1.CreateCategoryReq) (res *v1.CreateCategoryRes, err error) {
	id, err := service.Metadata().CreateCategory(ctx, req.Name)
	if err != nil {
		return
	}

	return &v1.CreateCategoryRes{ID: id}, nil
}

// UpdateCategory 更新分类
func (c *metadata) UpdateCategory(ctx context.Context, req *v1.UpdateCategoryReq) (res *v1.UpdateCategoryRes, err error) {
	err = service.Metadata().UpdateCategory(ctx, req.ID, req.Name)
	if err != nil {
		return
	}

	return
}

// DeleteCategory 删除分类
func (c *metadata) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryReq) (res *v1.DeleteCategoryRes, err error) {
	err = service.Metadata().DeleteCategory(ctx, req.ID)
	if err != nil {
		return
	}

	return
}

// GetCategory 获取分类
func (c *metadata) GetCategory(ctx context.Context, req *v1.GetCategoryReq) (res *v1.GetCategoryRes, err error) {
	category, err := service.Metadata().GetCategoryByID(ctx, req.ID)
	if err != nil {
		return
	}
	if category == nil {
		return nil, nil
	}

	return &v1.GetCategoryRes{CategoryInfo: c.convertCategoryModelToResponse(category)}, nil
}

// SearchCategory 搜索分类
func (c *metadata) SearchCategory(ctx context.Context, req *v1.SearchCategoryReq) (res *v1.SearchCategoryRes, err error) {
	outs, err := service.Metadata().SearchCategory(ctx, req.Name)
	if err != nil {
		return
	}

	res = &v1.SearchCategoryRes{
		List: make([]*v1.CategoryInfo, 0, len(outs)),
	}
	for _, out := range outs {
		res.List = append(res.List, c.convertCategoryModelToResponse(out))
	}
	return
}

func (c *metadata) convertCategoryModelToResponse(in *model.Category) (out *v1.CategoryInfo) {
	out = &v1.CategoryInfo{
		ID:         in.ID,
		Name:       in.Name,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}
	return
}
