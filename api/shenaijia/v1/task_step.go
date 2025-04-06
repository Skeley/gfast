/*
* @desc:任务步骤
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskAddStepReq struct {
	g.Meta                  `path:"/task/step" tags:"步骤管理" method:"post" summary:"添加工作步骤"`
	TaskId                  uint   `v:"required" json:"taskId"`
	StageId                 uint   `v:"required" json:"stageId"`
	StepName                string `v:"required" json:"stepName"`
	EstimatedCompletionDate string `json:"estimatedCompletionDate" v:"datetime"`
	Comment                 string `json:"comment"`
}
type TaskAddStepRes struct{}

type TaskUpdateStepReq struct {
	g.Meta                  `path:"/task/step" tags:"步骤管理" method:"put" summary:"更新工作步骤"`
	TaskId                  uint   `v:"required" json:"taskId"`
	StepId                  uint   `v:"required" json:"stepId"`
	StepName                string `v:"required" json:"stepName"`
	Comment                 string `json:"comment"`
	State                   *uint  `json:"state" dc:"0:未开始;1:进行中;2:已完成"`
	EstimatedCompletionDate string `v:"datetime" json:"estimatedCompletionDate"`
}
type TaskUpdateStepRes struct{}

type TaskDeleteStepReq struct {
	g.Meta `path:"/task/step" tags:"步骤管理" method:"delete" summary:"删除工作步骤"`
	StepId uint `v:"required" json:"stepId"`
}

type TaskDeleteStepRes struct{}

type TaskCompleteStepReq struct {
	g.Meta `path:"/task/step/complete" tags:"步骤管理" method:"post" summary:"完成步骤"`
	StepId uint `v:"required" json:"stepId"`
}

type TaskCompleteStepRes struct{}
