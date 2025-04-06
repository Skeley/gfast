package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"

	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	File = fileController{}
)

type fileController struct {
	commonController.BaseController
}

func (c *fileController) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	return service.File().Upload(ctx, req)
}
