package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	IVipOrder interface {
		CreateVipOrder(ctx context.Context, req *desk.CreateVipOrderReq) (res *desk.CreateVipOrderRes, err error)
		QueryPaymentInfo(ctx context.Context, req *desk.QueryVipOrderPaymentReq) (res *desk.QueryVipOrderPaymentRes, err error)
		QueryVipOrder(ctx context.Context, req *desk.QueryVipOrderInfoReq) (res *desk.QueryVipOrderInfoRes, err error)
		QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QueryVipOrderPaymentAlipayReq) (res *desk.QueryVipOrderPaymentAlipayRes, err error)
	}
)

var (
	localVipOrder IVipOrder
)

func VipOrder() IVipOrder {
	if localVipOrder == nil {
		panic("implement not found for interface IVipOrder, forgot register?")
	}
	return localVipOrder
}

// 注册发布接口
func RegisterVipOrder(i IVipOrder) {
	localVipOrder = i
}
