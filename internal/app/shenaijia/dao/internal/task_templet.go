// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskTempletDao is the data access object for the table task_templet.
type TaskTempletDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns TaskTempletColumns // columns contains all the column names of Table for convenient usage.
}

// TaskTempletColumns defines and stores column names for the table task_templet.
type TaskTempletColumns struct {
	Id   string //
	Name string //
	Type string //
}

// taskTempletColumns holds the columns for the table task_templet.
var taskTempletColumns = TaskTempletColumns{
	Id:   "id",
	Name: "name",
	Type: "type",
}

// NewTaskTempletDao creates and returns a new DAO object for table data access.
func NewTaskTempletDao() *TaskTempletDao {
	return &TaskTempletDao{
		group:   "default",
		table:   "task_templet",
		columns: taskTempletColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskTempletDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskTempletDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskTempletDao) Columns() TaskTempletColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskTempletDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskTempletDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskTempletDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
