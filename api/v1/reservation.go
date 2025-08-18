package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 预约相关API结构体

// ReserveGameReq 游戏预约请求
type ReserveGameReq struct {
	g.Meta `path:"/games/{game_id}/reserve" method:"post" tags:"Game Management/Reservation" summary:"Reserve Game"`
	model.Author
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// ReserveGameRes 游戏预约响应
type ReserveGameRes struct {
	g.Meta `mime:"application/json"`
}

// CancelReservationReq 取消预约请求
type CancelReservationReq struct {
	g.Meta `path:"/games/{game_id}/reserve" method:"delete" tags:"Game Management/Reservation" summary:"Cancel Reservation"`
	model.Author
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// CancelReservationRes 取消预约响应
type CancelReservationRes struct {
	g.Meta `mime:"application/json"`
}

// GetUserReservationsReq 获取用户预约列表请求
type GetUserReservationsReq struct {
	g.Meta `path:"/games/reservations" method:"get" tags:"Game Management/Reservation" summary:"Get User Reservations"`
	model.Author
	model.PageReq
}

// GetUserReservationsRes 获取用户预约列表响应
type GetUserReservationsRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// IsUserReservedReq 检查用户是否已预约请求
type IsUserReservedReq struct {
	g.Meta `path:"/games/{game_id}/is-reserved" method:"get" tags:"Game Management/Reservation" summary:"Check User Reservation Status"`
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// IsUserReservedRes 检查用户是否已预约响应
type IsUserReservedRes struct {
	g.Meta     `mime:"application/json"`
	IsReserved bool `json:"is_reserved" dc:"是否已预约"`
}
