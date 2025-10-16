package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 提交审核
type SubmitGameForReviewReq struct {
	g.Meta `path:"/games/{id}/submit-review" method:"post" tags:"Game Management/Status" summary:"Submit Game For Review"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type SubmitGameForReviewRes struct {
	g.Meta `mime:"application/json"`
}

// 获取待审核游戏列表
type ListInReviewReq struct {
	g.Meta `path:"/games/in-review" method:"get" tags:"Game Management/Status" summary:"List Games In Review"`
	model.AuthorRequired
	model.PageReq
}

type ListInReviewRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	model.PageRes
}

// 审核通过
type ApproveGameReq struct {
	g.Meta `path:"/games/{id}/approve" method:"post" tags:"Game Management/Status" summary:"Approve Game"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type ApproveGameRes struct {
	g.Meta `mime:"application/json"`
}

// 审核拒绝
type RejectGameReq struct {
	g.Meta `path:"/games/{id}/reject" method:"post" tags:"Game Management/Status" summary:"Reject Game"`
	model.AuthorRequired
	ID     int64  `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	Reason string `json:"reason" dc:"拒绝原因"`
}

type RejectGameRes struct {
	g.Meta `mime:"application/json"`
}

// 立即上架游戏
type PublishGameImmediatelyReq struct {
	g.Meta `path:"/games/{id}/publish-now" method:"post" tags:"Game Management/Status" summary:"Publish Game Immediately"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type PublishGameImmediatelyRes struct {
	g.Meta `mime:"application/json"`
}

// 预约上线
type SetGamePreRegisterReq struct {
	g.Meta `path:"/games/{id}/pre-register" method:"post" tags:"Game Management/Status" summary:"Set Game Pre Register"`
	model.AuthorRequired
	ID          int64       `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	PublishTime *gtime.Time `json:"publish_time" v:"required#发布时间不能为空" dc:"发布时间"`
}

type SetGamePreRegisterRes struct {
	g.Meta `mime:"application/json"`
}

// 下架游戏
type UnpublishGameReq struct {
	g.Meta `path:"/games/{id}/unpublish" method:"post" tags:"Game Management/Status" summary:"Unpublish Game"`
	model.AuthorRequired
	ID              int64  `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	UnpublishReason string `json:"unpublish_reason" dc:"下架原因"`
}

type UnpublishGameRes struct {
	g.Meta `mime:"application/json"`
}

// 取消预约发布
type CancelPreRegisterReq struct {
	g.Meta `path:"/games/{id}/cancel-pre-register" method:"post" tags:"Game Management/Status" summary:"Cancel Pre Register"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type CancelPreRegisterRes struct {
	g.Meta `mime:"application/json"`
}

// 更新游戏信息（回到初始状态）
type UpdateGameInfoReq struct {
	g.Meta `path:"/games/{id}/update-info" method:"post" tags:"Game Management/Status" summary:"Update Game Info (Reset to Init Status)"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type UpdateGameInfoRes struct {
	g.Meta `mime:"application/json"`
}

// 更新游戏版本（回到初始状态）
type UpdateGameVersionReq struct {
	g.Meta `path:"/games/{id}/update-version" method:"post" tags:"Game Management/Status" summary:"Update Game Version (Reset to Init Status)"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type UpdateGameVersionRes struct {
	g.Meta `mime:"application/json"`
}

// 获取游戏可用事件列表
type GetAvailableEventsReq struct {
	g.Meta `path:"/games/{id}/available-events" method:"get" tags:"Game Management/Status" summary:"Get Available Events for Game"`
	model.AuthorRequired
	ID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type GetAvailableEventsRes struct {
	g.Meta `mime:"application/json"`
	Events []GameEventInfo `json:"events" dc:"可用事件列表"`
}

// 游戏事件信息
type GameEventInfo struct {
	Event          model.GameEvent  `json:"event" dc:"事件代码"`
	EventName      string           `json:"event_name" dc:"事件名称"`
	NextStatus     model.GameStatus `json:"next_status" dc:"下一个状态"`
	NextStatusName string           `json:"next_status_name" dc:"下一个状态名称"`
}
