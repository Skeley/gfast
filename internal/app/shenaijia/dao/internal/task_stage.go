// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskStageDao is the data access object for the table task_stage.
type TaskStageDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns TaskStageColumns // columns contains all the column names of Table for convenient usage.
}

// TaskStageColumns defines and stores column names for the table task_stage.
type TaskStageColumns struct {
	Id        string // stage id
	TaskId    string //
	TempletId string // 所属任务模板
	Name      string // 阶段名
	Icon      string // 图标
	Comment   string // 描述
	Position  string // 任务流位置
}

// taskStageColumns holds the columns for the table task_stage.
var taskStageColumns = TaskStageColumns{
	Id:        "id",
	TaskId:    "task_id",
	TempletId: "templet_id",
	Name:      "name",
	Icon:      "icon",
	Comment:   "comment",
	Position:  "position",
}

// NewTaskStageDao creates and returns a new DAO object for table data access.
func NewTaskStageDao() *TaskStageDao {
	return &TaskStageDao{
		group:   "default",
		table:   "task_stage",
		columns: taskStageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TaskStageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TaskStageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TaskStageDao) Columns() TaskStageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TaskStageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TaskStageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TaskStageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
