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

type UserInfo struct {
	model.LoginUserRes
	Type string `json:"type"` // 物业 or 施工 or 合伙人
}

type LoginReq struct {
	g.Meta    `path:"/login" tags:"登陆" method:"get" summary:"登陆"`
	OpenId    string `v:"required" json:"openId"`
	LoginCode string `v:"required" json:"loginCode"`
	IV        string `v:"required" json:"iv"`
	// Tel    string `json:"tel"` // 从wechat拉取手机号码 登录
}

type LoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *UserInfo `json:"userInfo"`
	Token    string    `json:"token"`
}
