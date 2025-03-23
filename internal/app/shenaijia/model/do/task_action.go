// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskAction is the golang structure of table task_action for DAO operations like Where/Data.
type TaskAction struct {
	g.Meta     `orm:"table:task_action, do:true"`
	Task       interface{} // 任务id
	Step       interface{} // 步骤id
	State      interface{} // 0: 未开始, 1: 进行中, 2: 已完成
	ModifyDate *gtime.Time // 修改日期
	Comment    interface{} // 描述
	Images     interface{} // 图片url数组 {josn}
}
