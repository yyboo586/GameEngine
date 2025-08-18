package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 提交审核
type SubmitGameForReviewReq struct {
	g.Meta `path:"/games/{id}/submit-review" method:"post" tags:"Game Management/Status" summary:"Submit Game For Review"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type SubmitGameForReviewRes struct {
	g.Meta `mime:"application/json"`
}

// 审核游戏
type ApproveGameReq struct {
	g.Meta `path:"/games/{id}/approve" method:"post" tags:"Game Management/Status" summary:"Approve Game"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type ApproveGameRes struct {
	g.Meta `mime:"application/json"`
}

// 拒绝游戏
type RejectGameReq struct {
	g.Meta `path:"/games/{id}/reject" method:"post" tags:"Game Management/Status" summary:"Reject Game"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type RejectGameRes struct {
	g.Meta `mime:"application/json"`
}

// 立即上架游戏
type PublishGameImmediatelyReq struct {
	g.Meta `path:"/games/{id}/publish-now" method:"post" tags:"Game Management/Status" summary:"Publish Game Immediately"`
	model.Author
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type PublishGameImmediatelyRes struct {
	g.Meta `mime:"application/json"`
}

// 预约上线
type SetGamePreRegisterReq struct {
	g.Meta `path:"/games/{id}/pre-register" method:"post" tags:"Game Management/Status" summary:"Set Game Pre Register"`
	model.Author
	ID          int64       `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	PublishTime *gtime.Time `json:"publish_time" v:"required#发布时间不能为空" dc:"发布时间"`
}

type SetGamePreRegisterRes struct {
	g.Meta `mime:"application/json"`
}

// 下架游戏
type UnpublishGameReq struct {
	g.Meta `path:"/games/{id}/unpublish" method:"post" tags:"Game Management/Status" summary:"Unpublish Game"`
	model.Author
	ID              int64  `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	UnpublishReason string `json:"unpublish_reason" dc:"下架原因"`
}

type UnpublishGameRes struct {
	g.Meta `mime:"application/json"`
}

type ListInReviewReq struct {
	g.Meta `path:"/games/in-review" method:"get" tags:"Game Management/Status" summary:"List Games In Review"`
	model.Author
	model.PageReq
}

type ListInReviewRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	model.PageRes
}
