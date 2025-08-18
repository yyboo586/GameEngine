package dao

import "GameEngine/internal/dao/internal"

// gameFavoriteDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameFavoriteDao struct {
	*internal.GameFavoriteDao
}

var (
	// GameFavorite is globally public accessible object for table t_game_favorite operations.
	GameFavorite = gameFavoriteDao{
		internal.NewGameFavoriteDao(),
	}
)
