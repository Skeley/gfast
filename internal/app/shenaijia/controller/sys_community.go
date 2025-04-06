package controller

import (
	"context"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
)

var (
	SysCommunity = sysCommunityController{}
)

type sysCommunityController struct {
	commonController.BaseController
}

func (c *sysCommunityController) Search(ctx context.Context, req *v1.CommunitySearchReq) (res *v1.CommunityRes, err error) {
	return service.Community().Search(ctx, req)
}

func (c *sysCommunityController) Add(ctx context.Context, req *v1.CommunityAddReq) (res *v1.CommunityAddRes, err error) {
	return service.Community().Add(ctx, req)
}

func (c *sysCommunityController) Update(ctx context.Context, req *v1.CommunityUpdateReq) (res *v1.CommunityUpdateRes, err error) {
	return service.Community().Update(ctx, req)
}
