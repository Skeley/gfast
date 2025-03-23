package model

import "github.com/gogf/gf/v2/os/gtime"

type Project struct {
	Id                      uint        `json:"id" orm:"id" description:"项目id"`
	Name                    string      `json:"name" orm:"name"  description:"项目名"`
	CommunityName           string      `json:"communityName" orm:"community_name" description:"小区名"`
	StartDate               *gtime.Time `json:"startDate" orm:"start_date" description:"开工日期"`
	EstimatedCompletionDate *gtime.Time `json:"estimatedCompletionDate" orm:"estimated_completion_date" description:"预计完工日期"`
	CompletionDate          *gtime.Time `json:"completionDate" orm:"completion_date" description:"完工日期"`
	Progress                uint8       `json:"progress" orm:"progress" description:"进度"`
	InspectionReport        string      `json:"inspectionReport" orm:"inspection_report" description:"检查报告PDF链接"`
	AcceptanceReport        string      `json:"acceptanceReport" orm:"acceptance_report" description:"验收报告PDF链接"`
	Manager                 uint        `json:"manager" orm:"manager" description:"物业经理"`
	Associate               uint        `json:"associate" orm:"associate" description:"合伙人"`
}
