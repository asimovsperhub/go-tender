package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	Publish = cPublish{}
)

type cPublish struct {
}

func (*cPublish) PublishKnowledge(ctx context.Context, req *desk.KnowledgeReq) (res *desk.KnowledgeRes, err error) {
	res, err = service.Publish().KnowledgePublish(ctx, req)
	return
}

func (*cPublish) PublishKnowledgeListGet(ctx context.Context, req *desk.KnowledgeGetListReq) (res *desk.KnowledgeGetListRes, err error) {
	res, err = service.Publish().KnowledgePublishListGet(ctx, req)
	return
}
func (*cPublish) PublishKnowledgeGet(ctx context.Context, req *desk.KnowledgeGetReq) (res *desk.KnowledgeGetRes, err error) {
	res, err = service.Publish().KnowledgePublishGet(ctx, req)
	return
}

func (*cPublish) PublishKnowledgeEdit(ctx context.Context, req *desk.KnowledgeEditReq) (res *desk.KnowledgeEditRes, err error) {
	res, err = service.Publish().KnowledgePublishEdit(ctx, req)
	return
}

// KnowledgeDelReq
func (*cPublish) PublishKnowledgeDel(ctx context.Context, req *desk.KnowledgeDelReq) (res *desk.KnowledgeDelRes, err error) {
	res, err = service.Publish().KnowledgePublishDel(ctx, req)
	return
}

//KnowledgeBuyReq
func (*cPublish) KnowledgeBuy(ctx context.Context, req *desk.KnowledgeBuyReq) (res *desk.KnowledgeBuyRes, err error) {
	res, err = service.Publish().KnowledgeBuy(ctx, req)
	return
}
