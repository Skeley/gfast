/*
 * @desc:小区接口
 * @company:深爱家
 * @Author:sk
 * @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Community struct {
	Id     uint64 `json:"id"`
	Parent uint   `json:"parent"`
	Name   string `json:"name"`
}

type CommunityListAllReq struct {
	g.Meta `path:"/community/all" tags:"小区管理" method:"get" summary:"获取全部小区"`
}

type CommunityRes struct {
	g.Meta        `mime:"application/json"`
	CommunityList []*Community `json:"communityList"`
}

type CommunitySearchReq struct {
	g.Meta `path:"/community/search" tags:"小区管理" method:"get" summary:"搜索小区"`
	Name   string `json:"name"`
}

type CommunityListReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"get" summary:"获取小区列表. 如果不指定Parent则获取一级小区"`
	Parent *uint `json:"parent"`
}

type CommunityAddReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"post" summary:"添加小区"`
	Parent *uint  `json:"parent"`
	Name   string `json:"name"`
}

type CommunityAddRes struct{}

type CommunityUpdateReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"put" summary:"修改小区名"`
	Id     uint   `json:"Id"`
	Name   string `json:"name"`
}

type CommunityUpdateRes struct{}
