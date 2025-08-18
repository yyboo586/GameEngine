package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 游戏评分相关
type AddGameRatingReq struct {
	g.Meta `path:"/games/{game_id}/rating" method:"post" tags:"Game Management/Rating" summary:"Add Game Rating"`
	model.Author
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	Rating int   `json:"rating" v:"required|min:1|max:5#评分不能为空|评分必须在1-5之间|评分必须在1-5之间" dc:"评分(1-5)"`
}

type AddGameRatingRes struct {
	g.Meta `mime:"application/json"`
}
