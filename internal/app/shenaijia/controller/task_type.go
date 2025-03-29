package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	TaskType = taskTypeController{}
)

type taskTypeController struct {
	commonController.BaseController
}

func (c *taskTypeController) List(ctx context.Context, req *v1.TaskListTypeReq) (res *v1.TaskListTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskTypeController) Add(ctx context.Context, req *v1.TaskAddTypeReq) (res *v1.TaskAddTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskTypeController) Delete(ctx context.Context, req *v1.TaskDeleteTypeReq) (res *v1.TaskDeleteTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
