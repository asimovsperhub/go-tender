package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	ICollect interface {
		CollectAdd(ctx context.Context, req *desk.CollectAddReq) (res *desk.CollectAddRes, err error)
		CollectDel(ctx context.Context, req *desk.CollectDelReq) (res *desk.CollectDelRes, err error)
		CollectGetList(ctx context.Context, req *desk.CollectGetListReq) (res *desk.CollectGetListRes, err error)
		CollectGet(ctx context.Context, req *desk.CollectGetReq) (res *desk.CollectGetRes, err error)
		CollectCountGet(ctx context.Context, req *desk.CollectCountGetReq) (res *desk.CollectCountGetRes, err error)

		SubscribeAdd(ctx context.Context, req *desk.SubscribeAddReq) (res *desk.SubscribeAddRes, err error)
		SubscribeDel(ctx context.Context, req *desk.SubscribeDelReq) (res *desk.SubscribeDelRes, err error)
		SubscribeEdit(ctx context.Context, req *desk.SubscribeEditReq) (res *desk.SubscribeEditRes, err error)
		SubscribeList(ctx context.Context, req *desk.SubscribeListReq) (res *desk.SubscribeListRes, err error)
		SubscribeGet(ctx context.Context, req *desk.SubscribeGetReq) (res *desk.SubscribeGetRes, err error)
		SubscribeCan(ctx context.Context, req *desk.SubscribeCanReq) (res *desk.SubscribeCanRes, err error)
	}
)

var (
	localCollect ICollect
)

func Collect() ICollect {
	if localCollect == nil {
		panic("implement not found for interface ICollect, forgot register?")
	}
	return localCollect
}

// 注册发布接口
func RegisterCollect(i ICollect) {
	localCollect = i
}
