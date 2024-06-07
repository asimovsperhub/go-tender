package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"tender/internal/app/desk/controller"
	"tender/internal/app/desk/service"
	Syscontroller "tender/internal/app/system/controller"
	"tender/library/libRouter"
)

var R = new(Router)

type Router struct{}

// 前台
func (router *Router) BindControllerDesk(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/desk", func(group *ghttp.RouterGroup) {
		//context拦截器 + 用户信息
		group.Middleware(service.Middleware().Ctx)
		group.Bind(
			//注册
			controller.Register,
			controller.Login.Login,
			controller.Login.MsgLogin,
			controller.Login.UserPassEdit,
			controller.TnderData,
			controller.Index,
			Syscontroller.Enterprise.EnterpriseGet,
			controller.Publish.PublishKnowledgeGet,
			controller.Publish.PublishKnowledgeListGet,
			controller.Publish.KnowledgeBuy,
			Syscontroller.MemberManger.FindFee,
			Syscontroller.MemberManger.FindIn,
			Syscontroller.MemberManger.FindSu,
			controller.Bbs.BbsGet,
			controller.Bbs.BbsGetList,
			controller.Bbs.BbsLike,
			controller.Bbs.BbsCommentLikeReq,
			controller.Bbs.BbsBrowse,
			controller.Bbs.BbsGetListAll,
		)

		group.POST("/pay/wxnotify", controller.Notify.WxNotify)
		group.POST("/pay/alinotify", controller.Notify.AlipayNotify)
		group.POST("/pay/knowledgeorder/wxnotify", controller.KnowledgeOrderNotify.WxNotify)
		group.POST("/pay/knowledgeorder/alinotify", controller.KnowledgeOrderNotify.AlipayNotify)
		group.POST("/pay/rechargeorder/wxnotify", controller.RechargeOrderNotify.WxNotify)
		group.POST("/pay/rechargeorder/alinotify", controller.RechargeOrderNotify.AlipayNotify)
		group.POST("/pay/singlesubscription/wxnotify", controller.SingleSubscriptionOrderNotify.WxNotify)
		group.POST("/pay/singlesubscription/alinotify", controller.SingleSubscriptionOrderNotify.AlipayNotify)

		//登录验证拦截
		service.GfToken().Middleware(group)
		group.Bind(
			// 个人中心
			Syscontroller.Enterprise.EnterpriseAdd,
			Syscontroller.Enterprise.EnterpriseEdit,
			controller.Publish.PublishKnowledgeEdit,
			controller.Publish.PublishKnowledgeDel,
			controller.Publish.PublishKnowledge,
			controller.Collect,
			controller.VipOrder,
			controller.KnowledgeOrder,
			controller.Login.LoginOut,
			controller.UserInfo,
			controller.RechargeOrder,
			controller.SingleSubscriptionOrder,
			controller.Bbs.BbsEdit,
			controller.Bbs.BbsDel,
			controller.Bbs.BbsPublish,
			controller.Bbs.BbsComment,
			controller.Bbs.BbsCommentDel,
			controller.Bbs.BbsReply,
			controller.Bbs.BbsGetLike,
			controller.Bbs.FeedbackPublish,
		)
		//自动绑定定义的控制器
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
