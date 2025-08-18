package dao

import (
	"GameEngine/internal/dao/internal"
)

// gameDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameTagDao struct {
	*internal.GameTagDao
}

var (
	// GameTag is globally public accessible object for table tools_gen_table operations.
	GameTag = gameTagDao{
		internal.NewGameTagDao(),
	}
)

// Fill with you ideas below.
