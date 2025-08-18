package dao

import (
	"GameEngine/internal/dao/internal"
)

// gameDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameDao struct {
	*internal.GameDao
}

var (
	// Tag is globally public accessible object for table tools_gen_table operations.
	Game = gameDao{
		internal.NewGameDao(),
	}
)

// Fill with you ideas below.
