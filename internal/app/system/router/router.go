package router

import (
	"context"
	DeskController "tender/internal/app/desk/controller"
	"tender/internal/app/system/controller"
	"tender/internal/app/system/service"
	"tender/library/libRouter"

	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

// 后台
func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Login.Login,
			controller.Login.Qrcode,
			controller.Login.CheckScanQrcode,
			DeskController.Bbs.BbsGet,
		)
		// todo 给前端测试不用token  后边要改回来
		group.GET("/knowledge/download", controller.Download.DownloadKnowledgeFile)
		group.GET("/tender/download", controller.Download.DownloadTenderFile)
		//DownloadSingleTenderFile
		group.GET("/tender/single/download", controller.Download.DownloadSingleTenderFile)
		group.GET("/enterprise/download", controller.Download.DownloadEnterpriseFile)
		group.GET("/consultation/download", controller.Download.DownloadConsultationFile)
		// 反馈页批量下载附件
		group.GET("/feedback/download", controller.Download.DownloadFeedbackFile)

		// 会员导出 MemberExport
		group.GET("/memberExport", controller.Download.MemberExport)
		// 企业导出 EnterpriseExport
		group.GET("/enterpriseExport", controller.Download.EnterpriseExport)
		// 	未认证的企业导出 EnterpriseUncertifiedExport
		group.GET("/enterpriseUncertifiedExport", controller.Download.EnterpriseUncertifiedExport)
		// 反馈页导出 FeedbackExport
		group.GET("/feedbackExport", controller.Download.FeedbackExport)
		// 财务汇总导出 FinanceExport
		group.GET("/financeExport", controller.Download.FinanceExport)
		// 下载模版
		group.GET("/template", controller.Download.DownloadSystemTp)
		//登录验证拦截
		service.GfToken().Middleware(group)
		group.Bind(
			controller.Msg,
		)
		//context拦截器 + 权限拦截器 用户对应的菜单
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		//group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)

		// excel导入知识
		group.POST("/importFileExcel", controller.Upload.SystemImportFileExcel)
		//
		group.POST("/importFileWord", controller.Upload.SystemImportFileWord)
		group.POST("/importFilePdf", controller.Upload.SystemImportFilePdf)
		group.POST("/importFileVideo", controller.Upload.SystemImportFileVideo)
		//SystemImportFileOther
		group.POST("/importFileOther", controller.Upload.SystemImportFileOther)
		group.Bind(
			controller.User,
			controller.Menu,
			controller.Role,
			controller.MemberManger,
			controller.Personal,
			controller.EnterpriseManger,
			controller.Data,
			controller.KnowledgeManger,
			controller.SysIndexManger,
			controller.SysFinanceManger,
			controller.Login.LoginOut,
			DeskController.Publish.PublishKnowledgeEdit,
			controller.ForumManger,
		)
		//group.GET("/knowledge/download", controller.Download.DownloadKnowledgeFile)
		//group.GET("/tender/download", controller.Download.DownloadTenderFile)
		//group.GET("/enterprise/download", controller.Download.DownloadEnterpriseFile)
		//group.GET("/consultation/download", controller.Download.DownloadConsultationFile)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
