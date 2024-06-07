package cmd

import (
	"context"
	"tender/internal/app/common/service"
	"tender/internal/consts"
	"tender/internal/packed/websocket"
	"tender/internal/router"

	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
			g.Log().Info(ctx, "TENDER version:", consts.Version)
			service.InitializeS3(ctx)
			service.Initialize(ctx)
			//启动服务
			websocket.StartWebSocket(ctx)
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
				//注册路由
				group.ALL("/ws", websocket.WsPage)
			})
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact:     &goai.Contact{
			// Name: consts.OpenAPIContactName,
			//URL:  consts.OpenAPIContactUrl,
		},
	}
}
