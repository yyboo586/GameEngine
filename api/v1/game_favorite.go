package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AddGameFavoriteReq struct {
	g.Meta `path:"/games/{game_id}/favorite" method:"post" tags:"Game Management/Favorite" summary:"Add Game Favorite"`
	model.Author
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type AddGameFavoriteRes struct {
	g.Meta `mime:"application/json"`
}

type RemoveGameFavoriteReq struct {
	g.Meta `path:"/games/{game_id}/favorite" method:"delete" tags:"Game Management/Favorite" summary:"Remove Game Favorite"`
	model.Author
	GameID int64 `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
}

type RemoveGameFavoriteRes struct {
	g.Meta `mime:"application/json"`
}

type GetGameFavoriteReq struct {
	g.Meta `path:"/games/favorites" method:"get" tags:"Game Management/Favorite" summary:"Get User Game Favorite"`
	model.Author
	PageReq model.PageReq
}

type GetGameFavoriteRes struct {
	g.Meta `mime:"application/json"`
	List   []*Game `json:"list" dc:"游戏列表"`
	*model.PageRes
}
