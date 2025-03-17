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

func (c *taskController) StageAddStep(ctx context.Context, req *v1.StageAddStepReq) (res *v1.StageAddStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) StageDelStep(ctx context.Context, req *v1.StageDelStepReq) (res *v1.StageDelStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) StageUpdateStep(ctx context.Context, req *v1.StageUpdateStepReq) (res *v1.StageUpdateStepRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskAdd(ctx context.Context, req *v1.TaskAddReq) (res *v1.TaskAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskAddStage(ctx context.Context, req *v1.TaskAddStageReq) (res *v1.TaskAddStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskAddType(ctx context.Context, req *v1.TaskAddTypeReq) (res *v1.TaskAddTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskAddAction(ctx context.Context, req *v1.TaskAddWorkReq) (res *v1.TaskAddWorkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskDelete(ctx context.Context, req *v1.TaskDeleteReq) (res *v1.TaskDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskDeleteStage(ctx context.Context, req *v1.TaskDeleteStageReq) (res *v1.TaskDeleteStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskDeleteType(ctx context.Context, req *v1.TaskDeleteTypeReq) (res *v1.TaskDeleteTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskGet(ctx context.Context, req *v1.TaskGetReq) (res *v1.TaskGetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskList(ctx context.Context, req *v1.TaskListReq) (res *v1.TaskListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskListStage(ctx context.Context, req *v1.TaskListStageReq) (res *v1.TaskListStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskListType(ctx context.Context, req *v1.TaskListTypeReq) (res *v1.TaskListTypeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskListAction(ctx context.Context, req *v1.TaskListWorkReq) (res *v1.TaskListWorkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskUpdate(ctx context.Context, req *v1.TaskUpdateReq) (res *v1.TaskUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskUpdateStage(ctx context.Context, req *v1.TaskUpdateStageReq) (res *v1.TaskUpdateStageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *taskController) TaskUpdateAction(ctx context.Context, req *v1.TaskUpdateWorkReq) (res *v1.TaskUpdateWorkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
