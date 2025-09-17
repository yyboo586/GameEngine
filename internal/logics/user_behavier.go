package logics

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"GameEngine/internal/service"
	"context"
	"sync"
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

// 记录游戏行为
func (df *userBehavier) RecordBehavior(ctx context.Context, userID int64, gameID int64, behaviorType model.BehaviorType, ipAddress string, searchKeyword string) error {
	_, err := dao.UserBehavior.Ctx(ctx).Data(map[string]interface{}{
		dao.UserBehavior.Columns().UserID:        userID,
		dao.UserBehavior.Columns().GameID:        gameID,
		dao.UserBehavior.Columns().BehaviorType:  behaviorType,
		dao.UserBehavior.Columns().IPAddress:     ipAddress,
		dao.UserBehavior.Columns().SearchKeyword: searchKeyword,
	}).Insert()

	return err
}

// 获取搜索历史
func (df *userBehavier) GetSearchHistory(ctx context.Context, userID int64, pageReq *model.PageReq) (outs []*model.UserBehavior, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	query := dao.UserBehavior.Ctx(ctx).
		Where(dao.UserBehavior.Columns().UserID, userID)

	total, err := query.Count()
	if err != nil {
		return
	}

	var entities []*entity.UserBehavior
	err = query.
		Where(dao.UserBehavior.Columns().BehaviorType, model.BehaviorSearch).
		OrderDesc(dao.UserBehavior.Columns().BehaviorTime).
		Page(pageReq.Page, pageReq.Size).
		Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		outs = append(outs, df.convertUserBehaviorEntityToModel(entity))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

// 清空搜索历史
func (df *userBehavier) ClearSearchHistory(ctx context.Context, userID int64) error {
	_, err := dao.UserBehavior.Ctx(ctx).
		Where(dao.UserBehavior.Columns().UserID, userID).
		Delete()

	return err
}

// 获取玩过游戏历史
func (df *userBehavier) GetPlayHistory(ctx context.Context, userID int64, pageReq *model.PageReq) (outs []*model.UserBehavior, pageRes *model.PageRes, err error) {
	if pageReq.Page == 0 {
		pageReq.Page = 1
	}
	if pageReq.Size == 0 {
		pageReq.Size = 10
	}

	// 先获取去重后的游戏总数（按game_id分组）
	totalQuery := dao.UserBehavior.Ctx(ctx).
		Where(dao.UserBehavior.Columns().UserID, userID).
		Where(dao.UserBehavior.Columns().BehaviorType, model.BehaviorPlay).
		Group(dao.UserBehavior.Columns().GameID)

	total, err := totalQuery.Count()
	if err != nil {
		return
	}

	// 获取去重后的游戏记录（每个游戏只保留最新的一条记录）
	// 使用子查询的方式获取每个游戏的最新记录
	var entities []*entity.UserBehavior
	err = dao.UserBehavior.Ctx(ctx).
		Where(dao.UserBehavior.Columns().UserID, userID).
		Where(dao.UserBehavior.Columns().BehaviorType, model.BehaviorPlay).
		Where("(game_id, behavior_time) IN (SELECT game_id, MAX(behavior_time) FROM t_user_behavior WHERE user_id = ? AND behavior_type = ? GROUP BY game_id)", userID, model.BehaviorPlay).
		OrderDesc(dao.UserBehavior.Columns().BehaviorTime).
		Page(pageReq.Page, pageReq.Size).
		Scan(&entities)
	if err != nil {
		return
	}

	for _, entity := range entities {
		outs = append(outs, df.convertUserBehaviorEntityToModel(entity))
	}
	pageRes = &model.PageRes{
		Total:       total,
		CurrentPage: pageReq.Page,
	}
	return
}

/*
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
*/

// 转换实体到模型
func (df *userBehavier) convertUserBehaviorEntityToModel(in *entity.UserBehavior) *model.UserBehavior {
	return &model.UserBehavior{
		ID:            in.ID,
		UserID:        in.UserID,
		GameID:        in.GameID,
		SearchKeyword: in.SearchKeyword,
		BehaviorType:  model.BehaviorType(in.BehaviorType),
		BehaviorTime:  in.BehaviorTime,
		IPAddress:     in.IPAddress,
	}
}
