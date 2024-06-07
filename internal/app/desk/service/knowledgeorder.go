package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	IKnowledgeOrder interface {
		CreateOrder(ctx context.Context, req *desk.CreateKnowledgeOrderReq) (res *desk.CreateKnowledgeOrderRes, err error)
		QueryPaymentInfo(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentReq) (res *desk.QueryKnowledgeOrderPaymentRes, err error)
		QueryOrder(ctx context.Context, req *desk.QueryKnowledgeOrderInfoReq) (res *desk.QueryKnowledgeOrderInfoRes, err error)
		QueryPaymentInfoForAlipay(ctx context.Context, req *desk.QueryKnowledgeOrderPaymentAlipayReq) (res *desk.QueryKnowledgeOrderPaymentAlipayRes, err error)
	}
)

var (
	localKnowledgeOrder IKnowledgeOrder
)

func KnowledgeOrder() IKnowledgeOrder {
	if localKnowledgeOrder == nil {
		panic("implement not found for interface IKnowledgeOrder, forgot register?")
	}
	return localKnowledgeOrder
}

// 注册发布接口
func RegisterKnowledgeOrder(i IKnowledgeOrder) {
	localKnowledgeOrder = i
}
