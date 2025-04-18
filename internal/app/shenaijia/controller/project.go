package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	sysService "github.com/tiger1103/gfast/v3/internal/app/system/service"

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
	manager, e := service.Community().BoundUser(ctx, req.CommunityId)
	if e != nil {
		return nil, e
	}
	user := sysService.Context().GetLoginUser(ctx)
	creator := user.Id
	return service.Project().Add(ctx, manager, creator, req)
}

func (c *projectController) ProjectDelete(ctx context.Context, req *v1.ProjectDeleteReq) (res *v1.ProjectDeleteRes, err error) {
	return service.Project().Delete(ctx, req)
}

func (c *projectController) ProjectList(ctx context.Context, req *v1.ProjectListReq) (res *v1.ProjectListRes, err error) {
	userCtx := sysService.Context().Get(ctx)
	return service.Project().List(ctx, userCtx, req)
}

func (c *projectController) ProjectUpdate(ctx context.Context, req *v1.ProjectUpdateReq) (res *v1.ProjectUpdateRes, err error) {
	return service.Project().Update(ctx, req)
}
