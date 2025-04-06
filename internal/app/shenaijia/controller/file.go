package controller

import (
	"context"
	"fmt"

	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/service"
	sysService "github.com/tiger1103/gfast/v3/internal/app/system/service"

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
	req.FileName = fmt.Sprintf("%d-%s", sysService.Context().GetUserId(ctx), req.FileName)
	return service.File().Upload(ctx, req)
}
