package ranking

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
	"time"
)

// RankingLogic 榜单逻辑实现
type Ranking struct{}

// NewRanking 创建榜单逻辑实例
func NewRanking() service.IRanking {
	return &Ranking{}
}

// GetHotGames 获取热门游戏榜单
func (rl *Ranking) GetHotGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按热度分数排序
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc("(download_count * 0.5 + favorite_count * 0.3 + rating_score * 0.2)").
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetNewGames 获取新游榜单
func (rl *Ranking) GetNewGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取最近30天发布的游戏总数
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", thirtyDaysAgo).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().PublishTime+" >= ?", thirtyDaysAgo).
		OrderDesc(dao.Game.Columns().PublishTime).
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetTopRatedGames 获取高分游戏榜单
func (rl *Ranking) GetTopRatedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按平均评分排序
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		OrderDesc("(rating_score / rating_count)").
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetMostDownloadedGames 获取下载量榜单
func (rl *Ranking) GetMostDownloadedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按下载量排序
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc(dao.Game.Columns().DownloadCount).
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetMostFavoritedGames 获取收藏数榜单
func (rl *Ranking) GetMostFavoritedGames(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按收藏数排序
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		OrderDesc(dao.Game.Columns().FavoriteCount).
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetCategoryRanking 获取分类榜单
func (rl *Ranking) GetCategoryRanking(ctx context.Context, categoryID int64, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_category WHERE category_id = ?)", categoryID).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_category WHERE category_id = ?)", categoryID).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetTagRanking 获取标签榜单
func (rl *Ranking) GetTagRanking(ctx context.Context, tagID int64, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_tag WHERE tag_id = ?)", tagID).
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where("id IN (SELECT game_id FROM t_game_tag WHERE tag_id = ?)", tagID).
		OrderDesc("(download_count * 0.4 + favorite_count * 0.3 + rating_score * 0.2 + rating_count * 0.1)").
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// GetComprehensiveRanking 获取综合评分榜单
func (rl *Ranking) GetComprehensiveRanking(ctx context.Context, page, size int) ([]*v1.Game, *model.PageRes, error) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}

	// 计算偏移量
	offset := (page - 1) * size

	// 获取总数
	total, err := dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		Count()
	if err != nil {
		return nil, nil, err
	}

	// 获取游戏列表，按综合评分排序
	var games []*model.Game
	err = dao.Game.Ctx(ctx).
		Where(dao.Game.Columns().Status, model.GameStatusPublished).
		Where(dao.Game.Columns().RatingCount + " > 0").
		OrderDesc("(download_count * 0.3 + favorite_count * 0.2 + (rating_score / rating_count) * 0.5)").
		Offset(offset).
		Limit(size).
		Scan(&games)

	if err != nil {
		return nil, nil, err
	}

	// 转换为API响应格式
	var gameList []*v1.Game
	for _, game := range games {
		gameList = append(gameList, rl.convertModelToResponse(game))
	}

	// 构建分页信息
	pageRes := &model.PageRes{
		Total:       total,
		CurrentPage: page,
	}

	return gameList, pageRes, nil
}

// convertModelToResponse 将model.Game转换为v1.Game
func (rl *Ranking) convertModelToResponse(in *model.Game) *v1.Game {
	return &v1.Game{
		ID:             in.ID,
		Name:           in.Name,
		DistributeType: int(in.DistributeType),
		Developer:      in.Developer,
		Publisher:      in.Publisher,
		Description:    in.Description,
		Details:        in.Details,
		Status:         int(in.Status),
		PublishTime:    in.PublishTime,
		ReserveCount:   in.ReserveCount,
		RatingScore:    in.RatingScore,
		RatingCount:    in.RatingCount,
		FavoriteCount:  in.FavoriteCount,
		DownloadCount:  in.DownloadCount,
		CreateTime:     in.CreateTime,
		UpdateTime:     in.UpdateTime,
	}
}
