package dao

import (
	"GameEngine/internal/dao/internal"
)

// gameBehaviorDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type gameBehaviorDao struct {
	*internal.GameBehaviorDao
}

var (
	// GameBehavior is globally public accessible object for table t_game_behavior operations.
	GameBehavior = gameBehaviorDao{
		internal.NewGameBehaviorDao(),
	}
)

// Fill with you ideas below.
