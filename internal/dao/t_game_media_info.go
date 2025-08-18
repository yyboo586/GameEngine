package dao

import "GameEngine/internal/dao/internal"

// gameMediaInfoDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameMediaInfoDao struct {
	*internal.GameMediaInfoDao
}

var (
	// GameMediaInfo is globally public accessible object for table tools_gen_table operations.
	GameMediaInfo = gameMediaInfoDao{
		internal.NewGameMediaInfoDao(),
	}
)

// Fill with you ideas below.
