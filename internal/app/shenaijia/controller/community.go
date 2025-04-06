package controller

import (
	"context"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
)

var (
	Community = communityController{}
)

type communityController struct {
	commonController.BaseController
}

func (c *communityController) Search(ctx context.Context, req *v1.CommunitySearchReq) (res *v1.CommunityRes, err error) {
	return service.Community().Search(ctx, req)
}
