package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameMediaInfoDao is the data access object for table t_game_media_info.
type GameMediaInfoDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns GameMediaColumns // columns contains all the column names of Table for convenient usage.
}

// GameMediaColumns defines and stores column names for table t_game_media_info.
type GameMediaColumns struct {
	ID         string // 主键
	GameID     string // 游戏ID
	FileID     string // 文件ID
	MediaType  string // 媒体类型
	MediaUrl   string // 媒体URL
	Status     string // 状态
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// gameMediaColumns holds the columns for table t_game_media_info		.
var gameMediaColumns = GameMediaColumns{
	ID:         "id",
	GameID:     "game_id",
	FileID:     "file_id",
	MediaType:  "media_type",
	MediaUrl:   "media_url",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewGameMediaInfoDao creates and returns a new DAO object for table data access.
func NewGameMediaInfoDao() *GameMediaInfoDao {
	return &GameMediaInfoDao{
		group:   "default",
		table:   "t_game_media_info",
		columns: gameMediaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameMediaInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameMediaInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameMediaInfoDao) Columns() GameMediaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *GameMediaInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameMediaInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameMediaInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
