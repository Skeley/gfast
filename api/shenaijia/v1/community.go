/*
 * @desc:小区接口
 * @company:深爱家
 * @Author:sk
 * @Date:2025/03/12
 */

package v1

import "github.com/gogf/gf/v2/frame/g"

type Community struct {
	FirstLevelId  uint   `json:"firstLevelId"`
	SecondLevelId uint   `json:"secondLevelId"`
	Name          string `json:"name"`
}

type CommunityListFirstLevelReq struct {
	g.Meta `path:"/community" tags:"小区管理" method:"get" summary:"获取一级小区"`
}
type CommunityListFirstLevelRes struct {
	g.Meta        `mime:"application/json"`
	CommunityList []*Community `json:"communityList"`
}

type CommunityListTwoLevelReq struct {
	g.Meta       `path:"/community" tags:"小区管理" method:"get" summary:"获取二级小区"`
	FirstLevelId uint `v:"required" json:"firstLevelId"`
}
type CommunityListTwoLevelRes struct {
	g.Meta        `mime:"application/json"`
	CommunityList []*Community `json:"communityList"`
}
