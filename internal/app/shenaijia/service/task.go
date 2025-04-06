// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

type (
	ITask interface {
		SearchType(ctx context.Context, req *v1.TempletTypeSearchReq) (res *v1.TempletTypeSearchRes, err error)
		AddType(ctx context.Context, req *v1.TempletTypeAddReq) (res *v1.TempletTypeAddRes, err error)
		UpdateType(ctx context.Context, req *v1.TempletTypeUpdateReq) (res *v1.TempletTypeUpdateRes, err error)
		DeleteType(ctx context.Context, req *v1.TempletTypeDeleteReq) (res *v1.TempletTypeDeleteRes, err error)
		HasTemplet(ctx context.Context, templetType uint) (res bool, err error)
		SearchTemplet(ctx context.Context, req *v1.TaskTempletSearchReq) (res *v1.TaskTempletSearchRes, err error)
		AddTemplet(ctx context.Context, req *v1.TaskTempletAddReq) (res *v1.TaskTempletAddRes, err error)
		UpdateTemplet(ctx context.Context, req *v1.TaskTempletUpdateReq) (res *v1.TaskTempletUpdateRes, err error)
		DeleteTemplet(ctx context.Context, req *v1.TaskTempletDeleteReq) (res *v1.TaskTempletDeleteRes, err error)
		TempletSetFlow(ctx context.Context, req *v1.TaskTempletSetFlowReq) (res *v1.TaskTempletSetFlowRes, err error)
		TempletGetFlow(ctx context.Context, req *v1.TaskTempletGetFlowReq) (res *v1.TaskTempletGetFlowRes, err error)
		AddStep(ctx context.Context, req *v1.TaskAddStepReq) (res *v1.TaskAddStepRes, err error)
		UpdateStep(ctx context.Context, req *v1.TaskUpdateStepReq) (res *v1.TaskAddStepRes, err error)
		DeleteStep(ctx context.Context, req *v1.TaskDeleteStepReq) (res *v1.TaskDeleteStepRes, err error)
	}
)

var (
	localTask ITask
)

func Task() ITask {
	if localTask == nil {
		panic("implement not found for interface ITask, forgot register?")
	}
	return localTask
}

func RegisterTask(i ITask) {
	localTask = i
}
