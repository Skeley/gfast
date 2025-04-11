package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	SysProject = sysProjectController{}
)

type sysProjectController struct {
	commonController.BaseController
}

func (c *sysProjectController) Add(ctx context.Context, req *v1.SysProjectAddReq) (res *v1.SysProjectAddRes, err error) {
	return service.Project().SysAdd(ctx, req)
}

func (c *sysProjectController) Delete(ctx context.Context, req *v1.SysProjectDelReq) (res *v1.SysProjectDelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *sysProjectController) List(ctx context.Context, req *v1.SysProjectSearchReq) (res *v1.SysProjectSearchRes, err error) {
	res, err = service.Project().SysList(ctx, req)
	if err != nil {
		return res, err
	}
	var communityList []*v1.Community
	for _, v := range res.List {
		communityList = append(communityList, &v1.Community{
			Id:  v.CommunityId,
			Pid: v.CommunityPid,
		})
	}
	communityList, err = service.Community().FillParentName(ctx, communityList)
	if err != nil {
		return res, err
	}
	for i, _ := range res.List {
		res.List[i].CommunityParentName = communityList[i].ParentName
	}
	return res, nil
}

func (c *sysProjectController) Update(ctx context.Context, req *v1.SysProjectEditReq) (res *v1.SysProjectEditRes, err error) {
	return service.Project().SysEdit(ctx, req)
}

func (c *sysProjectController) Audit(ctx context.Context, req *v1.SysProjectAuditReq) (res *v1.SysProjectAuditRes, err error) {
	return service.Project().Audit(ctx, req)

}
