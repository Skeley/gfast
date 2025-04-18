package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	sysService "github.com/tiger1103/gfast/v3/internal/app/system/service"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	SysUser = sysUserController{}
)

type sysUserController struct {
	commonController.BaseController
}

func (c *sysUserController) BoundCommunity(ctx context.Context, req *v1.UserBoundCommunityReq) (res *v1.UserBoundCommunityRes, err error) {
	return service.User().BoundCommunity(ctx, req)
}

func (c *sysUserController) BindCommunity(ctx context.Context, req *v1.UserBindCommunityReq) (res *v1.UserBindCommunityRes, err error) {
	return service.User().BindCommunity(ctx, req)
}

func (c *sysUserController) UnbindCommunity(ctx context.Context, req *v1.UserUnbindCommunityReq) (res *v1.UserUnbindCommunityRes, err error) {
	return service.User().UnbindCommunity(ctx, req)
}

func (c *sysUserController) CheckIsNewUserReq(ctx context.Context, req *v1.CheckIsNewUserReq) (res *v1.CheckIsNewUserRes, err error) {
	sessionRsp, e := service.WeChat().Jscode2Session(ctx, 2, req.LoginCode)
	if e != nil {
		return nil, gerror.Newf("微信接口异常: %s ", e.Error())
	}

	isNew, e := sysService.SysUser().IsNewUser(ctx, sessionRsp.UnionID)
	if e != nil {
		return nil, e
	}
	res = &v1.CheckIsNewUserRes{
		NewUser: isNew,
	}
	return res, nil
}
