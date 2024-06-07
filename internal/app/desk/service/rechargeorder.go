package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	IRechargeOrder interface {
		CreateRechargeOrder(ctx context.Context, req *desk.CreateRechargeOrderReq) (res *desk.CreateRechargeOrderRes, err error)
		QueryPaymentInfo(ctx context.Context, req *desk.QueryRechargeOrderPaymentReq) (res *desk.QueryRechargeOrderPaymentRes, err error)
		QueryRechargeOrder(ctx context.Context, req *desk.QueryRechargeOrderInfoReq) (res *desk.QueryRechargeOrderInfoRes, err error)
		QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QueryRechargeOrderPaymentAlipayReq) (res *desk.QueryRechargeOrderPaymentAlipayRes, err error)
	}
)

var (
	localRechargeOrder IRechargeOrder
)

func RechargeOrder() IRechargeOrder {
	if localRechargeOrder == nil {
		panic("implement not found for interface IRechargeOrder, forgot register?")
	}
	return localRechargeOrder
}

// 注册发布接口
func RegisterRechargeOrder(i IRechargeOrder) {
	localRechargeOrder = i
}
