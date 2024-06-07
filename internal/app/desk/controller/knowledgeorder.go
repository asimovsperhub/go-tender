package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	KnowledgeOrder = cKnowledgeOrderController{}
)

type cKnowledgeOrderController struct {
	BaseController
}

func (*cKnowledgeOrderController) OrderCreate(ctx context.Context, req *desk.CreateKnowledgeOrderReq) (res *desk.CreateKnowledgeOrderRes, err error) {
	res, err = service.KnowledgeOrder().CreateOrder(ctx, req)
	return res, err
}

func (*cKnowledgeOrderController) OrderPayment(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentReq) (res *desk.QueryKnowledgeOrderPaymentRes, err error) {
	res, err = service.KnowledgeOrder().QueryPaymentInfo(ctx, req)
	return res, err
}

func (*cKnowledgeOrderController) OrderQuery(ctx context.Context, req *desk.QueryKnowledgeOrderInfoReq) (res *desk.QueryKnowledgeOrderInfoRes, err error) {
	res, err = service.KnowledgeOrder().QueryOrder(ctx, req)
	return res, err
}

func (*cKnowledgeOrderController) OrderPaymentForAlipay(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentAlipayReq) (res *desk.QueryKnowledgeOrderPaymentAlipayRes, err error) {
	res, err = service.KnowledgeOrder().QueryPaymentInfoForAlipay(ctx, req)
	return res, err
}
