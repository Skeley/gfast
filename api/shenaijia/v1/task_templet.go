package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
)

type TaskTempletSearchReq struct {
	g.Meta `path:"/task/templet" tags:"模板管理" method:"get" summary:"搜索模板"`
	Name   string `json:"name"`
	Type   string `json:"type"` // type id

	commonApi.PageReq
}

type TaskTempletSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.TaskTemplet `json:"list"`

	commonApi.ListRes
}

type TaskTempletAddReq struct {
	g.Meta `path:"/task/templet" tags:"模板管理" method:"post" summary:"添加模板"`
	Type   uint   `json:"type"` // type id
	Name   string `json:"name"`
}

type TaskTempletAddRes struct{}

type TaskTempletUpdateReq struct {
	g.Meta `path:"/task/templet" tags:"模板管理" method:"put" summary:"更新模板"`
	Id     uint   `json:"id"`
	Type   uint   `json:"type"` // type id
	Name   string `json:"name"`
}

type TaskTempletUpdateRes struct{}

type TaskTempletDeleteReq struct {
	g.Meta `path:"/task/templet" tags:"模板管理" method:"delete" summary:"删除模板"`
	Id     uint `json:"id"`
}

type TaskTempletDeleteRes struct{}

type TaskStep struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type TaskStage struct {
	Name  string      `json:"name"`
	Icon  string      `json:"icon"`
	Steps []*TaskStep `json:"steps"`
}

type TaskFlow struct {
	Stages []*TaskStage `json:"stages"`
}

type TaskTempletSetFlowReq struct {
	g.Meta    `path:"/task/templet/flow" tags:"模板管理" method:"post" summary:"设置流程"`
	TempletId uint     `json:"templetId"`
	Flow      TaskFlow `json:"flow"`
}

type TaskTempletSetFlowRes struct{}

type TaskTempletGetFlowReq struct {
	g.Meta    `path:"/task/templet/flow" tags:"模板管理" method:"get" summary:"获取流程"`
	TempletId uint `json:"templetId"`
}

type TaskTempletGetFlowRes struct {
	g.Meta `mime:"application/json"`

	Flow *TaskFlow `json:"flow"`
}
