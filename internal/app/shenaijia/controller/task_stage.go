package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	TaskStage = taskStageController{}
)

type taskStageController struct {
	commonController.BaseController
}

func (c *taskStageController) List(ctx context.Context, req *v1.TaskListStageReq) (res *v1.TaskListStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskStageController) Add(ctx context.Context, req *v1.TaskAddStageReq) (res *v1.TaskAddStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskStageController) Update(ctx context.Context, req *v1.TaskUpdateStageReq) (res *v1.TaskUpdateStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskStageController) Delete(ctx context.Context, req *v1.TaskDeleteStageReq) (res *v1.TaskDeleteStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
