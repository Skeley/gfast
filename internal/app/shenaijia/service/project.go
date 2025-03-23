// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	api "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
)

type (
	IProject interface {
		SysList(ctx context.Context, req *api.SysProjectListReq) (res *api.SysProjectListRes, err error)
		SysAdd(ctx context.Context, req *api.SysProjectAddReq) (res *api.SysProjectAddRes, err error)
		SysEdit(ctx context.Context, req *api.SysProjectEditReq) (res *api.SysProjectEditRes, err error)
		Delete(ctx context.Context, req *api.ProjectDeleteReq) (res *api.ProjectDeleteRes, err error)
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
