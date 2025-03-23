// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityUnion is the golang structure of table community_union for DAO operations like Where/Data.
type CommunityUnion struct {
	g.Meta           `orm:"table:community_union, do:true"`
	Id               interface{} //
	UserId           interface{} // 用户id
	CommunityMajorId interface{} // 小区主id
	CommunityMinorId interface{} // 小区次id
}
