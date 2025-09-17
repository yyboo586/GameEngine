package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

func (c *gameController) AddGame(ctx context.Context, req *v1.CreateGameReq) (res *v1.CreateGameRes, err error) {
	id, err := service.Game().CreateGame(ctx, req)
	if err != nil {
		return
	}

	res = &v1.CreateGameRes{
		ID: id,
	}
	return
}

func (c *gameController) DeleteGame(ctx context.Context, req *v1.DeleteGameReq) (res *v1.DeleteGameRes, err error) {
	err = service.Game().DeleteGame(ctx, req.ID)
	return
}

func (c *gameController) UpdateGame(ctx context.Context, req *v1.UpdateGameReq) (res *v1.UpdateGameRes, err error) {
	err = service.Game().UpdateGame(ctx, req)
	return
}

func (c *gameController) GetGameByID(ctx context.Context, req *v1.GetGameByIDReq) (res *v1.GetGameByIDRes, err error) {
	// 获取游戏详情
	out, err := service.Game().GetGameByID(ctx, req.ID)
	if err != nil {
		return
	}
	if out == nil {
		return
	}

	res = &v1.GetGameByIDRes{}
	tmps, err := c.getGameDetails(ctx, []*model.Game{out})
	if err != nil {
		return nil, err
	}
	res.Game = tmps[0]
	return
}

func (c *gameController) ListGame(ctx context.Context, req *v1.ListGameReq) (res *v1.ListGameRes, err error) {
	outs, pageRes, err := service.Game().ListGame(ctx, req)
	if err != nil {
		return nil, err
	}

	res = &v1.ListGameRes{
		List:    make([]*v1.Game, 0, len(outs)),
		PageRes: pageRes,
	}
	res.List, err = c.getGameDetails(ctx, outs)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	return
}

func (c *gameController) SearchGameByGameName(ctx context.Context, req *v1.SearchGameByGameNameReq) (res *v1.SearchGameByGameNameRes, err error) {
	outs, _, err := service.Game().SearchGameByGameName(ctx, req.Name, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	res = &v1.SearchGameByGameNameRes{
		List: make([]*v1.Game, 0, len(outs)),
	}
	res.List, err = c.getGameDetails(ctx, outs)
	if err != nil {
		return nil, err
	}
	// 登录用户：补充是否已预约和是否已收藏标记
	err = c.setUserGameStatus(ctx, res.List)
	if err != nil {
		return nil, err
	}
	value := ctx.Value(model.UserInfoKey)
	if value != nil {
		service.UserBehavior().RecordBehavior(ctx, value.(model.User).ID, 0, model.BehaviorSearch, "", req.Name)
	}
	return
}

func (c *gameController) getGameDetails(ctx context.Context, in []*model.Game) (out []*v1.Game, err error) {
	out = make([]*v1.Game, 0, len(in))
	for _, gameInfo := range in {
		v := c.convertModelToResponse(gameInfo)

		// 获取游戏媒体信息
		mediaInfos, err := service.Game().GetMediaInfo(ctx, gameInfo.ID)
		if err != nil {
			return nil, err
		}
		v.MediaInfos = make([]*v1.GameMediaInfo, 0, len(mediaInfos))
		for _, mediaInfo := range mediaInfos {
			v.MediaInfos = append(v.MediaInfos, c.convertMediaInfoModelToResponse(mediaInfo))
		}

		// 获取游戏分类
		category, err := service.Metadata().GetCategoryByGameID(ctx, gameInfo.ID)
		if err != nil {
			return nil, err
		}
		if category != nil {
			v.Category = MetadataController.convertCategoryModelToResponse(category)
		}

		// 获取游戏标签
		tags, err := service.Metadata().GetTagsByGameID(ctx, gameInfo.ID)
		if err != nil {
			return nil, err
		}
		for _, tag := range tags {
			v.Tags = append(v.Tags, MetadataController.convertTagModelToResponse(tag))
		}

		out = append(out, v)
	}
	return
}

func (c *gameController) convertModelToResponse(in *model.Game) (out *v1.Game) {
	out = &v1.Game{
		ID:             in.ID,
		Name:           in.Name,
		DistributeType: int(in.DistributeType),
		Developer:      in.Developer,
		Publisher:      in.Publisher,
		Description:    in.Description,
		Details:        in.Details,

		Status:       int(in.Status),
		PublishTime:  in.PublishTime,
		ReserveCount: in.ReserveCount,

		AverageRating: in.AverageRating,
		RatingScore:   in.RatingScore,
		RatingCount:   in.RatingCount,
		FavoriteCount: in.FavoriteCount,
		DownloadCount: in.DownloadCount,

		CreateTime: in.CreateTime,
		UpdateTime: in.UpdateTime,
	}
	return
}
