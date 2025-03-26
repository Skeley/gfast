package controller

import (
	"context"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Community = communityController{}
)

type communityController struct {
	commonController.BaseController
}

func (c *communityController) ListAll(ctx context.Context, req *v1.CommunityListAllReq) (res *v1.CommunityRes, err error) {
	return &v1.CommunityRes{}, nil
}

func (c *communityController) Search(ctx context.Context, req *v1.CommunitySearchReq) (res *v1.CommunityRes, err error) {
	return &v1.CommunityRes{}, nil
}

func (c *communityController) List(ctx context.Context, req *v1.CommunityListReq) (res *v1.CommunityRes, err error) {
	return &v1.CommunityRes{}, nil
}

func (c *communityController) Add(ctx context.Context, req *v1.CommunityAddReq) (res *v1.CommunityAddRes, err error) {
	return &v1.CommunityAddRes{}, nil
}

func (c *communityController) Update(ctx context.Context, req *v1.CommunityUpdateReq) (res *v1.CommunityUpdateRes, err error) {
	return &v1.CommunityUpdateRes{}, nil
}
