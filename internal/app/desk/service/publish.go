package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	IPublish interface {
		KnowledgePublish(ctx context.Context, req *desk.KnowledgeReq) (res *desk.KnowledgeRes, err error)
		KnowledgePublishListGet(ctx context.Context, req *desk.KnowledgeGetListReq) (res *desk.KnowledgeGetListRes, err error)
		KnowledgePublishGet(ctx context.Context, req *desk.KnowledgeGetReq) (res *desk.KnowledgeGetRes, err error)
		KnowledgePublishEdit(ctx context.Context, req *desk.KnowledgeEditReq) (res *desk.KnowledgeEditRes, err error)
		KnowledgePublishDel(ctx context.Context, req *desk.KnowledgeDelReq) (res *desk.KnowledgeDelRes, err error)

		KnowledgeBuy(ctx context.Context, req *desk.KnowledgeBuyReq) (res *desk.KnowledgeBuyRes, err error)
	}
)

var (
	localPublish IPublish
)

func Publish() IPublish {
	if localPublish == nil {
		panic("implement not found for interface IPublish, forgot register?")
	}
	return localPublish
}

// 注册发布接口
func RegisterPublish(i IPublish) {
	localPublish = i
}
