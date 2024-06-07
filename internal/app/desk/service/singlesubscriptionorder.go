package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	ISingleSubscriptionOrder interface {
		CreateSingleSubscriptionOrder(ctx context.Context, req *desk.CreateSingleSubscriptionOrderReq) (res *desk.CreateSingleSubscriptionOrderRes, err error)
		QueryPaymentInfo(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentReq) (res *desk.QuerySingleSubscriptionOrderPaymentRes, err error)
		QuerySingleSubscriptionOrder(ctx context.Context, req *desk.QuerySingleSubscriptionOrderInfoReq) (res *desk.QuerySingleSubscriptionOrderInfoRes, err error)
		QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QuerySingleSubscriptionOrderPaymentAlipayReq) (res *desk.QuerySingleSubscriptionOrderPaymentAlipayRes, err error)
	}
)

var (
	localSingleSubscriptionOrder ISingleSubscriptionOrder
)

func SingleSubscriptionOrder() ISingleSubscriptionOrder {
	if localSingleSubscriptionOrder == nil {
		panic("implement not found for interface ISingleSubscriptionOrder, forgot register?")
	}
	return localSingleSubscriptionOrder
}

// 注册发布接口
func RegisterSingleSubscriptionOrder(i ISingleSubscriptionOrder) {
	localSingleSubscriptionOrder = i
}
