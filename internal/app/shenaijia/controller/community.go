package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Community = communityController{}
)

type communityController struct {
	commonController.BaseController
}

func (c *communityController) ListMajor(ctx context.Context, req *v1.CommunityListMajorReq) (res *v1.CommunityListMajorRes, err error) {
	return service.Community().ListMajor(ctx, req)
}

func (c *communityController) ListMinor(ctx context.Context, req *v1.CommunityListMinorReq) (res *v1.CommunityListMinorRes, err error) {
	return service.Community().ListMinor(ctx, req)
}

func (c *communityController) SysAdd(ctx context.Context, req *v1.SysCommunityAddReq) (res *v1.SysCommunityAddRes, err error) {
	return service.Community().SysAdd(ctx, req)
}

func (c *communityController) SysUpdate(ctx context.Context, req *v1.SysCommunityUpdateReq) (res *v1.SysCommunityUpdateRes, err error) {
	return service.Community().SysUpdate(ctx, req)
}
