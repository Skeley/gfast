// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskActionDao is the data access object for the table task_action.
type TaskActionDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns TaskActionColumns // columns contains all the column names of Table for convenient usage.
}

// TaskActionColumns defines and stores column names for the table task_action.
type TaskActionColumns struct {
	Task       string // 任务id
	Step       string // 步骤id
	State      string // 0: 未开始, 1: 进行中, 2: 已完成
	ModifyDate string // 修改日期
	Comment    string // 描述
	Images     string // 图片url数组 {josn}
}

// taskActionColumns holds the columns for the table task_action.
var taskActionColumns = TaskActionColumns{
	Task:       "task",
	Step:       "step",
	State:      "state",
	ModifyDate: "modify_date",
	Comment:    "comment",
	Images:     "images",
}

// NewTaskActionDao creates and returns a new DAO object for table data access.
func NewTaskActionDao() *TaskActionDao {
	return &TaskActionDao{
		group:   "default",
		table:   "task_action",
		columns: taskActionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskActionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskActionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskActionDao) Columns() TaskActionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskActionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskActionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskActionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
