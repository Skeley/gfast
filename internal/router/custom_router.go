package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/controller"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func (router *Router) BindShenAiJiaController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1/shenaijia", func(group *ghttp.RouterGroup) {
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)

		group.Bind(
			controller.Login,
		)
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx)
		group.Bind(
			controller.Project,
			controller.Community,
			controller.TaskTemplet,
			controller.SysTempletType,
			controller.File,
			controller.SysUser,
		)
	})
}
