// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTypeDao is the data access object for the table user_type.
type UserTypeDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of the current DAO.
	columns UserTypeColumns // columns contains all the column names of Table for convenient usage.
}

// UserTypeColumns defines and stores column names for the table user_type.
type UserTypeColumns struct {
	Type string // 用户类型
	Name string // name
}

// userTypeColumns holds the columns for the table user_type.
var userTypeColumns = UserTypeColumns{
	Type: "type",
	Name: "name",
}

// NewUserTypeDao creates and returns a new DAO object for table data access.
func NewUserTypeDao() *UserTypeDao {
	return &UserTypeDao{
		group:   "default",
		table:   "user_type",
		columns: userTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserTypeDao) Columns() UserTypeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
