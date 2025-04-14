// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	Id            string //
	UserName      string // 用户名
	Mobile        string // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	UserNickname  string // 用户昵称
	Openid        string //
	Unionid       string //
	Birthday      string // 生日
	UserPassword  string // 登录密码;cmf_password加密
	UserSalt      string // 加密盐
	UserStatus    string // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string // 用户登录邮箱
	Sex           string // 性别;0:保密,1:男,2:女
	Avatar        string // 用户头像
	DeptId        string // 部门id
	UserTypes     string // 用户类型, 1(物业经理), 2(施工经理), 3(合伙人)
	Remark        string // 备注
	IsAdmin       string // 是否后台管理员 1 是  0   否
	Address       string // 联系地址
	Describe      string // 描述信息
	LastLoginIp   string // 最后登录ip
	LastLoginTime string // 最后登录时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	Id:            "id",
	UserName:      "user_name",
	Mobile:        "mobile",
	UserNickname:  "user_nickname",
	Openid:        "openid",
	Unionid:       "unionid",
	Birthday:      "birthday",
	UserPassword:  "user_password",
	UserSalt:      "user_salt",
	UserStatus:    "user_status",
	UserEmail:     "user_email",
	Sex:           "sex",
	Avatar:        "avatar",
	DeptId:        "dept_id",
	UserTypes:     "user_types",
	Remark:        "remark",
	IsAdmin:       "is_admin",
	Address:       "address",
	Describe:      "describe",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
