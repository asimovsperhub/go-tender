package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "tender/internal/app/common/router"
	commonService "tender/internal/app/common/service"
	deskRouter "tender/internal/app/desk/router"
	systemRouter "tender/internal/app/system/router"
	"tender/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//跨域处理，安全起见正式环境请注释该行
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定后台路由
		systemRouter.R.BindController(ctx, group)
		//绑定前台接口
		deskRouter.R.BindControllerDesk(ctx, group)
		// 绑定公共路由
		commonRouter.R.BindController(ctx, group)
		//自动绑定定义的模块
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
