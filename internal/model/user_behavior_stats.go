package model

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/os/gtime"
)

type BehaviorType int

const (
	_ BehaviorType = iota
	BehaviorSearch
	BehaviorPlay
	BehaviorDownload
)

func GetBehaviorTypeString(behaviorType BehaviorType) string {
	switch behaviorType {
	case BehaviorSearch:
		return "Search"
	case BehaviorPlay:
		return "Play"
	case BehaviorDownload:
		return "Download"
	default:
		return "Unknown"
	}
}

type UserBehavior struct {
	ID            int64        `json:"id" dc:"ID"`
	UserID        int64        `json:"user_id" dc:"用户ID"`
	GameID        int64        `json:"game_id" dc:"游戏ID"`
	BehaviorType  BehaviorType `json:"behavior_type" dc:"行为类型"`
	SearchKeyword string       `json:"search_keyword" dc:"搜索关键词"`
	BehaviorTime  *gtime.Time  `json:"behavior_time" dc:"行为时间"`
	IPAddress     string       `json:"ip_address" dc:"IP地址"`
}

func GetUserInfo(ctx context.Context) (userInfo *User, err error) {
	var ok bool
	value := ctx.Value(UserInfoKey)
	if value == nil {
		err = errors.New("token is required")
		return
	}
	v, ok := value.(User)
	if !ok {
		err = errors.New("user info is not a User")
		return
	}
	userInfo = &v
	return
}
