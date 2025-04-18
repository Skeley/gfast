// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	sysModel "github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type (
	IProject interface {
		List(ctx context.Context, userCtx *sysModel.Context, req *api.ProjectListReq) (res *api.ProjectListRes, err error)
		Add(ctx context.Context, manager uint64, creator uint64, req *api.ProjectAddReq) (res *api.ProjectAddRes, err error)
		Update(ctx context.Context, req *api.ProjectUpdateReq) (res *api.ProjectUpdateRes, err error)
		SysList(ctx context.Context, req *api.SysProjectSearchReq) (res *api.SysProjectSearchRes, err error)
		SysAdd(ctx context.Context, req *api.SysProjectAddReq) (res *api.SysProjectAddRes, err error)
		SysEdit(ctx context.Context, req *api.SysProjectEditReq) (res *api.SysProjectEditRes, err error)
		Delete(ctx context.Context, req *api.ProjectDeleteReq) (res *api.ProjectDeleteRes, err error)
		Audit(ctx context.Context, req *api.SysProjectAuditReq) (res *api.SysProjectAuditRes, err error)
	}
)

var (
	localProject IProject
)

func Project() IProject {
	if localProject == nil {
		panic("implement not found for interface IProject, forgot register?")
	}
	return localProject
}

func RegisterProject(i IProject) {
	localProject = i
}
