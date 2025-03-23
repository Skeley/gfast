// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TaskStep is the golang structure of table task_step for DAO operations like Where/Data.
type TaskStep struct {
	g.Meta  `orm:"table:task_step, do:true"`
	Id      interface{} // 步骤id
	StageId interface{} // stage id
	Name    interface{} // 步骤名
	Comment interface{} // 描述
	Order   interface{} // 排序位置
}
