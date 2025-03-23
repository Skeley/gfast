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

type SysProjectListReq struct {
	g.Meta `path:"/project" tags:"工程项目管理" method:"get" summary:"获取项目信息"`
	commonApi.PageReq

	CommunityMajorId    uint     `p:"communityMajorId"`
	CommunityMinorId    uint     `p:"communityMinorId"`
	StartDateRange      []string `p:"startDateRange" v:"foreach|date-format:Y-m-d"`
	CompletionDateRange []string `p:"completionDateRange" v:"foreach|date-format:Y-m-d"`
	Manager             uint     `p:"manager"`
	Associate           uint     `p:"associate"`
}

type SysProjectListRes struct {
	g.Meta      `mime:"application/json"`
	ProjectList []*model.Project
	commonApi.ListRes
}

type SysProjectAddReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"post" summary:"添加项目"`
	Name                    string `p:"name"`
	CommunityMajorId        uint   `p:"communityMajorId" v:"required"`
	CommunityMinorId        uint   `p:"communityMinorId" v:"required"`
	StartDate               string `p:"startDate" v:"date-format"`
	EstimatedCompletionDate string `p:"estimatedCompletionDate" v:"date-format:Y-m-d|after:StartDate"`
	CompletionDate          string `p:"completionDate" v:"date-format:Y-m-d|after:StartDate"`
	InspectionReport        string `p:"inspectionReport"`
	AcceptanceReport        string `p:"acceptanceReport"`
	Progress                uint   `v:"required|between:0,100" json:"progress"`
	Manager                 uint   `p:"manager" v:"required"`
	Associate               uint   `p:"associate"`
}

type SysProjectAddRes struct{}

type SysProjectEditReq struct {
	g.Meta                  `path:"/project" tags:"工程项目管理" method:"put" summary:"更新项目信息"`
	ProjectId               int64  `p:"projectId" v:"required|min:1#主键ID不能为空|主键ID必须为大于0的值"`
	Name                    string `p:"name"`
	CommunityMajorId        uint   `p:"communityMajorId" v:"required"`
	CommunityMinorId        uint   `p:"communityMinorId" v:"required"`
	StartDate               string `p:"startDate" v:"date-format"`
	EstimatedCompletionDate string `p:"estimatedCompletionDate" v:"date-format:Y-m-d|after:StartDate"`
	CompletionDate          string `p:"completionDate" v:"date-format:Y-m-d|after:StartDate"`
	InspectionReport        string `p:"inspectionReport"`
	AcceptanceReport        string `p:"acceptanceReport"`
	Progress                uint   `v:"required|between:0,100" json:"progress"`
	Manager                 uint   `p:"manager" v:"required"`
	Associate               uint   `p:"associate"`
}

type SysProjectEditRes struct{}

type SysProjectDelReq struct {
	g.Meta    `path:"/project/delete" tags:"工程项目管理" method:"delete" summary:"删除项目信息"`
	ProjectId uint `v:"required" json:"projectId"`
}
type SysProjectDelRes struct{}
