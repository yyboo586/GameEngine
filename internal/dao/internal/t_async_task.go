package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AsyncTaskDao is the data access object for table t_async_task.
type AsyncTaskDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AsyncTaskColumns // columns contains all the column names of Table for convenient usage.
}

// AsyncTaskColumns defines and stores column names for table t_async_task.
type AsyncTaskColumns struct {
	ID            string // 主键
	CustomID      string // 自定义任务ID
	TaskType      string // 任务类型
	Status        string // 任务状态
	RetryCount    string // 重试次数
	Content       string // 任务内容
	Version       string // 版本标识
	NextRetryTime string // 下次处理时间
	CreateTime    string // 创建时间
	UpdateTime    string // 更新时间
}

// asyncTaskColumns holds the columns for table t_async_task.
var asyncTaskColumns = AsyncTaskColumns{
	ID:            "id",
	CustomID:      "custom_id",
	TaskType:      "task_type",
	Status:        "status",
	RetryCount:    "retry_count",
	Content:       "content",
	Version:       "version",
	NextRetryTime: "next_retry_time",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// NewAsyncTaskDao creates and returns a new DAO object for table data access.
func NewAsyncTaskDao() *AsyncTaskDao {
	return &AsyncTaskDao{
		group:   "default",
		table:   "t_async_task",
		columns: asyncTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AsyncTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AsyncTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AsyncTaskDao) Columns() AsyncTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current DAO.
func (dao *AsyncTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AsyncTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AsyncTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
