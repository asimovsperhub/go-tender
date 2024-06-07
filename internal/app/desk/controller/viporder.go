package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	VipOrder = cVipOrderController{}
)

type cVipOrderController struct {
	BaseController
}

func (*cVipOrderController) VipOrderCreate(ctx context.Context, req *desk.CreateVipOrderReq) (res *desk.CreateVipOrderRes, err error) {
	res, err = service.VipOrder().CreateVipOrder(ctx, req)
	return res, err
}

func (*cVipOrderController) VipOrderPayment(ctx context.Context, req *desk.QueryVipOrderPaymentReq) (res *desk.QueryVipOrderPaymentRes, err error) {
	res, err = service.VipOrder().QueryPaymentInfo(ctx, req)
	return res, err
}

func (*cVipOrderController) VipOrderQuery(ctx context.Context, req *desk.QueryVipOrderInfoReq) (res *desk.QueryVipOrderInfoRes, err error) {
	res, err = service.VipOrder().QueryVipOrder(ctx, req)
	return res, err
}

func (*cVipOrderController) VipOrderPaymentForAlipay(ctx context.Context, req *desk.QueryVipOrderPaymentAlipayReq) (res *desk.QueryVipOrderPaymentAlipayRes, err error) {
	res, err = service.VipOrder().QueryPaymentInfoForAlipay(ctx, req)
	return res, err
}
