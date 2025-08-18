package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

// CreateTag 创建标签
func (c *metadata) CreateTag(ctx context.Context, req *v1.CreateTagReq) (res *v1.CreateTagRes, err error) {
	id, err := service.Metadata().CreateTag(ctx, req.Name)
	if err != nil {
		return
	}

	return &v1.CreateTagRes{ID: id}, nil
}

// UpdateTag 更新标签
func (c *metadata) UpdateTag(ctx context.Context, req *v1.UpdateTagReq) (res *v1.UpdateTagRes, err error) {
	err = service.Metadata().UpdateTag(ctx, req.ID, req.Name)
	if err != nil {
		return
	}

	return
}

// DeleteTag 删除标签
func (c *metadata) DeleteTag(ctx context.Context, req *v1.DeleteTagReq) (res *v1.DeleteTagRes, err error) {
	err = service.Metadata().DeleteTag(ctx, req.ID)
	if err != nil {
		return
	}

	return
}

// GetTag 获取标签
func (c *metadata) GetTag(ctx context.Context, req *v1.GetTagReq) (res *v1.GetTagRes, err error) {
	tag, err := service.Metadata().GetTagByID(ctx, req.ID)
	if err != nil {
		return
	}
	if tag == nil {
		return nil, nil
	}

	return &v1.GetTagRes{TagInfo: c.convertTagModelToResponse(tag)}, nil
}

// ListTag 获取标签列表
func (c *metadata) ListTag(ctx context.Context, req *v1.GetTagListReq) (res *v1.GetTagListRes, err error) {
	outs, err := service.Metadata().ListTag(ctx)
	if err != nil {
		return
	}

	res = &v1.GetTagListRes{
		List: make([]*v1.TagInfo, 0, len(outs)),
	}
	for _, out := range outs {
		res.List = append(res.List, c.convertTagModelToResponse(out))
	}
	return
}

func (c *metadata) convertTagModelToResponse(in *model.Tag) (out *v1.TagInfo) {
	out = &v1.TagInfo{
		ID:         in.ID,
		Name:       in.Name,
		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}
	return
}
