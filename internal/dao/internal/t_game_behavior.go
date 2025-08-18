package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameBehaviorDao is the data access object for table t_game_behavior.
type GameBehaviorDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns GameBehaviorColumns // columns contains all the column names of Table for convenient usage.
}

// GameBehaviorColumns defines and stores column names for table t_game_behavior.
type GameBehaviorColumns struct {
	ID           string // 主键
	UserID       string // 用户ID
	GameID       string // 游戏ID
	BehaviorType string // 行为类型
	BehaviorTime string // 行为时间
	IPAddress    string // IP地址
}

// gameBehaviorColumns holds the columns for table t_game_behavior.
var gameBehaviorColumns = GameBehaviorColumns{
	ID:           "id",
	UserID:       "user_id",
	GameID:       "game_id",
	BehaviorType: "behavior_type",
	BehaviorTime: "behavior_time",
	IPAddress:    "ip_address",
}

// NewGameBehaviorDao creates and returns a new DAO object for table data access.
func NewGameBehaviorDao() *GameBehaviorDao {
	return &GameBehaviorDao{
		group:   "default",
		table:   "t_game_behavior",
		columns: gameBehaviorColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameBehaviorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameBehaviorDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameBehaviorDao) Columns() GameBehaviorColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GameBehaviorDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameBehaviorDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameBehaviorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
