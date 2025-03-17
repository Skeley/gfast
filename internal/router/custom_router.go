package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/controller"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

func (router *Router) BindShenAiJiaController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/shenaijia", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			controller.Login,
		)
		service.GfToken().Middleware(group)
		group.Bind(
			controller.Project,
			controller.Community,
			controller.Task,
			controller.File,
		)
	})
}
