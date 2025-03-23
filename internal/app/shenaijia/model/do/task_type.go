// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TaskType is the golang structure of table task_type for DAO operations like Where/Data.
type TaskType struct {
	g.Meta `orm:"table:task_type, do:true"`
	Type   interface{} // 类型
	Name   interface{} // 类型名
}
