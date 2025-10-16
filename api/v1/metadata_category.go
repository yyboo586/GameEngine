package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*
分类管理
1、均为 管理控制台/开发者后台 调用的接口，所以需要令牌。
*/

// CreateCategoryReq 创建分类请求
type CreateCategoryReq struct {
	g.Meta `path:"/categories" method:"post" tags:"Metadata/Category" summary:"Create Category"`
	model.AuthorRequired
	Name string `json:"name" v:"required|length:1,6#分类名称不能为空|分类名称长度不能超过6个字符" dc:"分类名称"`
}

// CreateCategoryRes 创建分类响应
type CreateCategoryRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"分类ID"`
}

// UpdateCategoryReq 更新分类请求
type UpdateCategoryReq struct {
	g.Meta `path:"/categories/{id}" method:"put" tags:"Metadata/Category" summary:"Update Category"`
	model.AuthorRequired
	ID   int64  `p:"id" v:"required#分类ID不能为空" dc:"分类ID"`
	Name string `json:"name" v:"required|length:1,6#分类名称不能为空|分类名称长度不能超过6个字符" dc:"分类名称"`
}

// UpdateCategoryRes 更新分类响应
type UpdateCategoryRes struct {
	g.Meta `mime:"application/json"`
}

// DeleteCategoryReq 删除分类请求
type DeleteCategoryReq struct {
	g.Meta `path:"/categories/{id}" method:"delete" tags:"Metadata/Category" summary:"Delete Category"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

// DeleteCategoryRes 删除分类响应
type DeleteCategoryRes struct {
	g.Meta `mime:"application/json"`
}

// GetCategoryReq 获取分类请求
type GetCategoryReq struct {
	g.Meta `path:"/categories/{id}" method:"get" tags:"Metadata/Category" summary:"Get Category"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

// GetCategoryRes 获取分类响应
type GetCategoryRes struct {
	g.Meta `mime:"application/json"`
	*CategoryInfo
}

// SearchCategoryReq 搜索分类请求
type SearchCategoryReq struct {
	g.Meta `path:"/categories" method:"get" tags:"Metadata/Category" summary:"Search Category"`
	model.AuthorRequired
	Name string `json:"name" dc:"分类名称"`
}

// SearchCategoryRes 搜索分类响应
type SearchCategoryRes struct {
	g.Meta `mime:"application/json"`
	List   []*CategoryInfo `json:"list" dc:"分类列表"`
	*model.PageRes
}

type CategoryInfo struct {
	ID         int64       `json:"id" dc:"分类ID"`
	Name       string      `json:"name" dc:"分类名称"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}
