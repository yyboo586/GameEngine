package dao

import (
	"GameEngine/internal/dao/internal"
)

// tagDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tagDao struct {
	*internal.TagDao
}

var (
	// Tag is globally public accessible object for table tools_gen_table operations.
	Tag = tagDao{
		internal.NewTagDao(),
	}
)

// Fill with you ideas below.
