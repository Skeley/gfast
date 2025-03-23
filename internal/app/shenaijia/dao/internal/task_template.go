// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskTemplateDao is the data access object for the table task_template.
type TaskTemplateDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of the current DAO.
	columns TaskTemplateColumns // columns contains all the column names of Table for convenient usage.
}

// TaskTemplateColumns defines and stores column names for the table task_template.
type TaskTemplateColumns struct {
	Type string // 类型
	Name string // 默认项目名
}

// taskTemplateColumns holds the columns for the table task_template.
var taskTemplateColumns = TaskTemplateColumns{
	Type: "type",
	Name: "name",
}

// NewTaskTemplateDao creates and returns a new DAO object for table data access.
func NewTaskTemplateDao() *TaskTemplateDao {
	return &TaskTemplateDao{
		group:   "default",
		table:   "task_template",
		columns: taskTemplateColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskTemplateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskTemplateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskTemplateDao) Columns() TaskTemplateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskTemplateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskTemplateDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskTemplateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
