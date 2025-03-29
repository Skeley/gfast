/*
* @desc:任务接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Task struct {
	Id                      uint      `json:"id"`
	ProjectId               uint      `json:"projectId"`
	Name                    string    `json:"name"`
	Type                    uint      `json:"type"`
	StartDate               time.Time `json:"startDate"`
	EstimatedCompletionDate time.Time `json:"estimatedCompletionDate"`
	CompletionDate          time.Time `json:"completionDate"`
	Progress                uint8     `json:"progress"`
	Standard                bool      `json:"standard"`
}

type TaskSearchReq struct {
	g.Meta    `path:"/task/list" tags:"项目任务管理" method:"get" summary:"搜索项目下任务"`
	ProjectId uint   `v:"required" json:"projectId"`
	Name      string `json:"name"`
}

type TaskSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*Task `json:"list"`
}

type TaskAddReq struct {
	g.Meta                  `path:"/task" tags:"项目任务管理" method:"post" summary:"添加任务"`
	ProjectId               uint      `v:"required" json:"projectId"`
	Name                    string    `v:"required" json:"name"`
	Type                    uint      `v:"required" json:"type"`
	StartDate               time.Time `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate time.Time `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
}

type TaskAddRes struct{}

type TaskUpdateReq struct {
	g.Meta                  `path:"/task" tags:"项目任务管理" method:"put" summary:"更新任务"`
	TaskId                  uint      `v:"required" json:"taskId"`
	Name                    string    `v:"required" json:"name"`
	StartDate               time.Time `v:"required|date-format:Y-m-d" json:"startDate"`
	EstimatedCompletionDate time.Time `v:"required|date-format:Y-m-d" json:"estimatedCompletionDate"`
	CompletionDate          time.Time `v:"date-format:Y-m-d" json:"completionDate"`
	Progress                uint8     `v:"required|between:0,100" json:"progress"`
}

type TaskUpdateRes struct{}

type TaskDeleteReq struct {
	g.Meta `path:"/task" tags:"项目任务管理" method:"delete" summary:"删除任务"`
	TaskId uint `v:"required" json:"taskId"`
}

type TaskDeleteRes struct{}

type TaskListTypeReq struct {
	g.Meta `path:"/task/type/list" tags:"项目任务管理" method:"get" summary:"获取所有任务类型"`
}

type TaskListTypeRes struct {
	g.Meta   `mime:"application/json"`
	TypeList []struct {
		Id       uint   `json:"id"`
		Type     string `json:"type"`
		Standard bool   `json:"standard"`
	} `json:"typeList"`
}

type TaskAddTypeReq struct {
	g.Meta `path:"/task/type" tags:"项目任务管理" method:"post" summary:"添加任务类型"`
	Type   string `v:"required" json:"type"`
}

type TaskAddTypeRes struct{}

type TaskDeleteTypeReq struct {
	g.Meta `path:"/task/type" tags:"项目任务管理" method:"post" summary:"添加任务类型"`
	Type   string `v:"required" json:"type"`
}

type TaskDeleteTypeRes struct{}
