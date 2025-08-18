package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 预约相关API结构体

// ReserveGameReq 游戏预约请求
type ReserveGameReq struct {
	g.Meta `path:"/games/{id}/reserve" method:"post" tags:"Game Management/Reservation" summary:"Reserve Game"`
	model.Author
	GameID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// ReserveGameRes 游戏预约响应
type ReserveGameRes struct {
	g.Meta `mime:"application/json"`
}

// CancelReservationReq 取消预约请求
type CancelReservationReq struct {
	g.Meta `path:"/games/{id}/reserve" method:"delete" tags:"Game Management/Reservation" summary:"Cancel Reservation"`
	model.Author
	GameID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
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

// GetThisMonthNewGamesReq 获取本月新游戏请求
type GetThisMonthNewGamesReq struct {
	g.Meta `path:"/games/this-month-new" method:"get" tags:"Game Management/Reservation" summary:"Get This Month New Games"`
	model.PageReq
}

// GetThisMonthNewGamesRes 获取本月新游戏响应
type GetThisMonthNewGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// GetUpcomingGamesReq 获取即将上新游戏请求
type GetUpcomingGamesReq struct {
	g.Meta `path:"/games/upcoming" method:"get" tags:"Game Management/Reservation" summary:"Get Upcoming Games"`
	model.PageReq
}

// GetUpcomingGamesRes 获取即将上新游戏响应
type GetUpcomingGamesRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// TODO List
// GetUserReservationsRes 获取用户预约列表响应
type GetUserReservationsRes struct {
	g.Meta  `mime:"application/json"`
	List    []*Game        `json:"list" dc:"游戏列表"`
	PageRes *model.PageRes `json:"page_res" dc:"分页信息"`
}

// GetGameReservationCountReq 获取游戏预约数量请求
type GetGameReservationCountReq struct {
	g.Meta `path:"/games/{id}/reservation-count" method:"get" tags:"Game Management/Reservation" summary:"Get Game Reservation Count"`
	GameID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// GetGameReservationCountRes 获取游戏预约数量响应
type GetGameReservationCountRes struct {
	g.Meta `mime:"application/json"`
	Count  int64 `json:"count" dc:"预约数量"`
}

// IsUserReservedReq 检查用户是否已预约请求
type IsUserReservedReq struct {
	g.Meta `path:"/games/{id}/reservation-status" method:"get" tags:"Game Management/Reservation" summary:"Check User Reservation Status"`
	UserID int64 `q:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	GameID int64 `p:"id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

// IsUserReservedRes 检查用户是否已预约响应
type IsUserReservedRes struct {
	g.Meta     `mime:"application/json"`
	IsReserved bool `json:"is_reserved" dc:"是否已预约"`
}

// GetBatchReservationStatusReq 批量获取预约状态请求
type GetBatchReservationStatusReq struct {
	g.Meta  `path:"/games/batch-reservation-status" method:"get" tags:"Game Management/Reservation" summary:"Get Batch Reservation Status"`
	UserID  int64   `q:"user_id" v:"required#用户ID不能为空" dc:"用户ID"`
	GameIDs []int64 `q:"game_ids" v:"required#游戏ID列表不能为空" dc:"游戏ID列表"`
}

// GetBatchReservationStatusRes 批量获取预约状态响应
type GetBatchReservationStatusRes struct {
	g.Meta    `mime:"application/json"`
	StatusMap map[int64]bool `json:"status_map" dc:"预约状态映射"`
}

// GetPopularReservationGamesReq 获取热门预约游戏请求
type GetPopularReservationGamesReq struct {
	g.Meta `path:"/games/popular-reservations" method:"get" tags:"Game Management/Reservation" summary:"Get Popular Reservation Games"`
	model.PageReq
}

// GetPopularReservationGamesRes 获取热门预约游戏响应
type GetPopularReservationGamesRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	Total  int64   `json:"total" dc:"总数"`
}
