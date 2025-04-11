/*
* @desc:项目接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/11
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
)

type SysProjectSearchReq struct {
	g.Meta `path:"/project" tags:"工程项目管理" method:"get" summary:"搜索项目信息"`
	commonApi.PageReq

	ProjectName         string   `json:"projectName"`
	CommunityId         string   `json:"communityId"`
	StartDateRange      []string `json:"startDateRange" v:"foreach|datetime"`
	CompletionDateRange []string `json:"completionDateRange" v:"foreach|datetime"`
	Creator             string   `json:"creator"`
	Manager             string   `json:"manager"`
	Associate           string   `json:"associate"`
}

type SysProjectSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.ProjectRes `json:"list"`
	commonApi.ListRes
}

type SysProjectAddReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"post" summary:"添加项目"`
	ProjectName             string `p:"projectName"`
	CommunityId             uint   `p:"communityId" v:"required"`
	StartDate               string `p:"startDate" v:"datetime"`
	EstimatedCompletionDate string `p:"estimatedCompletionDate" v:"datetime|after:StartDate"`
	InspectionReport        string `p:"inspectionReport"`
	Progress                uint   `v:"required|between:0,100" json:"progress"`
	Creator                 string `p:"creator"`
	Manager                 uint   `p:"manager" v:"required"`
}

type SysProjectAddRes struct{}

type SysProjectEditReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"put" summary:"更新项目信息"`
	ProjectId               int64  `p:"projectId" v:"required|min:1#主键ID不能为空|主键ID必须为大于0的值"`
	ProjectName             string `p:"projectName"`
	CommunityId             uint   `p:"communityId" v:"required"`
	StartDate               string `p:"startDate" v:"datetime"`
	EstimatedCompletionDate string `p:"estimatedCompletionDate" v:"datetime|after:StartDate"`
	CompletionDate          string `p:"completionDate" v:"datetime|after:StartDate"`
	InspectionReport        string `p:"inspectionReport"`
	AcceptanceReport        string `p:"acceptanceReport"`
	Progress                uint   `v:"required|between:0,100" json:"progress"`
	Creator                 string `p:"creator" v:"required"`
	Manager                 uint   `p:"manager" v:"required"`
	Associate               uint   `p:"associate"`
}

type SysProjectEditRes struct{}

type SysProjectDelReq struct {
	g.Meta    `path:"/project/delete" tags:"工程项目管理" method:"delete" summary:"删除项目信息"`
	ProjectId uint `v:"required" json:"projectId"`
}
type SysProjectDelRes struct{}

type SysProjectAuditReq struct {
	g.Meta    `path:"/project/audit" tags:"工程项目管理" method:"post" summary:"项目审核"`
	ProjectId uint   `v:"required" json:"projectId"`
	Associate string `json:"associate"`
}

type SysProjectAuditRes struct {
	g.Meta `mime:"application/json"`
	Pass   bool `json:"pass"`
}
