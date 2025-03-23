// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserType is the golang structure of table user_type for DAO operations like Where/Data.
type UserType struct {
	g.Meta `orm:"table:user_type, do:true"`
	Type   interface{} // 用户类型
	Name   interface{} // name
}
