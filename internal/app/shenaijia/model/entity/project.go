// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Project is the golang structure for table project.
type Project struct {
	Id                      uint        `json:"id"                      orm:"id"                        description:"项目id"`
	Valid                   int         `json:"valid"                   orm:"valid"                     description:"是否有效"`
	Audited                 int         `json:"audited"                 orm:"audited"                   description:""`
	ProjectName             string      `json:"projectName"             orm:"project_name"              description:"项目名"`
	CommunityId             uint        `json:"communityId"             orm:"community_id"              description:"小区id"`
	StartDate               *gtime.Time `json:"startDate"               orm:"start_date"                description:"开工日期"`
	EstimatedCompletionDate *gtime.Time `json:"estimatedCompletionDate" orm:"estimated_completion_date" description:"预计完工日期"`
	CompletionDate          *gtime.Time `json:"completionDate"          orm:"completion_date"           description:"完工日期"`
	Progress                uint        `json:"progress"                orm:"progress"                  description:"进度"`
	InspectionReport        string      `json:"inspectionReport"        orm:"inspection_report"         description:"检查报告PDF链接"`
	AcceptanceReport        string      `json:"acceptanceReport"        orm:"acceptance_report"         description:"验收报告PDF链接"`
	State                   int         `json:"state"                   orm:"state"                     description:"项目状态: 0(未审核), 1(审核通过)"`
	Creator                 int         `json:"creator"                 orm:"creator"                   description:"创建人"`
	Manager                 uint        `json:"manager"                 orm:"manager"                   description:"物业经理"`
	Associate               uint        `json:"associate"               orm:"associate"                 description:"合伙人"`
}
