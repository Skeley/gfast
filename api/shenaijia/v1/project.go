/*
* @desc:项目接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/11
 */

package v1

import (
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProjectListReq struct {
	g.Meta   `path:"/project/list" tags:"工程项目管理" method:"get" summary:"获取用户所有项目信息"`
	UserType uint `v:"required" json:"userType"`
	commonApi.PageReq
}

type ProjectListRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ProjectRes `json:"list"`
	commonApi.ListRes
}

type ProjectAddReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"post" summary:"创建项目"`
	CommunityId             uint   `v:"required" json:"communityId"`
	ProjectName             string `v:"required" json:"projectName"`
	StartDate               string `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate string `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
}

type ProjectAddRes struct{}

type ProjectUpdateReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"put" summary:"更新项目"`
	ProjectId               uint   `v:"required" json:"projectId"`
	ProjectName             string `v:"required" json:"projectName"`
	StartDate               string `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate string `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
	CompletionDate          string `v:"date-format:Y-m-d" json:"completionDate"`
	Progress                uint8  `v:"required|between:0,100" json:"progress"`
	InspectionReport        string `v:"url" json:"inspectionReport"`
	AcceptanceReport        string `v:"url" json:"acceptanceReport"`
}

type ProjectUpdateRes struct{}

type ProjectDeleteReq struct {
	g.Meta    `path:"/project" tags:"工程项目管理" method:"delete" summary:"软删除项目"`
	ProjectId uint `v:"required" json:"projectId"`
}

type ProjectDeleteRes struct{}
