package controller

import (
	v1 "GameEngine/api/v1"
	"GameEngine/internal/model"
	"GameEngine/internal/service"
	"context"
)

func (c *gameController) PreUploadMediaInfo(ctx context.Context, req *v1.PreUploadMediaInfoReq) (res *v1.PreUploadMediaInfoRes, err error) {
	in := &model.PreUploadReq{
		FileName:    req.FileName,
		ContentType: req.ContentType,
		Size:        req.FileSize,
	}

	switch model.GameMediaType(req.Type) {
	case model.GameMediaTypeIcon, model.GameMediaTypeScreenshot, model.GameMediaTypeVideo:
		in.BucketID = "public-bucket"
	default:
		in.BucketID = "private-bucket"
	}

	out, err := service.FileEngine().PreUpload(ctx, in)
	if err != nil {
		return nil, err
	}

	mediaInfo := &model.GameMediaInfo{
		GameID:    req.GameID,
		FileID:    out.ID,
		MediaType: model.GameMediaType(req.Type),
		MediaUrl:  out.VisitURL,
		Status:    model.GameMediaStatusInit,
	}
	err = service.Game().AddMediaInfo(ctx, mediaInfo)
	if err != nil {
		return nil, err
	}

	res = &v1.PreUploadMediaInfoRes{
		FileID:       out.ID,
		OriginalName: out.OriginalName,
		UploadURL:    out.UploadURL,
	}
	return
}

func (c *gameController) ReportUploadResult(ctx context.Context, req *v1.ReportUploadResult) (res *v1.ReportUploadResultRes, err error) {
	err = service.FileEngine().ReportUploadResult(ctx, req.FileID, req.Success)
	if err != nil {
		return
	}

	status := model.GameMediaStatusSuccess
	if !req.Success {
		status = model.GameMediaStatusFailed
	}
	err = service.Game().UpdateMediaInfoStatusByFileID(ctx, req.FileID, status)
	if err != nil {
		return
	}

	return
}

func (c *gameController) PreDownloadMediaInfo(ctx context.Context, req *v1.PreDownloadMediaInfoReq) (res *v1.PreDownloadMediaInfoRes, err error) {
	userInfo, err := model.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	err = service.Game().Download(ctx, req.GameID, userInfo.ID)
	if err != nil {
		return nil, err
	}

	out, err := service.FileEngine().PreDownload(ctx, req.FileID)
	if err != nil {
		return nil, err
	}

	service.UserBehavior().RecordBehavior(ctx, userInfo.ID, req.GameID, model.BehaviorDownload, "")

	res = &v1.PreDownloadMediaInfoRes{
		DownloadURL: out.DownloadURL,
		ExpiresAt:   out.ExpiresAt,
		ExpiresIn:   out.ExpiresIn,
	}
	return
}

func (c *gameController) ReportDownloadResult(ctx context.Context, req *v1.ReportDownloadResult) (res *v1.ReportDownloadResultRes, err error) {

	return
}

func (c *gameController) SaveMediaInfo(ctx context.Context, req *v1.SaveMediaInfoReq) (res *v1.SaveMediaInfoRes, err error) {
	mediaInfos := make([]*model.GameMediaInfo, 0)
	for _, info := range req.MediaInfos {
		mediaInfos = append(mediaInfos, &model.GameMediaInfo{
			FileID:    info.FileID,
			MediaType: model.GameMediaType(info.MediaType),
			MediaUrl:  info.MediaUrl,
		})
	}
	err = service.Game().UpdateMediaInfoByGameID(ctx, req.GameID, mediaInfos)
	if err != nil {
		return
	}
	return
}

func (c *gameController) convertMediaInfoModelToResponse(in *model.GameMediaInfo) (out *v1.GameMediaInfo) {
	out = &v1.GameMediaInfo{
		ID:        in.ID,
		FileID:    in.FileID,
		MediaType: int(in.MediaType),
		MediaUrl:  in.MediaUrl,
		Status:    int(in.Status),
	}
	return
}
