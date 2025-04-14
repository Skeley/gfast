package controller

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	sysService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
)

var (
	Login = loginController{}
)

type loginController struct {
	commonController.BaseController
}

func (c *loginController) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	sessionRsp, e := service.WeChat().Jscode2Session(ctx, req.LoginCode)
	if e != nil {
		return nil, gerror.Newf("微信接口异常: %s ", e.Error())
	}
	if req.PhoneCode != "" {
		phone, e := service.WeChat().GetPhoneNumber(ctx, req.PhoneCode)
		if e != nil {
			return nil, gerror.Newf("微信接口异常: %s ", e.Error())
		}
		user, err = sysService.SysUser().GetUserByMobile(ctx, phone)
		if err == nil {
			tctx := gctx.New()
			if e := sysService.SysUser().BindUnionId(tctx, user.Id, sessionRsp.UnionID); e != nil {
				g.Log().Errorf(tctx, "绑定uninoid异常: %s", e.Error())
			}
		}
	} else {
		user, err = sysService.SysUser().GetUserByUnionId(ctx, sessionRsp.UnionID)
	}
	if err != nil {
		// 保存登录失败的日志信息
		sysService.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
			Status:    0,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "WeChat",
		})
		err = gerror.New("登陆失败, 后段服务异常或用户信息不完整")
		return
	}
	user.UserPassword = ""
	err = sysService.SysUser().UpdateLoginInfo(ctx, user.Id, ip)
	if err != nil {
		return
	}
	// 报存登录成功的日志信息
	sysService.SysLoginLog().Invoke(gctx.New(), &model.LoginLogParams{
		Status:    1,
		Username:  user.UserName,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "WeChat",
	})
	key := fmt.Sprintf("WeChat-%s-%s", gconv.String(user.Id), gmd5.MustEncryptString(user.Mobile))
	token, err = sysService.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	res = &v1.LoginRes{
		Token: token,
		User:  user,
	}
	//用户在线状态保存
	sysService.SysUserOnline().Invoke(gctx.New(), &model.SysUserOnlineParams{
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(token),
		Token:     token,
		Username:  user.UserName,
		Ip:        ip,
	})
	return
}
