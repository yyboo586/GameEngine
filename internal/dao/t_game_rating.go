package dao

import "GameEngine/internal/dao/internal"

// gameRatingDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameRatingDao struct {
	*internal.GameRatingDao
}

var (
	// GameRating is globally public accessible object for table t_game_rating operations.
	GameRating = gameRatingDao{
		internal.NewGameRatingDao(),
	}
)
