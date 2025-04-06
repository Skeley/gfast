package controller

import (
	"context"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
)

var (
	TaskTemplet = taskTempletController{}
)

type taskTempletController struct {
	commonController.BaseController
}

func (c *taskTempletController) Search(ctx context.Context, req *v1.TaskTempletSearchReq) (res *v1.TaskTempletSearchRes, err error) {
	return service.Task().SearchTemplet(ctx, req)
}

func (c *taskTempletController) Add(ctx context.Context, req *v1.TaskTempletAddReq) (res *v1.TaskTempletAddRes, err error) {
	return service.Task().AddTemplet(ctx, req)
}

func (c *taskTempletController) Update(ctx context.Context, req *v1.TaskTempletUpdateReq) (res *v1.TaskTempletUpdateRes, err error) {
	return service.Task().UpdateTemplet(ctx, req)
}

func (c *taskTempletController) Delete(ctx context.Context, req *v1.TaskTempletDeleteReq) (res *v1.TaskTempletDeleteRes, err error) {
	return service.Task().DeleteTemplet(ctx, req)
}

func (c *taskTempletController) SetFlow(ctx context.Context, req *v1.TaskTempletSetFlowReq) (res *v1.TaskTempletSetFlowRes, err error) {
	return service.Task().TempletSetFlow(ctx, req)
}

func (c *taskTempletController) GetFlow(ctx context.Context, req *v1.TaskTempletGetFlowReq) (res *v1.TaskTempletGetFlowRes, err error) {
	return service.Task().TempletGetFlow(ctx, req)
}
