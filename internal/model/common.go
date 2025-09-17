package model

type Author struct {
	Authorization string `p:"Authorization" in:"header" v:"required#Authorization不能为空" dc:"Bearer {{token}}"`
}

type PageReq struct {
	Page int `json:"page" d:"1" dc:"页码"`
	Size int `json:"size" d:"10" dc:"每页数量"`
}

type PageRes struct {
	Total       int `json:"total" dc:"总数"`
	CurrentPage int `json:"current_page" dc:"当前页码"`
}

type UserCtxKey string

const (
	UserInfoKey UserCtxKey = "user_info"
)

type User struct {
	ID   int64  `json:"id" dc:"用户ID"`
	Name string `json:"name" dc:"用户名"`
}

// ReservationUser 预约用户信息
type ReservationUser struct {
	ID          int64  `json:"id" dc:"预约记录ID"`
	UserID      int64  `json:"user_id" dc:"用户ID"`
	UserName    string `json:"user_name" dc:"用户名"`
	ReserveTime string `json:"reserve_time" dc:"预约时间"`
}
