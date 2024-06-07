package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	Collect = cCollect{}
)

type cCollect struct {
}

func (*cCollect) CollectAdd(ctx context.Context, req *desk.CollectAddReq) (res *desk.CollectAddRes, err error) {
	res, err = service.Collect().CollectAdd(ctx, req)
	return
}

func (*cCollect) CollectDel(ctx context.Context, req *desk.CollectDelReq) (res *desk.CollectDelRes, err error) {
	res, err = service.Collect().CollectDel(ctx, req)
	return
}

//CollectGetListReq
func (*cCollect) CollectGetList(ctx context.Context, req *desk.CollectGetListReq) (res *desk.CollectGetListRes, err error) {
	res, err = service.Collect().CollectGetList(ctx, req)
	return
}

//CollectGetReq
func (*cCollect) CollectGet(ctx context.Context, req *desk.CollectGetReq) (res *desk.CollectGetRes, err error) {
	res, err = service.Collect().CollectGet(ctx, req)
	return
}

//SubscribeAddReq

func (*cCollect) SubscribeAdd(ctx context.Context, req *desk.SubscribeAddReq) (res *desk.SubscribeAddRes, err error) {
	res, err = service.Collect().SubscribeAdd(ctx, req)
	return
}

func (*cCollect) SubscribeDel(ctx context.Context, req *desk.SubscribeDelReq) (res *desk.SubscribeDelRes, err error) {
	res, err = service.Collect().SubscribeDel(ctx, req)
	return
}

// SubscribeEdit
func (*cCollect) SubscribeEdit(ctx context.Context, req *desk.SubscribeEditReq) (res *desk.SubscribeEditRes, err error) {
	res, err = service.Collect().SubscribeEdit(ctx, req)
	return
}

//SubscribeList
func (*cCollect) SubscribeList(ctx context.Context, req *desk.SubscribeListReq) (res *desk.SubscribeListRes, err error) {
	res, err = service.Collect().SubscribeList(ctx, req)
	return
}

//SubscribeGetReq
func (*cCollect) SubscribeGet(ctx context.Context, req *desk.SubscribeGetReq) (res *desk.SubscribeGetRes, err error) {
	res, err = service.Collect().SubscribeGet(ctx, req)
	return
}

//SubscribeCanReq
func (*cCollect) SubscribeCan(ctx context.Context, req *desk.SubscribeCanReq) (res *desk.SubscribeCanRes, err error) {
	res, err = service.Collect().SubscribeCan(ctx, req)
	return
}
