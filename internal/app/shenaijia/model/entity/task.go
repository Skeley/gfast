// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Task is the golang structure for table task.
type Task struct {
	Id                      uint64      `json:"id"                      orm:"id"                        description:"任务id"`
	PorjectId               uint        `json:"porjectId"               orm:"porject_id"                description:"项目"`
	Name                    string      `json:"name"                    orm:"name"                      description:"项目名"`
	Type                    uint        `json:"type"                    orm:"type"                      description:"类型"`
	StartDate               *gtime.Time `json:"startDate"               orm:"start_date"                description:"开工日期"`
	EstimatedCompletionDate *gtime.Time `json:"estimatedCompletionDate" orm:"estimated_completion_date" description:"预计完工日期"`
	CompletionDate          *gtime.Time `json:"completionDate"          orm:"completion_date"           description:"完工日期"`
	Progress                uint        `json:"progress"                orm:"progress"                  description:"进度"`
	AcceptanceReport        string      `json:"acceptanceReport"        orm:"acceptance_report"         description:"验收报告PDF链接"`
}
