package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	RechargeOrder = cRechargeOrderController{}
)

type cRechargeOrderController struct {
	BaseController
}

func (*cRechargeOrderController) RechargeOrderCreate(ctx context.Context, req *desk.CreateRechargeOrderReq) (res *desk.CreateRechargeOrderRes, err error) {
	res, err = service.RechargeOrder().CreateRechargeOrder(ctx, req)
	return res, err
}

func (*cRechargeOrderController) RechargeOrderPayment(ctx context.Context, req *desk.QueryRechargeOrderPaymentReq) (res *desk.QueryRechargeOrderPaymentRes, err error) {
	res, err = service.RechargeOrder().QueryPaymentInfo(ctx, req)
	return res, err
}

func (*cRechargeOrderController) RechargeOrderQuery(ctx context.Context, req *desk.QueryRechargeOrderInfoReq) (res *desk.QueryRechargeOrderInfoRes, err error) {
	res, err = service.RechargeOrder().QueryRechargeOrder(ctx, req)
	return res, err
}

func (*cRechargeOrderController) RechargeOrderPaymentForAlipay(ctx context.Context, req *desk.QueryRechargeOrderPaymentAlipayReq) (res *desk.QueryRechargeOrderPaymentAlipayRes, err error) {
	res, err = service.RechargeOrder().QueryPaymentInfoForAlipay(ctx, req)
	return res, err
}
