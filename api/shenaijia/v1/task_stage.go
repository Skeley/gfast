/*
* @desc:任务阶段
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
)

type TaskListStageReq struct {
	g.Meta    `path:"/task/stage" tags:"模板任务阶段管理" method:"get" summary:"获取模板的任务阶段"`
	TempletId uint `v:"required" json:"templetId"`
}

type TaskListStageRes struct {
	g.Meta `mime:"application/json"`
	List   []model.TaskStage `json:"list"`
}

type TaskAddStageReq struct {
	g.Meta    `path:"/task/stage" tags:"模板任务阶段管理" method:"post" summary:"模版下添加任务阶段"`
	TempletId uint   `v:"required" json:"templetId"`
	Name      string `v:"required" json:"name"`
	Icon      string `v:"required|url" json:"icon"`
	Position  uint   `v:"required" json:"position"`
}
type TaskAddStageRes struct{}

type TaskUpdateStageReq struct {
	g.Meta  `path:"/task/stage" tags:"模板任务阶段管理" method:"put" summary:"更新模板的任务阶段"`
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
