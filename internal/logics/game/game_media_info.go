package game

import (
	"GameEngine/internal/dao"
	"GameEngine/internal/model"
	"GameEngine/internal/model/entity"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// 游戏媒体信息相关方法
// TODO: 游戏图标应该只有一个。
func (gg *Game) AddMediaInfo(ctx context.Context, mediaInfo *model.GameMediaInfo) (err error) {
	_, err = dao.GameMediaInfo.Ctx(ctx).Data(map[string]interface{}{
		dao.GameMediaInfo.Columns().GameID:    mediaInfo.GameID,
		dao.GameMediaInfo.Columns().FileID:    mediaInfo.FileID,
		dao.GameMediaInfo.Columns().MediaType: mediaInfo.MediaType,
		dao.GameMediaInfo.Columns().MediaUrl:  mediaInfo.MediaUrl,
		dao.GameMediaInfo.Columns().Status:    mediaInfo.Status,
	}).Insert()

	return
}

// TODO: 先删除，再插入? 还是对比差异，只更新差异部分？
func (gg *Game) UpdateMediaInfoByGameID(ctx context.Context, gameID int64, mediaInfos []*model.GameMediaInfo) (err error) {
	// 1. 查询当前游戏的所有媒体文件
	oldMediaInfos, err := gg.GetMediaInfo(ctx, gameID)
	if err != nil {
		return err
	}

	oldMediaInfosMap := make(map[string]*model.GameMediaInfo, len(oldMediaInfos))
	for _, mediaInfo := range oldMediaInfos {
		oldMediaInfosMap[mediaInfo.FileID] = mediaInfo
	}
	newMediaInfosMap := make(map[string]*model.GameMediaInfo, len(mediaInfos))
	for _, mediaInfo := range mediaInfos {
		newMediaInfosMap[mediaInfo.FileID] = mediaInfo
	}

	toRemove := make([]string, 0)
	for fileID := range oldMediaInfosMap {
		if _, exists := newMediaInfosMap[fileID]; !exists {
			toRemove = append(toRemove, fileID)
			delete(oldMediaInfosMap, fileID)
		}
	}
	toUpdate := make([]string, 0)
	for fileID, v := range oldMediaInfosMap {
		if v.MediaUrl == "" {
			toUpdate = append(toUpdate, fileID)
		}
	}

	// 3. 执行数据库操作
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GameMediaInfo.Ctx(ctx).TX(tx).
			Where(dao.GameMediaInfo.Columns().GameID, gameID).
			WhereIn(dao.GameMediaInfo.Columns().FileID, toRemove).
			Delete()
		if err != nil {
			return err
		}

		if len(toUpdate) > 0 {
			for _, fileID := range toUpdate {
				v := newMediaInfosMap[fileID]
				if v != nil {
					dataUpdate := map[string]interface{}{
						dao.GameMediaInfo.Columns().MediaType: v.MediaType,
						dao.GameMediaInfo.Columns().MediaUrl:  v.MediaUrl,
					}
					_, err = dao.GameMediaInfo.Ctx(ctx).TX(tx).
						Where(dao.GameMediaInfo.Columns().GameID, gameID).
						Where(dao.GameMediaInfo.Columns().FileID, fileID).
						Data(dataUpdate).
						Update()
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return
}

func (gg *Game) UpdateMediaInfoStatusByFileID(ctx context.Context, fileID string, status model.GameMediaStatus) (err error) {
	_, err = dao.GameMediaInfo.Ctx(ctx).Where(dao.GameMediaInfo.Columns().FileID, fileID).Data(map[string]interface{}{
		dao.GameMediaInfo.Columns().Status: status,
	}).Update()
	return
}

func (gg *Game) GetMediaInfo(ctx context.Context, gameID int64) (out []*model.GameMediaInfo, err error) {
	var entities []*entity.GameMediaInfo
	err = dao.GameMediaInfo.Ctx(ctx).Where(dao.GameMediaInfo.Columns().GameID, gameID).Scan(&entities)
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		out = append(out, gg.convertMediaInfoEntityToModel(entity))
	}
	return
}

func (gg *Game) CheckMediaInfo(ctx context.Context, gameInfo *model.Game) (err error) {
	var exists bool
	exists, err = dao.GameMediaInfo.Ctx(ctx).
		Where(dao.GameMediaInfo.Columns().GameID, gameInfo.ID).
		Where(dao.GameMediaInfo.Columns().MediaType, model.GameMediaTypeIcon).
		Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("图标媒体文件不存在")
	}

	exists, err = dao.GameMediaInfo.Ctx(ctx).
		Where(dao.GameMediaInfo.Columns().GameID, gameInfo.ID).
		Where(dao.GameMediaInfo.Columns().MediaType, model.GameMediaTypeScreenshot).
		Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("截图媒体文件不存在")
	}

	exists, err = dao.GameMediaInfo.Ctx(ctx).
		Where(dao.GameMediaInfo.Columns().GameID, gameInfo.ID).
		Where(dao.GameMediaInfo.Columns().MediaType, model.GameMediaTypeVideo).
		Exist()
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("视频媒体文件不存在")
	}
	if !exists {
		return fmt.Errorf("视频媒体文件不存在")
	}

	switch gameInfo.DistributeType {
	case model.GameDistributeTypeAPK:
		exists, err = dao.GameMediaInfo.Ctx(ctx).
			Where(dao.GameMediaInfo.Columns().GameID, gameInfo.ID).
			Where(dao.GameMediaInfo.Columns().MediaType, model.GameMediaTypeApkFile).
			Exist()
		if !exists {
			return fmt.Errorf("APK媒体文件不存在")
		}
	case model.GameDistributeTypeLink:
		exists, err = dao.GameMediaInfo.Ctx(ctx).
			Where(dao.GameMediaInfo.Columns().GameID, gameInfo.ID).
			Where(dao.GameMediaInfo.Columns().MediaType, model.GameMediaTypeH5Link).
			Exist()
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("H5链接媒体文件不存在")
		}
	}

	return nil
}

// SetH5Link 设置H5链接
func (gg *Game) SetH5Link(ctx context.Context, gameID int64, link string) error {
	_, err := dao.GameMediaInfo.Ctx(ctx).
		Data(map[string]interface{}{
			dao.GameMediaInfo.Columns().GameID:    gameID,
			dao.GameMediaInfo.Columns().FileID:    "",
			dao.GameMediaInfo.Columns().MediaType: model.GameMediaTypeH5Link,
			dao.GameMediaInfo.Columns().MediaUrl:  link,
			dao.GameMediaInfo.Columns().Status:    model.GameMediaStatusSuccess,
		}).Insert()
	return err
}

func (gg *Game) convertMediaInfoEntityToModel(in *entity.GameMediaInfo) (out *model.GameMediaInfo) {
	out = &model.GameMediaInfo{
		ID:        in.ID,
		GameID:    in.GameID,
		FileID:    in.FileID,
		MediaType: model.GameMediaType(in.MediaType),
		MediaUrl:  in.MediaUrl,
		Status:    model.GameMediaStatus(in.Status),
	}
	return
}
