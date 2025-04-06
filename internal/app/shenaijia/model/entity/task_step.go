// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TaskStep is the golang structure for table task_step.
type TaskStep struct {
	Id                      uint64      `json:"id"                      orm:"id"                        description:"步骤id"`
	TaskId                  int64       `json:"taskId"                  orm:"task_id"                   description:""`
	StageId                 uint64      `json:"stageId"                 orm:"stage_id"                  description:"stage id"`
	Name                    string      `json:"name"                    orm:"name"                      description:"步骤名"`
	State                   uint        `json:"state"                   orm:"state"                     description:"0:未开始;1:进行中;2:已完成"`
	EstimatedCompletionDate *gtime.Time `json:"estimatedCompletionDate" orm:"estimated_completion_date" description:"预计完工日期"`
	CompletionDate          *gtime.Time `json:"completionDate"          orm:"completion_date"           description:""`
	Comment                 string      `json:"comment"                 orm:"comment"                   description:"描述"`
	Description             string      `json:"description"             orm:"description"               description:""`
	Position                uint        `json:"position"                orm:"position"                  description:"排序位置"`
}
