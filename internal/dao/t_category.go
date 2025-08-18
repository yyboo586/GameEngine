package dao

import (
	"GameEngine/internal/dao/internal"
)

// categoryDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type categoryDao struct {
	*internal.CategoryDao
}

var (
	// Category is globally public accessible object for table t_game_category operations.
	Category = categoryDao{
		internal.NewCategoryDao(),
	}
)

// Fill with you ideas below.
