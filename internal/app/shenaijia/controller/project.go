package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Project = projectController{}
)

type projectController struct {
	commonController.BaseController
}

func (c *projectController) ProjectAdd(ctx context.Context, req *v1.ProjectAddReq) (res *v1.ProjectAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *projectController) ProjectDelete(ctx context.Context, req *v1.ProjectDeleteReq) (res *v1.ProjectDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *projectController) ProjectGet(ctx context.Context, req *v1.ProjectGetReq) (res *v1.ProjectGetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *projectController) ProjectList(ctx context.Context, req *v1.ProjectListReq) (res *v1.ProjectListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *projectController) ProjectUpdate(ctx context.Context, req *v1.ProjectUpdateReq) (res *v1.ProjectUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
