package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	SysProject = sysProjectController{}
)

type sysProjectController struct {
	commonController.BaseController
}

func (c *sysProjectController) ProjectAdd(ctx context.Context, req *v1.SysProjectAddReq) (res *v1.SysProjectAddRes, err error) {
	return service.Project().SysAdd(ctx, req)
}

func (c *sysProjectController) ProjectDelete(ctx context.Context, req *v1.SysProjectDelReq) (res *v1.SysProjectDelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *sysProjectController) ProjectList(ctx context.Context, req *v1.SysProjectListReq) (res *v1.SysProjectListRes, err error) {
	return service.Project().SysList(ctx, req)
}

func (c *sysProjectController) ProjectUpdate(ctx context.Context, req *v1.SysProjectEditReq) (res *v1.SysProjectEditRes, err error) {
	return service.Project().SysEdit(ctx, req)
}
