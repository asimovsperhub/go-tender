package service

import (
	"context"
	"tender/api/v1/system"
)

type IKnowledgeManger interface {
	List(ctx context.Context, req *system.KnowledgeSearchReq) (res *system.KnowledgeSearchRes, err error)
	ImportList(ctx context.Context, req *system.KnowledgeReq) (res *system.KnowledgeRes, err error)
	KnowledgeReview(ctx context.Context, req *system.KnowledgeReviewReq) (res *system.KnowledgeReviewRes, err error)
	KnowledgeList(ctx context.Context, req *system.KnowledgeReviewSearchReq) (res *system.KnowledgeReviewSearchRes, err error)
	Del(ctx context.Context, req *system.KnowledgeDelReq) (res *system.KnowledgeDelRes, err error)
	ProcessVideo(ctx context.Context, req *system.KnowledgeProcessVideoReq) (res *system.KnowledgeProcessVideoRes, err error)
	ProcessPdf(ctx context.Context, req *system.KnowledgeProcessPdfReq) (res *system.KnowledgeProcessPdfRes, err error)
}

var localKnowledgeManger IKnowledgeManger

func SysKnowledgeManger() IKnowledgeManger {
	if localKnowledgeManger == nil {
		panic("implement not found for interface SysKnowledgeManger, forgot register?")
	}
	return localKnowledgeManger
}

func RegisterKnowledgeManger(i IKnowledgeManger) {
	localKnowledgeManger = i
}
