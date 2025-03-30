/*
 * @desc:小区接口
 * @company:深爱家
 * @Author:sk
 * @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/model/entity"
)

type CommunitySearchReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"get" summary:"搜索小区"`
	commonApi.PageReq

	Parent string `json:"parent"`
	Name   string `json:"name"`
}

type CommunityRes struct {
	g.Meta `mime:"application/json"`
	List   []*entity.Community `json:"list"`

	commonApi.ListRes
}

type CommunityAddReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"post" summary:"添加小区"`
	Parent string `json:"parent"`
	Name   string `json:"name"`
}

type CommunityAddRes struct{}

type CommunityUpdateReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"put" summary:"修改小区名"`
	Id     uint64 `json:"Id"`
	Name   string `json:"name"`
}

type CommunityUpdateRes struct{}
