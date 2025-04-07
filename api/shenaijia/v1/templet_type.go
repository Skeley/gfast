package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model"
)

type TempletTypeSearchReq struct {
	g.Meta   `path:"/task/templet/type" tags:"模板类型管理" method:"get" summary:"搜索模板类型"`
	Name     string `json:"name"`
	Standard *bool  `json:"standard"`

	commonApi.PageReq
}

type TempletTypeSearchRes struct {
	g.Meta `mime:"application/json"`
	List   []*model.TaskType `json:"list"`

	commonApi.ListRes
}

type TempletTypeAddReq struct {
	g.Meta   `path:"/task/templet/type" tags:"模板类型管理" method:"post" summary:"添加模板类型"`
	Name     string `json:"name"`
	Standard bool   `json:"standard" description:"是否为标准类型"`
}
type TempletTypeAddRes struct{}

type TempletTypeUpdateReq struct {
	g.Meta   `path:"/task/templet/type" tags:"模板类型管理" method:"put" summary:"更新"`
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Standard bool   `json:"standard" description:"是否为标准类型"`
}

type TempletTypeUpdateRes struct{}

type TempletTypeDeleteReq struct {
	g.Meta `path:"/task/templet/type" tags:"模板类型管理" method:"delete" summary:"删除"`
	Id     uint `json:"id"`
}
type TempletTypeDeleteRes struct{}
