package service

import (
	"GameEngine/internal/model"
	"context"
)

// IUserBehavior 用户行为服务接口
type IUserBehavior interface {
	// 搜索历史管理
	AddSearchHistory(ctx context.Context, userID int64, keyword string, resultCount int) error
	GetSearchHistory(ctx context.Context, userID int64, limit int) ([]*model.SearchHistory, error)
	ClearSearchHistory(ctx context.Context, userID int64) error
	GetPopularSearchKeywords(ctx context.Context, days int, limit int) ([]string, error)

	// 游戏行为记录
	RecordGameBehavior(ctx context.Context, userID, gameID int64, behaviorType model.BehaviorType, ipAddress string) error
	GetUserGameHistory(ctx context.Context, userID int64, limit int) ([]*model.GameBehavior, error)
	GetGameBehaviorStats(ctx context.Context, gameID int64) (*model.GameBehaviorStats, error)

	// 用户偏好分析
	GetUserPreferences(ctx context.Context, userID int64) (*model.UserPreferences, error)
	GetUserFavoriteCategories(ctx context.Context, userID int64) ([]*model.Category, error)
	GetUserFavoriteTags(ctx context.Context, userID int64) ([]*model.Tag, error)

	// 行为统计
	GetUserActivityStats(ctx context.Context, userID int64, days int) (*model.UserActivityStats, error)
	GetGamePopularityStats(ctx context.Context, gameID int64, days int) (*model.GamePopularityStats, error)
}

var localUserBehavior IUserBehavior

func UserBehavior() IUserBehavior {
	if localUserBehavior == nil {
		panic("implement not found for interface IUserBehavior, forgot register?")
	}
	return localUserBehavior
}

func RegisterUserBehavior(i IUserBehavior) {
	localUserBehavior = i
}
