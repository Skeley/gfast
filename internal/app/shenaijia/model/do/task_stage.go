// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TaskStage is the golang structure of table task_stage for DAO operations like Where/Data.
type TaskStage struct {
	g.Meta  `orm:"table:task_stage, do:true"`
	Id      interface{} // stage id
	Type    interface{} // 所属任务类型
	Name    interface{} // 阶段名
	Icon    interface{} // 图标
	Comment interface{} // 描述
	Order   interface{} // 任务流位置
}
