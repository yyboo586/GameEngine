package dao

import (
	"GameEngine/internal/dao/internal"
)

// searchHistoryDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type searchHistoryDao struct {
	*internal.SearchHistoryDao
}

var (
	// SearchHistory is globally public accessible object for table t_search_history operations.
	SearchHistory = searchHistoryDao{
		internal.NewSearchHistoryDao(),
	}
)

// Fill with you ideas below.
