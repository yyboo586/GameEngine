package logics

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"sync"
	"time"
)

var (
	userBehavierOnce     sync.Once
	userBehavierInstance *userBehavier
)

type userBehavier struct {
}

func NewUserBehavier() service.IUserBehavior {
	userBehavierOnce.Do(func() {
		userBehavierInstance = &userBehavier{}
	})
	return userBehavierInstance
}

// 添加搜索历史
func (df *userBehavier) AddSearchHistory(ctx context.Context, userID int64, keyword string, resultCount int) error {
	// 记录搜索历史
	_, err := dao.SearchHistory.Ctx(ctx).Data(map[string]interface{}{
		dao.SearchHistory.Columns().UserID:        userID,
		dao.SearchHistory.Columns().SearchKeyword: keyword,
		dao.SearchHistory.Columns().ResultCount:   resultCount,
	}).Insert()

	return err
}

// 获取搜索历史
func (df *userBehavier) GetSearchHistory(ctx context.Context, userID int64, limit int) ([]*model.SearchHistory, error) {
	if limit == 0 {
		limit = 20
	}

	var entities []*entity.SearchHistory
	err := dao.SearchHistory.Ctx(ctx).
		Where(dao.SearchHistory.Columns().UserID, userID).
		OrderDesc(dao.SearchHistory.Columns().SearchTime).
		Limit(limit).
		Scan(&entities)
	if err != nil {
		return nil, err
	}

	var result []*model.SearchHistory
	for _, entity := range entities {
		result = append(result, df.convertSearchHistoryEntityToModel(entity))
	}

	return result, nil
}

// 清空搜索历史
func (df *userBehavier) ClearSearchHistory(ctx context.Context, userID int64) error {
	_, err := dao.SearchHistory.Ctx(ctx).
		Where(dao.SearchHistory.Columns().UserID, userID).
		Delete()

	return err
}

// 获取热门搜索关键词
func (df *userBehavier) GetPopularSearchKeywords(ctx context.Context, days int, limit int) ([]string, error) {
	if limit == 0 {
		limit = 10
	}

	// 统计最近7天的热门搜索词
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	var keywords []string
	err := dao.SearchHistory.Ctx(ctx).
		Fields("search_keyword, COUNT(*) as count").
		Where(dao.SearchHistory.Columns().SearchTime+" >= ?", sevenDaysAgo).
		Group("search_keyword").
		OrderDesc("count").
		Limit(limit).
		Scan(&keywords)

	if err != nil {
		return nil, err
	}

	return keywords, nil
}

// 记录游戏行为
func (df *userBehavier) RecordGameBehavior(ctx context.Context, userID int64, gameID int64, behaviorType model.BehaviorType, ipAddress string) error {
	_, err := dao.GameBehavior.Ctx(ctx).Data(map[string]interface{}{
		dao.GameBehavior.Columns().UserID:       userID,
		dao.GameBehavior.Columns().GameID:       gameID,
		dao.GameBehavior.Columns().BehaviorType: int(behaviorType),
		dao.GameBehavior.Columns().IPAddress:    ipAddress,
	}).Insert()

	return err
}

// 获取用户游戏历史
func (df *userBehavier) GetUserGameHistory(ctx context.Context, userID int64, limit int) ([]*model.GameBehavior, error) {
	if limit == 0 {
		limit = 50
	}

	var entities []*entity.GameBehavior
	err := dao.GameBehavior.Ctx(ctx).
		Where(dao.GameBehavior.Columns().UserID, userID).
		OrderDesc(dao.GameBehavior.Columns().BehaviorTime).
		Limit(limit).
		Scan(&entities)

	if err != nil {
		return nil, err
	}

	var result []*model.GameBehavior
	for _, entity := range entities {
		result = append(result, df.convertGameBehaviorEntityToModel(entity))
	}

	return result, nil
}

func (df *userBehavier) GetGameBehaviorStats(ctx context.Context, gameID int64) (out *model.GameBehaviorStats, err error) {
	return nil, nil
}

// 用户偏好分析
func (df *userBehavier) GetUserPreferences(ctx context.Context, userID int64) (out *model.UserPreferences, err error) {
	return nil, nil
}

func (df *userBehavier) GetUserFavoriteCategories(ctx context.Context, userID int64) (out []*model.Category, err error) {
	return nil, nil
}

func (df *userBehavier) GetUserFavoriteTags(ctx context.Context, userID int64) (out []*model.Tag, err error) {
	return nil, nil
}

// 行为统计
func (df *userBehavier) GetUserActivityStats(ctx context.Context, userID int64, days int) (out *model.UserActivityStats, err error) {
	return nil, nil
}

func (df *userBehavier) GetGamePopularityStats(ctx context.Context, gameID int64, days int) (out *model.GamePopularityStats, err error) {
	return nil, nil
}

// 转换实体到模型
func (df *userBehavier) convertSearchHistoryEntityToModel(in *entity.SearchHistory) *model.SearchHistory {
	return &model.SearchHistory{
		ID:            in.ID,
		UserID:        in.UserID,
		SearchKeyword: in.SearchKeyword,
		SearchTime:    in.SearchTime,
		ResultCount:   in.ResultCount,
	}
}

func (df *userBehavier) convertGameBehaviorEntityToModel(in *entity.GameBehavior) *model.GameBehavior {
	return &model.GameBehavior{
		ID:           in.ID,
		UserID:       in.UserID,
		GameID:       in.GameID,
		BehaviorType: model.BehaviorType(in.BehaviorType),
		BehaviorTime: in.BehaviorTime,
		IPAddress:    in.IPAddress,
	}
}
