/*
* @desc:耗材接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Material struct {
	Id    uint   `json:"id"`
	Type  uint   `json:"type"`
	Brand uint   `json:"brand"`
	Name  string `json:"name"`
	Image []byte `json:"image"`
	Price uint   `json:"price"`
}

type MaterialGetReq struct {
	g.Meta     `path:"/material" tags:"耗材管理" method:"get" summary:"获取耗材"`
	MaterialId uint `json:"materialId"`
}

type MaterialGetRes struct {
	g.Meta   `mime:"application/json"`
	Material *Material `json:"material"`
}

type MaterialAddReq struct {
	g.Meta `path:"/material" tags:"耗材管理" method:"post" summary:"添加耗材"`
	Type   uint   `json:"type"`
	Brand  uint   `json:"brand"`
	Name   string `json:"name"`
	Image  []byte `json:"image"`
	Price  uint   `json:"price"`
}

type MaterialAddRes struct{}

type MaterialUpdateReq struct {
	g.Meta     `path:"/material" tags:"耗材管理" method:"put" summary:"更新耗材"`
	MaterialId uint   `json:"materialId"`
	Type       uint   `json:"type"`
	Brand      uint   `json:"brand"`
	Name       string `json:"name"`
	Image      []byte `json:"image"`
	Price      uint   `json:"price"`
}

type MaterialUpdateRes struct{}

type MaterialDeleteReq struct {
	g.Meta     `path:"/material" tags:"耗材管理" method:"delete" summary:"删除耗材"`
	MaterialId uint `json:"materialId"`
}

type MaterialDeleteRes struct{}

type MaterialListTypeReq struct {
	g.Meta `path:"/material/type" tags:"耗材类型管理" method:"get" summary:"获取全部耗材类型"`
}

type MaterialListTypeRes struct {
	g.Meta   `mime:"application/json"`
	TypeList []struct {
		Id       uint   `json:"id"`
		TypeName string `json:"typeName"`
	} `json:"typeList"`
}

type MaterialAddTypeReq struct {
	g.Meta   `path:"/material/type" tags:"耗材类型管理" method:"post" summary:"添加耗材类型"`
	TypeName string `json:"typeName"`
}

type MaterialAddTypeRes struct{}

type MaterialUpdateTypeReq struct {
	g.Meta   `path:"/material/type" tags:"耗材类型管理" method:"put" summary:"修改耗材类型名"`
	TypeId   uint   `json:"typeId"`
	TypeName string `json:"typeName"`
}

type MaterialUpdateTypeRes struct{}

type MaterialDeleteTypeReq struct {
	g.Meta `path:"/material/type" tags:"耗材类型管理" method:"delete" summary:"删除耗材类型"`
	Type   uint `json:"type"`
}

type MaterialDeleteTypeRes struct{}

type MaterialSearchBrandReq struct {
	g.Meta `path:"/material/brand" tags:"耗材品牌管理" method:"get" summary:"搜索某类型全部品牌"`
	Type   uint `json:"type"`
}

type MaterialSearchBrandRes struct {
	g.Meta    `mime:"application/json"`
	BrandList []struct {
		Id        uint   `json:"id"`
		BrandName string `json:"brandName"`
	} `json:"brandList"`
}

type MaterialAddBrandReq struct {
	g.Meta    `path:"/material/brand" tags:"耗材品牌管理" method:"post" summary:"添加品牌"`
	Type      uint   `json:"type"`
	BrandName string `json:"brandName"`
}

type MaterialAddBrandRes struct{}

type MaterialUpdateBrandReq struct {
	g.Meta    `path:"/material/brand" tags:"耗材品牌管理" method:"put" summary:"修改品牌名"`
	BrandId   uint   `json:"brandId"`
	BrandName string `json:"brandName"`
}

type MaterialUpdateBrandRes struct{}

type MaterialDeleteBrandReq struct {
	g.Meta `path:"/material/brand" tags:"耗材品牌管理" method:"delete" summary:"删除品牌"`
	Brand  uint `json:"brand"`
}

type MaterialDeleteBrandRes struct{}

type MaterialSearchReq struct {
	g.Meta `path:"/material/list" tags:"耗材管理" method:"get" summary:"检索耗材"`
	Type   uint `json:"type"`
	Brand  uint `json:"brand"`
}

type MaterialSearchRes struct {
	g.Meta   `mime:"application/json"`
	Material []*Material `json:"materials"`
}
