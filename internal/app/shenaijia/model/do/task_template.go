// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TaskTemplate is the golang structure of table task_template for DAO operations like Where/Data.
type TaskTemplate struct {
	g.Meta `orm:"table:task_template, do:true"`
	Type   interface{} // 类型
	Name   interface{} // 默认项目名
}
