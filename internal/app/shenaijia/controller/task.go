package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Task = taskController{}
)

type taskController struct {
	commonController.BaseController
}

func (c *taskController) Search(ctx context.Context, req v1.TaskSearchReq) (res *v1.TaskSearchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) Add(ctx context.Context, req v1.TaskAddReq) (res *v1.TaskAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) Update(ctx context.Context, req v1.TaskUpdateReq) (res *v1.TaskUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) Delete(ctx context.Context, req v1.TaskDeleteReq) (res *v1.TaskDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) ListStage(ctx context.Context, req v1.TaskListStageReq) (res *v1.TaskListStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) AddStage(ctx context.Context, req v1.TaskAddStageReq) (res *v1.TaskAddStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) UpdateStage(ctx context.Context, req v1.TaskAddStageReq) (res *v1.TaskAddStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) DeleteStage(ctx context.Context, req v1.TaskDeleteStageReq) (res *v1.TaskDeleteStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) ListStep(ctx context.Context, req v1.TaskListStepReq) (res *v1.TaskListStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) AddStep(ctx context.Context, req *v1.TaskAddStepReq) (res *v1.TaskAddStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) UpdateStep(ctx context.Context, req *v1.TaskAddStepReq) (res *v1.TaskAddStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) DeleteStep(ctx context.Context, req *v1.TaskDeleteStepReq) (res *v1.TaskDeleteStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) CompleteStep(ctx context.Context, req *v1.TaskCompleteStepReq) (res *v1.TaskCompleteStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
