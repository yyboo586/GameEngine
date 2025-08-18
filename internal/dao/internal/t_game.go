package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameDao is the data access object for table t_game.
type GameDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns GameColumns // columns contains all the column names of Table for convenient usage.
}

// GameColumns defines and stores column names for table t_game.
type GameColumns struct {
	ID             string // 主键
	Name           string // 标签名称
	DistributeType string // 游戏类型
	Developer      string // 开发商
	Publisher      string // 发行商
	Description    string // 游戏描述
	Details        string // 游戏详情

	Status       string // 状态
	PublishTime  string // 发布时间
	ReserveCount string // 预约次数

	FavoriteCount string // 收藏次数
	RatingScore   string // 评分总分
	RatingCount   string // 评分次数
	DownloadCount string // 下载次数
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// gameColumns holds the columns for table t_game.
var gameColumns = GameColumns{
	ID:             "id",
	Name:           "name",
	DistributeType: "distribute_type",
	Developer:      "developer",
	Publisher:      "publisher",
	Description:    "description",
	Details:        "details",

	Status:       "status",
	PublishTime:  "publish_time",
	ReserveCount: "reserve_count",

	FavoriteCount: "favorite_count",
	RatingScore:   "rating_score",
	RatingCount:   "rating_count",
	DownloadCount: "download_count",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewGameTagDao creates and returns a new DAO object for table data access.
func NewGameDao() *GameDao {
	return &GameDao{
		group:   "default",
		table:   "t_game",
		columns: gameColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameDao) Columns() GameColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
