/*
* @desc:项目接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/11
 */

package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Project struct {
	Id                      uint      `json:"id"`
	Name                    string    `json:"name"`
	Community               string    `json:"community"`
	StartDate               time.Time `json:"startDate"`
	EstimatedCompletionDate time.Time `json:"estimatedCompletionDate"`
	CompletionDate          time.Time `json:"completionDate"`
	Progress                uint8     `json:"progress"`
	InspectionReport        string    `json:"inspectionReport"` // url
	AcceptanceReport        string    `json:"acceptanceReport"` // url
}

type ProjectListReq struct {
	g.Meta `path:"/project/list" tags:"工程项目管理" method:"get" summary:"获取用户所有项目信息"`
	UserId uint `v:"required"  json:"userId"`
}

type ProjectListRes struct {
	g.Meta      `mime:"application/json"`
	ProjectList []*Project `json:"projectList"`
}

type ProjectGetReq struct {
	g.Meta    `path:"/project" tags:"工程项目管理" method:"get" summary:"获取项目信息"`
	ProjectId uint `v:"required" json:"projectId"`
}

type ProjectGetRes struct {
	g.Meta  `mime:"application/json"`
	Project *Project `json:"project"`
}

type ProjectAddReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"post" summary:"创建项目"`
	CommunityMajorId        uint   `v:"required" json:"communityMajorId" `
	CommunityMinorId        uint   `v:"required" json:"communityMinorId" `
	ProjectName             string `v:"required" json:"projectName"`
	StartDate               string `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate string `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
}

type ProjectAddRes struct{}

type ProjectUpdateReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"put" summary:"更新项目"`
	ProjectId               uint   `v:"required" json:"projectId"`
	ProjectName             string `v:"required" json:"projectName"`
	CommunityMajorId        uint   `v:"required" json:"communityMajorId" `
	CommunityMinorId        uint   `v:"required" json:"communityMinorId" `
	StartDate               string `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate string `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
	CompletionDate          string `v:"date-format:Y-m-d" json:"completionDate"`
	Progress                uint8  `v:"required|between:0,100" json:"progress"`
}

type ProjectUpdateRes struct{}

type ProjectDeleteReq struct {
	g.Meta    `path:"/project" tags:"工程项目管理" method:"delete" summary:"软删除项目"`
	ProjectId uint `v:"required" json:"projectId"`
}

type ProjectDeleteRes struct{}
