package dao

import (
	"GameEngine/internal/dao/internal"
)

// gameReserveDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameReserveDao struct {
	*internal.GameReserveDao
}

var (
	// GameReserve is globally public accessible object for table t_game_reserve operations.
	GameReserve = gameReserveDao{
		internal.NewGameReserveDao(),
	}
)

// Fill with you ideas below.
