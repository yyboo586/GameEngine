package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameBehaviorStats 游戏行为统计
type GameBehaviorStats struct {
	GameID        int64       `json:"game_id" dc:"游戏ID"`
	ViewCount     int64       `json:"view_count" dc:"查看次数"`
	DownloadCount int64       `json:"download_count" dc:"下载次数"`
	FavoriteCount int64       `json:"favorite_count" dc:"收藏次数"`
	RatingCount   int64       `json:"rating_count" dc:"评分次数"`
	AvgRating     float64     `json:"avg_rating" dc:"平均评分"`
	LastActivity  *gtime.Time `json:"last_activity" dc:"最后活动时间"`
	CreateTime    *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime    *gtime.Time `json:"update_time" dc:"更新时间"`
}

// UserPreferences 用户偏好
type UserPreferences struct {
	UserID           int64       `json:"user_id" dc:"用户ID"`
	FavoriteCategory string      `json:"favorite_category" dc:"偏好分类"`
	FavoriteTags     []string    `json:"favorite_tags" dc:"偏好标签"`
	GamePlayStyle    string      `json:"game_play_style" dc:"游戏风格偏好"`
	LastActiveTime   *gtime.Time `json:"last_active_time" dc:"最后活跃时间"`
	CreateTime       *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime       *gtime.Time `json:"update_time" dc:"更新时间"`
}

// UserActivityStats 用户活动统计
type UserActivityStats struct {
	UserID           int64       `json:"user_id" dc:"用户ID"`
	TotalViews       int64       `json:"total_views" dc:"总查看次数"`
	TotalDownloads   int64       `json:"total_downloads" dc:"总下载次数"`
	TotalFavorites   int64       `json:"total_favorites" dc:"总收藏次数"`
	TotalRatings     int64       `json:"total_ratings" dc:"总评分次数"`
	ActiveDays       int         `json:"active_days" dc:"活跃天数"`
	LastActivityTime *gtime.Time `json:"last_activity_time" dc:"最后活动时间"`
	CreateTime       *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime       *gtime.Time `json:"update_time" dc:"更新时间"`
}

// GamePopularityStats 游戏热度统计
type GamePopularityStats struct {
	GameID           int64       `json:"game_id" dc:"游戏ID"`
	DailyViews       int64       `json:"daily_views" dc:"日查看次数"`
	DailyDownloads   int64       `json:"daily_downloads" dc:"日下载次数"`
	DailyFavorites   int64       `json:"daily_favorites" dc:"日收藏次数"`
	DailyRatings     int64       `json:"daily_ratings" dc:"日评分次数"`
	WeeklyViews      int64       `json:"weekly_views" dc:"周查看次数"`
	WeeklyDownloads  int64       `json:"weekly_downloads" dc:"周下载次数"`
	WeeklyFavorites  int64       `json:"weekly_favorites" dc:"周收藏次数"`
	WeeklyRatings    int64       `json:"weekly_ratings" dc:"周评分次数"`
	MonthlyViews     int64       `json:"monthly_views" dc:"月查看次数"`
	MonthlyDownloads int64       `json:"monthly_downloads" dc:"月下载次数"`
	MonthlyFavorites int64       `json:"monthly_favorites" dc:"月收藏次数"`
	MonthlyRatings   int64       `json:"monthly_ratings" dc:"月评分次数"`
	CreateTime       *gtime.Time `json:"create_time" dc:"创建时间"`
	UpdateTime       *gtime.Time `json:"update_time" dc:"更新时间"`
}
