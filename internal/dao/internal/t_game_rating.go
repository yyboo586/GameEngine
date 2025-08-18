package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameRatingDao is the data access object for table t_game_rating.
type GameRatingDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns GameRatingColumns // columns contains all the column names of Table for convenient usage.
}

// GameRatingColumns defines and stores column names for table t_game_rating.
type GameRatingColumns struct {
	ID         string // 主键
	GameID     string // 游戏ID
	UserID     string // 用户ID
	Score      string // 评分总分
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// gameRatingColumns holds the columns for table t_game_rating.
var gameRatingColumns = GameRatingColumns{
	ID:         "id",
	GameID:     "game_id",
	UserID:     "user_id",
	Score:      "score",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewGameRatingDao creates and returns a new DAO object for table data access.
func NewGameRatingDao() *GameRatingDao {
	return &GameRatingDao{
		group:   "default",
		table:   "t_game_rating",
		columns: gameRatingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GameRatingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GameRatingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GameRatingDao) Columns() GameRatingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *GameRatingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GameRatingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GameRatingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
