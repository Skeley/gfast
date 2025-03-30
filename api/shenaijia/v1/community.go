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
)

type Community struct {
	Id         uint   `json:"id"              description:""`
	Pid        uint   `json:"pid"             description:""`
	Name       string `json:"name"  description:"小区名"`
	ParentName string `json:"parentName" description:"上级小区名"`
}

type CommunitySearchReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"get" summary:"搜索小区"`
	commonApi.PageReq

	Pid  string `json:"parentId"`
	Name string `json:"name"`
}

type CommunityRes struct {
	g.Meta `mime:"application/json"`
	List   []*Community `json:"list"`

	commonApi.ListRes
}

type CommunityAddReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"post" summary:"添加小区"`
	Pid    string `json:"parentId"`
	Name   string `json:"name"`
}

type CommunityAddRes struct{}

type CommunityUpdateReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"put" summary:"修改小区名"`
	Id     uint64 `json:"Id"`
	Name   string `json:"name"`
	Pid    string `json:"parentId"`
}

type CommunityUpdateRes struct{}
