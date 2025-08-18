package dao

import (
	"GameEngine/internal/dao/internal"
)

// gameCategoryDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameCategoryDao struct {
	*internal.GameCategoryDao
}

var (
	// GameCategory is globally public accessible object for table tools_gen_table operations.
	GameCategory = gameCategoryDao{
		internal.NewGameCategoryDao(),
	}
)

// Fill with you ideas below.
