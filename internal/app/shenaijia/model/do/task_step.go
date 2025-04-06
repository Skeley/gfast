// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskStep is the golang structure of table task_step for DAO operations like Where/Data.
type TaskStep struct {
	g.Meta                  `orm:"table:task_step, do:true"`
	Id                      interface{} // 步骤id
	TaskId                  interface{} //
	StageId                 interface{} // stage id
	Name                    interface{} // 步骤名
	State                   interface{} // 0:未开始;1:进行中;2:已完成
	EstimatedCompletionDate *gtime.Time // 预计完工日期
	CompletionDate          *gtime.Time //
	Comment                 interface{} // 描述
	Position                interface{} // 排序位置
}
