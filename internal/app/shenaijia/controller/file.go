package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"

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
	if req.File.Size > 20*1024*1024 { // max file size: 20M
		return nil, gerror.New("文件大小超限，必须小于20M")
	}
	req.File.Filename = fmt.Sprintf("%d-%s", sysService.Context().GetUserId(ctx), req.File.Filename)
	return service.File().Upload(ctx, req)
}
