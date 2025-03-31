// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectDao is the data access object for the table project.
type ProjectDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns ProjectColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectColumns defines and stores column names for the table project.
type ProjectColumns struct {
	Id                      string // 项目id
	Valid                   string // 是否有效
	ProjectName             string // 项目名
	CommunityId             string // 小区id
	StartDate               string // 开工日期
	EstimatedCompletionDate string // 预计完工日期
	CompletionDate          string // 完工日期
	Progress                string // 进度
	InspectionReport        string // 检查报告PDF链接
	AcceptanceReport        string // 验收报告PDF链接
	State                   string // 项目状态: 0(未审核), 1(审核通过)
	Creator                 string // 创建人
	Manager                 string // 物业经理
	Associate               string // 合伙人
}

// projectColumns holds the columns for the table project.
var projectColumns = ProjectColumns{
	Id:                      "id",
	Valid:                   "valid",
	ProjectName:             "project_name",
	CommunityId:             "community_id",
	StartDate:               "start_date",
	EstimatedCompletionDate: "estimated_completion_date",
	CompletionDate:          "completion_date",
	Progress:                "progress",
	InspectionReport:        "inspection_report",
	AcceptanceReport:        "acceptance_report",
	State:                   "state",
	Creator:                 "creator",
	Manager:                 "manager",
	Associate:               "associate",
}

// NewProjectDao creates and returns a new DAO object for table data access.
func NewProjectDao() *ProjectDao {
	return &ProjectDao{
		group:   "default",
		table:   "project",
		columns: projectColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProjectDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProjectDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProjectDao) Columns() ProjectColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProjectDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProjectDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ProjectDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
