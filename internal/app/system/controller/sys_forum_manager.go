package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var ForumManger = ForumMangerController{}

type ForumMangerController struct {
	BaseController
}

func (c *ForumMangerController) List(ctx context.Context, req *system.BbsGetListAllReq) (res *system.BbsGetListAllRes, err error) {
	res, err = service.SysForumManger().BbsGetListAll(ctx, req)
	return
}

//BbsGetReq
func (c *ForumMangerController) BbsGet(ctx context.Context, req *system.BbsGetReq) (res *system.BbsGetRes, err error) {
	res, err = service.SysForumManger().BbsGet(ctx, req)
	return
}

//BbsDelReq
func (c *ForumMangerController) BbsDel(ctx context.Context, req *system.BbsDelReq) (res *system.BbsDelRes, err error) {
	res, err = service.SysForumManger().BbsDel(ctx, req)
	return
}

//BbsRestoreReq
func (c *ForumMangerController) BbsRestore(ctx context.Context, req *system.BbsRestoreReq) (res *system.BbsRestoreRes, err error) {
	res, err = service.SysForumManger().BbsRestore(ctx, req)
	return
}

//BbsForeverReq
func (c *ForumMangerController) BbsForever(ctx context.Context, req *system.BbsForeverReq) (res *system.BbsForeverRes, err error) {
	res, err = service.SysForumManger().BbsForever(ctx, req)
	return
}

//BbsGetReviewListReq

func (c *ForumMangerController) BbsGetReviewList(ctx context.Context, req *system.BbsGetReviewListReq) (res *system.BbsGetReviewListRes, err error) {
	res, err = service.SysForumManger().BbsGetReviewList(ctx, req)
	return
}

//BbsReviewReq

func (c *ForumMangerController) BbsReview(ctx context.Context, req *system.BbsReviewReq) (res *system.BbsReviewRes, err error) {
	res, err = service.SysForumManger().BbsReview(ctx, req)
	return
}

//BbsTopReq

func (c *ForumMangerController) BbsTop(ctx context.Context, req *system.BbsTopReq) (res *system.BbsTopRes, err error) {
	res, err = service.SysForumManger().BbsTop(ctx, req)
	return
}
