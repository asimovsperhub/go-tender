package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var KnowledgeManger = KnowledgeMangerController{}

type KnowledgeMangerController struct {
	BaseController
}

func (c *KnowledgeMangerController) List(ctx context.Context, req *system.KnowledgeSearchReq) (res *system.KnowledgeSearchRes, err error) {
	res, err = service.SysKnowledgeManger().List(ctx, req)
	return
}

//KnowledgeReq
func (c *KnowledgeMangerController) ImportList(ctx context.Context, req *system.KnowledgeReq) (res *system.KnowledgeRes, err error) {
	res, err = service.SysKnowledgeManger().ImportList(ctx, req)
	return
}
func (c *KnowledgeMangerController) Review(ctx context.Context, req *system.KnowledgeReviewReq) (res *system.KnowledgeReviewRes, err error) {
	res, err = service.SysKnowledgeManger().KnowledgeReview(ctx, req)
	return
}

func (c *KnowledgeMangerController) Del(ctx context.Context, req *system.KnowledgeDelReq) (res *system.KnowledgeDelRes, err error) {
	res, err = service.SysKnowledgeManger().Del(ctx, req)
	return
}

func (c *KnowledgeMangerController) KnowledgeList(ctx context.Context, req *system.KnowledgeReviewSearchReq) (res *system.KnowledgeReviewSearchRes, err error) {
	res, err = service.SysKnowledgeManger().KnowledgeList(ctx, req)
	return
}

func (c *KnowledgeMangerController) ProcessVideo(ctx context.Context, req *system.KnowledgeProcessVideoReq) (res *system.KnowledgeProcessVideoRes, err error) {
	res, err = service.SysKnowledgeManger().ProcessVideo(ctx, req)
	return
}

func (s *KnowledgeMangerController) ProcessPdf(ctx context.Context, req *system.KnowledgeProcessPdfReq) (res *system.KnowledgeProcessPdfRes, err error) {
	res, err = service.SysKnowledgeManger().ProcessPdf(ctx, req)
	return
}
