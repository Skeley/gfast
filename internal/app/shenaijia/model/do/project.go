// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Project is the golang structure of table project for DAO operations like Where/Data.
type Project struct {
	g.Meta                  `orm:"table:project, do:true"`
	Id                      interface{} // 项目id
	Valid                   interface{} // 是否有效
	Name                    interface{} // 项目名
	CommunityMajorId        interface{} // 小区主id
	CommunityMinorId        interface{} // 小区次id
	StartDate               *gtime.Time // 开工日期
	EstimatedCompletionDate *gtime.Time // 预计完工日期
	CompletionDate          *gtime.Time // 完工日期
	Progress                interface{} // 进度
	InspectionReport        interface{} // 检查报告PDF链接
	AcceptanceReport        interface{} // 验收报告PDF链接
	Manager                 interface{} // 物业经理
	Associate               interface{} // 合伙人
}
