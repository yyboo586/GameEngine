package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameTagDao is the data access object for table t_game_tag.
type GameTagDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns GameTagColumns // columns contains all the column names of Table for convenient usage.
}

// GameTagColumns defines and stores column names for table t_game_tag.
type GameTagColumns struct {
	ID         string // 主键
	GameID     string // 游戏ID
	TagID      string // 标签ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// gameTagColumns holds the columns for table t_game_tag.
var gameTagColumns = GameTagColumns{
	ID:         "id",
	GameID:     "game_id",
	TagID:      "tag_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewGameTagDao creates and returns a new DAO object for table data access.
func NewGameTagDao() *GameTagDao {
	return &GameTagDao{
		group:   "default",
		table:   "t_game_tag",
		columns: gameTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameTagDao) Columns() GameTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
