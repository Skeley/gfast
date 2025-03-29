/*
* @desc:任务阶段
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import "github.com/gogf/gf/v2/frame/g"

type TaskStage struct {
	Id       int    `json:"id"`
	TaskType int    `json:"taskType"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
}

type TaskListStageReq struct {
	g.Meta   `path:"/task/stage" tags:"任务阶段管理" method:"get" summary:"获取任务阶段"`
	TaskType uint `v:"required" json:"taskType"`
}

type TaskListStageRes struct {
	g.Meta `mime:"application/json"`
	List   []TaskStage `json:"list"`
}

type TaskAddStageReq struct {
	g.Meta   `path:"/task/stage" tags:"任务阶段管理" method:"post" summary:"添加任务阶段"`
	TaskType uint   `v:"required" json:"taskType"`
	Name     string `v:"required" json:"name"`
	Icon     string `v:"required|url" json:"icon"`
	Order    uint   `v:"required" json:"order"`
}
type TaskAddStageRes struct{}

type TaskUpdateStageReq struct {
	g.Meta  `path:"/task/stage" tags:"任务阶段管理" method:"put" summary:"更新任务阶段"`
	StageId int    `v:"required" json:"stageId"`
	Name    string `v:"required" json:"name"`
	Icon    string `v:"required|url" json:"icon"`
}
type TaskUpdateStageRes struct{}

type TaskDeleteStageReq struct {
	g.Meta  `path:"/task/stage" tags:"任务阶段管理" method:"delete" summary:"删除任务阶段"`
	StageId int `v:"required" json:"stageId"`
}
type TaskDeleteStageRes struct{}
