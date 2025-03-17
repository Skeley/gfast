package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "github.com/tiger1103/gfast/v3/api/shenaijia/v1"
	commonController "github.com/tiger1103/gfast/v3/internal/app/common/controller"
)

var (
	Login = loginController{}
)

type loginController struct {
	commonController.BaseController
}

func (c *loginController) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
