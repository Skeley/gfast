// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskAction is the golang structure for table task_action.
type TaskAction struct {
	Task       uint64      `json:"task"       orm:"task"        description:"任务id"`
	Step       uint        `json:"step"       orm:"step"        description:"步骤id"`
	State      int         `json:"state"      orm:"state"       description:"0: 未开始, 1: 进行中, 2: 已完成"`
	ModifyDate *gtime.Time `json:"modifyDate" orm:"modify_date" description:"修改日期"`
	Comment    string      `json:"comment"    orm:"comment"     description:"描述"`
	Images     string      `json:"images"     orm:"images"      description:"图片url数组 {josn}"`
}
