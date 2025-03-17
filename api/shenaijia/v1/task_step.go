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

type Action struct {
	Id         int       `json:"id"`
	Comment    string    `json:"comment"`
	ImageList  []string  `json:"image_list"`
	Stats      uint      `json:"stats"`
	ModifyDate time.Time `json:"modify_date"`
}

type TaskStep struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Type   uint    `json:"type"`
	Action *Action `json:"action"`
}

type TaskListWorkReq struct {
	g.Meta  `path:"/task/stage/work" tags:"工作管理" method:"get" summary:"获取全部工作内容"`
	StageId uint   `v:"required" json:"stageId"`
	TaskId  string `v:"required" json:"taskId"`
}

type TaskListWorkRes struct {
	g.Meta   `mime:"application/json"`
	StepList []TaskStep `json:"stepList"`
}

type TaskAddWorkReq struct {
	g.Meta    `path:"/task/stage/work" tags:"步骤管理" method:"put" summary:"添加工作内容"`
	StageId   uint     `v:"required" json:"stageId"`
	TaskId    string   `v:"required" json:"taskId"`
	Comment   string   `v:"required" json:"comment"`
	ImageList []string `v:"foreach|url" json:"image_list"`
}
type TaskAddWorkRes struct{}

type TaskUpdateWorkReq struct {
	g.Meta    `path:"/task/stage/work" tags:"步骤管理" method:"put" summary:"更新工作内容"`
	ActionId  *Action  `v:"required" json:"actionId"`
	Comment   string   `v:"required" json:"comment"`
	Stats     uint     `v:"required|in:0,1,2" dc:"0:未开始;1:进行中;2:已完成" json:"stats"`
	ImageList []string `v:"foreach|url" json:"imageList"`
}
type TaskUpdateWorkRes struct{}

type StageAddStepReq struct {
	g.Meta  `path:"/task/stage/step" tags:"步骤管理" method:"post" summary:"添加步骤"`
	StageId uint   `v:"required" json:"stageId"`
	Name    string `v:"required" json:"name"`
	Comment string `v:"required" json:"comment"`
	Order   uint   `v:"required" json:"order"`
}
type StageAddStepRes struct{}

type StageUpdateStepReq struct {
	g.Meta  `path:"/task/stage/step" tags:"步骤管理" method:"put" summary:"更新步骤"`
	StepId  uint   `v:"required" json:"stepId"`
	Name    string `v:"required" json:"name"`
	Comment string `v:"required" json:"comment"`
}
type StageUpdateStepRes struct{}

type StageDelStepReq struct {
	g.Meta `path:"/task/stage/step" tags:"步骤管理" method:"delete" summary:"删除步骤"`
	StepId uint `v:"required" json:"stepId"`
}
type StageDelStepRes struct{}
