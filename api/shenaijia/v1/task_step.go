/*
* @desc:任务步骤
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type TaskStep struct {
	Id                      int       `json:"id"`
	Name                    string    `json:"name"`
	Comment                 string    `json:"comment"`
	ImageList               []string  `json:"image_list"`
	State                   uint      `json:"state" dc:"0:进行中;1:已完成"`
	EstimatedCompletionDate time.Time `json:"estimatedCompletionDate"`
	CompletionDate          time.Time `json:"completionDate"`
}

type TaskListStepReq struct {
	g.Meta  `path:"/task/stage/step" tags:"步骤管理" method:"get" summary:"获取全部工作步骤"`
	StageId uint   `v:"required" json:"stageId"`
	TaskId  string `v:"required" json:"taskId"`
}

type TaskListStepRes struct {
	g.Meta   `mime:"application/json"`
	StepList []TaskStep `json:"stepList"`
}

type TaskAddStepReq struct {
	g.Meta  `path:"/task/stage/step" tags:"步骤管理" method:"post" summary:"添加工作步骤"`
	StageId uint   `v:"required" json:"stageId"`
	TaskId  string `v:"required" json:"taskId"`

	StepName                string    `v:"required" json:"stepName"`
	EstimatedCompletionDate time.Time `json:"estimatedCompletionDate"`
	Comment                 string    `v:"required" json:"comment"`
	ImageList               []string  `v:"foreach|url" json:"image_list"`
}
type TaskAddStepRes struct{}

type TaskUpdateStepReq struct {
	g.Meta                  `path:"/task/stage/step" tags:"步骤管理" method:"put" summary:"更新工作步骤"`
	StepId                  uint     `v:"required" json:"stepId"`
	StepName                string   `v:"required" json:"stepName"`
	Comment                 string   `v:"required" json:"comment"`
	State                   uint     `v:"required|in:0,1,2" dc:"0:未开始;1:进行中;2:已完成" json:"state"`
	EstimatedCompletionDate string   `json:"estimatedCompletionDate"`
	ImageList               []string `v:"foreach|url" json:"imageList"`
}
type TaskUpdateStepRes struct{}

type TaskDeleteStepReq struct {
	g.Meta `path:"/task/stage/step" tags:"步骤管理" method:"delete" summary:"删除工作步骤"`
	StepId uint `v:"required" json:"stepId"`
}

type TaskDeleteStepRes struct{}

type TaskCompleteStepReq struct {
	g.Meta `path:"/task/stage/complete_step" tags:"步骤管理" method:"put" summary:"完成步骤"`
	StepId uint `v:"required" json:"stepId"`
}

type TaskCompleteStepRes struct{}
