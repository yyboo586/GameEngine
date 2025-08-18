package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CategoryDao is the data access object for table t_category.
type CategoryDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CategoryColumns // columns contains all the column names of Table for convenient usage.
}

// CategoryColumns defines and stores column names for table t_category.
type CategoryColumns struct {
	ID         string // 主键
	Name       string // 分类名称
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// categoryColumns holds the columns for table t_category.
var categoryColumns = CategoryColumns{
	ID:         "id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewCategoryDao creates and returns a new DAO object for table data access.
func NewCategoryDao() *CategoryDao {
	return &CategoryDao{
		group:   "default",
		table:   "t_category",
		columns: categoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CategoryDao) Columns() CategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *CategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
