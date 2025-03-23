// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityDao is the data access object for the table community.
type CommunityDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns CommunityColumns // columns contains all the column names of Table for convenient usage.
}

// CommunityColumns defines and stores column names for the table community.
type CommunityColumns struct {
	MajorId       string // 主id
	MinorId       string // 次id
	CommunityName string // 小区名
}

// communityColumns holds the columns for the table community.
var communityColumns = CommunityColumns{
	MajorId:       "major_id",
	MinorId:       "minor_id",
	CommunityName: "community_name",
}

// NewCommunityDao creates and returns a new DAO object for table data access.
func NewCommunityDao() *CommunityDao {
	return &CommunityDao{
		group:   "default",
		table:   "community",
		columns: communityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommunityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommunityDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommunityDao) Columns() CommunityColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommunityDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommunityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *CommunityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
