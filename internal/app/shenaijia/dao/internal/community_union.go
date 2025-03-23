// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityUnionDao is the data access object for the table community_union.
type CommunityUnionDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of the current DAO.
	columns CommunityUnionColumns // columns contains all the column names of Table for convenient usage.
}

// CommunityUnionColumns defines and stores column names for the table community_union.
type CommunityUnionColumns struct {
	Id               string //
	UserId           string // 用户id
	CommunityMajorId string // 小区主id
	CommunityMinorId string // 小区次id
}

// communityUnionColumns holds the columns for the table community_union.
var communityUnionColumns = CommunityUnionColumns{
	Id:               "id",
	UserId:           "user_id",
	CommunityMajorId: "community_major_id",
	CommunityMinorId: "community_minor_id",
}

// NewCommunityUnionDao creates and returns a new DAO object for table data access.
func NewCommunityUnionDao() *CommunityUnionDao {
	return &CommunityUnionDao{
		group:   "default",
		table:   "community_union",
		columns: communityUnionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommunityUnionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommunityUnionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommunityUnionDao) Columns() CommunityUnionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommunityUnionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommunityUnionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CommunityUnionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
