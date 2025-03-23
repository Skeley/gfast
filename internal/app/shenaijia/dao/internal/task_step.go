// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskStepDao is the data access object for the table task_step.
type TaskStepDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns TaskStepColumns // columns contains all the column names of Table for convenient usage.
}

// TaskStepColumns defines and stores column names for the table task_step.
type TaskStepColumns struct {
	Id      string // 步骤id
	StageId string // stage id
	Name    string // 步骤名
	Comment string // 描述
	Order   string // 排序位置
}

// taskStepColumns holds the columns for the table task_step.
var taskStepColumns = TaskStepColumns{
	Id:      "id",
	StageId: "stage_id",
	Name:    "name",
	Comment: "comment",
	Order:   "order",
}

// NewTaskStepDao creates and returns a new DAO object for table data access.
func NewTaskStepDao() *TaskStepDao {
	return &TaskStepDao{
		group:   "default",
		table:   "task_step",
		columns: taskStepColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskStepDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskStepDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskStepDao) Columns() TaskStepColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskStepDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskStepDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskStepDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
