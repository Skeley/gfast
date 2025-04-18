/*
* @desc:用户接口
* @company:深爱家
* @Author:sk
* @Date:2025/03/12
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type LoginReq struct {
	g.Meta    `path:"/login" tags:"登陆" method:"post" summary:"登陆"`
	LoginCode string `v:"required" json:"loginCode"`
	PhoneCode string `json:"phoneCode"`
	LoginType uint   `v:"required" json:"loginType"`
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	User   *model.LoginUserRes `json:"user"`
	Token  string              `json:"token"`
}

type UserAuthReq struct {
	g.Meta    `path:"/user/auth" tags:"用户管理" method:"get" summary:"验证"`
	LoginCode string `v:"required" json:"loginCode"`
}

type UserAuthRes struct {
	g.Meta `mime:"application/json"`
	Pass   bool `json:"pass"`
}

type UserBoundCommunityReq struct {
	g.Meta `path:"/user/community" tags:"用户管理" method:"get" summary:"查询用户已绑定小区"`
	UserId int64 `p:"userId" v:"required#用户id不能为空"`
}

type UserBoundCommunityRes struct {
	g.Meta `mime:"application/json"`
	List   []*Community `json:"list"`
}

type UserBindCommunityReq struct {
	g.Meta      `path:"/user/community" tags:"用户管理" method:"put" summary:"用户绑定小区"`
	UserId      int64 `p:"userId" v:"required#用户id不能为空"`
	CommunityId uint  `p:"communityId"`
}
type UserBindCommunityRes struct{}

type UserUnbindCommunityReq struct {
	g.Meta      `path:"/user/community" tags:"用户管理" method:"delete" summary:"用户解绑小区"`
	UserId      int64  `p:"userId" v:"required#用户id不能为空"`
	CommunityId string `p:"communityId"`
}
type UserUnbindCommunityRes struct{}
