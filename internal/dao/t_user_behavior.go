package dao

import (
	"GameEngine/internal/dao/internal"
)

// userBehaviorDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type userBehaviorDao struct {
	*internal.UserBehaviorDao
}

var (
	// UserBehavior is globally public accessible object for table t_user_behavior operations.
	UserBehavior = userBehaviorDao{
		internal.NewUserBehaviorDao(),
	}
)

// Fill with you ideas below.
