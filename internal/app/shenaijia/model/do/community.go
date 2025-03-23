// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Community is the golang structure of table community for DAO operations like Where/Data.
type Community struct {
	g.Meta        `orm:"table:community, do:true"`
	MajorId       interface{} // 主id
	MinorId       interface{} // 次id
	CommunityName interface{} // 小区名
}
