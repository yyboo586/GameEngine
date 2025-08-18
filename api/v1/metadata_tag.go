package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateTagReq 创建标签请求
type CreateTagReq struct {
	g.Meta `path:"/tags" method:"post" tags:"Metadata/Tag" summary:"Create Tag"`
	model.Author
	Name string `json:"name" v:"required|length:1,6#标签名称不能为空|标签名称长度不能超过6个字符" dc:"标签名称"`
}

// CreateTagRes 创建标签响应
type CreateTagRes struct {
	g.Meta `mime:"application/json"`
	ID     int64 `json:"id" dc:"标签ID"`
}

// DeleteTagReq 删除标签请求
type DeleteTagReq struct {
	g.Meta `path:"/tags/{id}" method:"delete" tags:"Metadata/Tag" summary:"Delete Tag"`
	model.Author
	ID int64 `p:"id" v:"required#标签ID不能为空" dc:"标签ID"`
}

// DeleteTagRes 删除标签响应
type DeleteTagRes struct {
	g.Meta `mime:"application/json"`
}

// UpdateTagReq 更新标签请求
type UpdateTagReq struct {
	g.Meta `path:"/tags/{id}" method:"put" tags:"Metadata/Tag" summary:"Update Tag"`
	model.Author
	ID   int64  `p:"id" v:"required#标签ID不能为空" dc:"标签ID"`
	Name string `json:"name" v:"required|length:1,6#标签名称不能为空|标签名称长度不能超过6个字符" dc:"标签名称"`
}

// UpdateTagRes 更新标签响应
type UpdateTagRes struct {
	g.Meta `mime:"application/json"`
}

// GetTagReq 获取标签请求
type GetTagReq struct {
	g.Meta `path:"/tags/{id}" method:"get" tags:"Metadata/Tag" summary:"Get Tag"`
	model.Author
	ID int64 `p:"id" v:"required#标签ID不能为空" dc:"标签ID"`
}

// GetTagRes 获取标签响应
type GetTagRes struct {
	g.Meta `mime:"application/json"`
	*TagInfo
}

// GetTagListReq 获取标签列表请求
type GetTagListReq struct {
	g.Meta `path:"/tags" method:"get" tags:"Metadata/Tag" summary:"Get Tag List"`
	model.Author
}

// GetTagListRes 获取标签列表响应
type GetTagListRes struct {
	g.Meta `mime:"application/json"`
	List   []*TagInfo `json:"list" dc:"标签列表"`
}

type TagInfo struct {
	ID         int64       `json:"id" dc:"标签ID"`
	Name       string      `json:"name" dc:"标签名称"`
	CreateTime *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime *gtime.Time `json:"update_time" dc:"更新时间"`
}
