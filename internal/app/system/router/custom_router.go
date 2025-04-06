package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/shenaijia/controller"
)

func (router *Router) BindShenAiJiaController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/shenaijia", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.SysProject,
			controller.Community,
			controller.Task,
			controller.File,
			controller.TempletType,
			controller.TaskTemplet,
			controller.User,
		)
	})
}
