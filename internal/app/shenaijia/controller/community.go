package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Community = communityController{}
)

type communityController struct {
	commonController.BaseController
}

func (c *communityController) CommunityListFirstLevel(ctx context.Context, req *v1.CommunityListFirstLevelReq) (res *v1.CommunityListFirstLevelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (c *communityController) CommunityListTwoLevel(ctx context.Context, req *v1.CommunityListTwoLevelReq) (res *v1.CommunityListTwoLevelRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
