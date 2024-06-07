package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	SingleSubscriptionOrder = cSingleSubscriptionOrderController{}
)

type cSingleSubscriptionOrderController struct {
	BaseController
}

func (*cSingleSubscriptionOrderController) SingleSubscriptionOrderCreate(ctx context.Context, req *desk.CreateSingleSubscriptionOrderReq) (res *desk.CreateSingleSubscriptionOrderRes, err error) {
	res, err = service.SingleSubscriptionOrder().CreateSingleSubscriptionOrder(ctx, req)
	return res, err
}

func (*cSingleSubscriptionOrderController) SingleSubscriptionOrderPayment(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentReq) (res *desk.QuerySingleSubscriptionOrderPaymentRes, err error) {
	res, err = service.SingleSubscriptionOrder().QueryPaymentInfo(ctx, req)
	return res, err
}

func (*cSingleSubscriptionOrderController) SingleSubscriptionOrderQuery(ctx context.Context, req *desk.QuerySingleSubscriptionOrderInfoReq) (res *desk.QuerySingleSubscriptionOrderInfoRes, err error) {
	res, err = service.SingleSubscriptionOrder().QuerySingleSubscriptionOrder(ctx, req)
	return res, err
}

func (*cSingleSubscriptionOrderController) SingleSubscriptionOrderPaymentForAlipay(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentAlipayReq) (res *desk.QuerySingleSubscriptionOrderPaymentAlipayRes, err error) {
	res, err = service.SingleSubscriptionOrder().QueryPaymentInfoForAlipay(ctx, req)
	return res, err
}
