// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta           `orm:"table:user_info, do:true"`
	Id               interface{} // 用户id
	Type             interface{} // 用户类型
	CommunityMajorId interface{} // 小区主id
	CommunityMinorId interface{} // 小区次id
	Inviter          interface{} // 邀请人id
}
