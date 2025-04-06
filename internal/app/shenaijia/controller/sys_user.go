package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"

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
