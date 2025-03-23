/*
 * @desc:小区接口
 * @company:深爱家
 * @Author:sk
 * @Date:2025/03/12
 */

package v1

import "github.com/gogf/gf/v2/frame/g"

type Community struct {
	MajorId uint   `orm:"major_id" json:"majorId"`
	MinorId uint   `orm:"minor_id" json:"minorId"`
	Name    string `orm:"community_name" json:"name"`
}

type CommunityListMajorReq struct {
	g.Meta `path:"/community/major" tags:"小区管理" method:"get" summary:"获取一级小区"`
}
type CommunityListMajorRes struct {
	g.Meta        `mime:"application/json"`
	CommunityList []*Community `json:"communityList"`
}

type CommunityListMinorReq struct {
	g.Meta  `path:"/community/minor" tags:"小区管理" method:"get" summary:"获取二级小区"`
	MajorId uint `v:"required" json:"majorId"`
}
type CommunityListMinorRes struct {
	g.Meta        `mime:"application/json"`
	CommunityList []*Community `json:"communityList"`
}

type SysCommunityAddReq struct {
	g.Meta  `path:"/community" tags:"小区管理" method:"post" summary:"添加小区"`
	MajorId uint   `json:"majorId"`
	MinorId uint   `json:"minorId"`
	Name    string `json:"name"`
}

type SysCommunityAddRes struct{}

type SysCommunityUpdateReq struct {
	g.Meta  `path:"/community" tags:"小区管理" method:"put" summary:"修改小区名"`
	MajorId uint   `json:"majorId"`
	MinorId uint   `json:"minorId"`
	Name    string `json:"name"`
}

type SysCommunityUpdateRes struct{}
