// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure of table task for DAO operations like Where/Data.
type Task struct {
	g.Meta                  `orm:"table:task, do:true"`
	Id                      interface{} // 任务id
	PorjectId               interface{} // 项目
	Name                    interface{} // 项目名
	Type                    interface{} // 类型
	StartDate               *gtime.Time // 开工日期
	EstimatedCompletionDate *gtime.Time // 预计完工日期
	CompletionDate          *gtime.Time // 完工日期
	Progress                interface{} // 进度
	AcceptanceReport        interface{} // 验收报告PDF链接
}
