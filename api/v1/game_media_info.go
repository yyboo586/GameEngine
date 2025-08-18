package v1

import (
	"GameEngine/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type PreUploadMediaInfoReq struct {
	g.Meta `path:"/games/{game_id}/media-info/pre-upload" method:"post" tags:"Game Management/MediaInfo" summary:"Pre Upload"`
	model.Author
	GameID      int64  `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	Type        int    `json:"type" v:"required#文件类型不能为空" dc:"文件类型(1:图标,2:截图,3:视频,4:APK文件)"`
	FileName    string `json:"file_name" v:"required#文件名称不能为空" dc:"文件名称"`
	FileSize    int64  `json:"file_size" v:"required#文件大小不能为空" dc:"文件大小"`
	ContentType string `json:"content_type" v:"required#文件类型不能为空" dc:"文件类型"`
}

type PreUploadMediaInfoRes struct {
	g.Meta       `mime:"application/json"`
	FileID       string `json:"file_id" dc:"文件ID"`
	OriginalName string `json:"original_name" dc:"文件名称"`
	UploadURL    string `json:"upload_url" dc:"上传URL"`
}

type ReportUploadResult struct {
	g.Meta `mime:"application/json" path:"/games/{game_id}/media-info/upload-result" method:"post" tags:"Game Management/MediaInfo" summary:"Report Upload Result"`
	model.Author
	GameID  int64  `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	FileID  string `json:"file_id" v:"required#文件ID不能为空" dc:"文件ID"`
	Success bool   `json:"success" v:"required#上传结果不能为空" dc:"上传结果"`
}

type ReportUploadResultRes struct {
	g.Meta `mime:"application/json"`
}

type PreDownloadMediaInfoReq struct {
	g.Meta `path:"/games/{game_id}/media-info/pre-download" method:"post" tags:"Game Management/MediaInfo" summary:"Pre Download"`
	model.Author
	GameID int64  `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	FileID string `json:"fileID" v:"required#文件ID不能为空" dc:"文件ID"`
}

type PreDownloadMediaInfoRes struct {
	g.Meta      `mime:"application/json"`
	DownloadURL string `json:"download_url" dc:"下载URL"`
	ExpiresAt   string `json:"expires_at" dc:"过期时间"`
	ExpiresIn   int64  `json:"expires_in" dc:"过期时间"`
}

type ReportDownloadResult struct {
	g.Meta `mime:"application/json" path:"/games/{game_id}/media-info/download-result" method:"post" tags:"Game Management/MediaInfo" summary:"Report Download Result"`
	model.Author
	GameID  int64  `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	FileID  string `json:"file_id" v:"required#文件ID不能为空" dc:"文件ID"`
	Success bool   `json:"success" v:"required#下载结果不能为空" dc:"下载结果"`
}

type ReportDownloadResultRes struct {
	g.Meta `mime:"application/json"`
}

type SaveMediaInfoReq struct {
	g.Meta `path:"/games/{game_id}/media-info/save" method:"post" tags:"Game Management/MediaInfo" summary:"Save Game Media Info"`
	model.Author
	GameID     int64            `p:"game_id" v:"required#游戏ID不能为空" dc:"游戏ID"`
	MediaInfos []*GameMediaInfo `json:"media_infos" v:"required#媒体信息不能为空" dc:"媒体信息"`
}

type SaveMediaInfoRes struct {
	g.Meta `mime:"application/json"`
}

type GameMediaInfo struct {
	ID        int64  `json:"id" dc:"媒体信息ID"`
	FileID    string `json:"file_id" dc:"文件ID"`
	MediaType int    `json:"media_type" dc:"媒体类型"`
	MediaUrl  string `json:"media_url" dc:"媒体URL"`
	Status    int    `json:"status" dc:"媒体状态"`
}
