package controller

import (
	"context"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
)

var (
	SysTempletType = sysTempletTypeController{}
)

type sysTempletTypeController struct {
	commonController.BaseController
}

func (c *sysTempletTypeController) Search(ctx context.Context, req *v1.TempletTypeSearchReq) (res *v1.TempletTypeSearchRes, err error) {
	return service.Task().SearchType(ctx, req)
}

func (c *sysTempletTypeController) Add(ctx context.Context, req *v1.TempletTypeAddReq) (res *v1.TempletTypeAddRes, err error) {
	return service.Task().AddType(ctx, req)
}

func (c *sysTempletTypeController) Update(ctx context.Context, req *v1.TempletTypeUpdateReq) (res *v1.TempletTypeUpdateRes, err error) {
	return service.Task().UpdateType(ctx, req)
}

func (c *sysTempletTypeController) Delete(ctx context.Context, req *v1.TempletTypeDeleteReq) (res *v1.TempletTypeDeleteRes, err error) {
	return service.Task().DeleteType(ctx, req)
}
