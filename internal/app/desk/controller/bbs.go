package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var (
	Bbs = cBbs{}
)

type cBbs struct {
}

func (*cBbs) BbsPublish(ctx context.Context, req *desk.BbsPublishReq) (res *desk.BbsPublishRes, err error) {
	res, err = service.Bbs().BbsPublish(ctx, req)
	return
}

//BbsGetListReq

func (*cBbs) BbsGetList(ctx context.Context, req *desk.BbsGetListReq) (res *desk.BbsGetListRes, err error) {
	res, err = service.Bbs().BbsGetList(ctx, req)
	return
}

//BbsGetListAllReq
func (*cBbs) BbsGetListAll(ctx context.Context, req *desk.BbsGetListAllReq) (res *desk.BbsGetListAllRes, err error) {
	res, err = service.Bbs().BbsGetListAll(ctx, req)
	return
}

//BbsGetReq

func (*cBbs) BbsGet(ctx context.Context, req *desk.BbsGetReq) (res *desk.BbsGetRes, err error) {
	res, err = service.Bbs().BbsGet(ctx, req)
	return
}

//BbsEditReq
func (*cBbs) BbsEdit(ctx context.Context, req *desk.BbsEditReq) (res *desk.BbsEditRes, err error) {
	res, err = service.Bbs().BbsEdit(ctx, req)
	return
}

//BbsDelReq
func (*cBbs) BbsDel(ctx context.Context, req *desk.BbsDelReq) (res *desk.BbsDelRes, err error) {
	res, err = service.Bbs().BbsDel(ctx, req)
	return
}

//BbsBrowseReq
func (*cBbs) BbsBrowse(ctx context.Context, req *desk.BbsBrowseReq) (res *desk.BbsBrowseRes, err error) {
	res, err = service.Bbs().BbsBrowse(ctx, req)
	return
}

//BbsLikeReq
func (*cBbs) BbsLike(ctx context.Context, req *desk.BbsLikeReq) (res *desk.BbsLikeRes, err error) {
	res, err = service.Bbs().BbsLike(ctx, req)
	return
}

//BbsCommentReq
func (*cBbs) BbsComment(ctx context.Context, req *desk.BbsCommentReq) (res *desk.BbsCommentRes, err error) {
	res, err = service.Bbs().BbsComment(ctx, req)
	return
}

//BbsCommentLikeReq
func (*cBbs) BbsCommentLikeReq(ctx context.Context, req *desk.BbsCommentLikeReq) (res *desk.BbsCommentLikeRes, err error) {
	res, err = service.Bbs().BbsCommentLike(ctx, req)
	return
}

//BbsCommentDelReq
func (*cBbs) BbsCommentDel(ctx context.Context, req *desk.BbsCommentDelReq) (res *desk.BbsCommentDelRes, err error) {
	res, err = service.Bbs().BbsCommentDel(ctx, req)
	return
}

//BbsReplyReq
func (*cBbs) BbsReply(ctx context.Context, req *desk.BbsReplyReq) (res *desk.BbsReplyRes, err error) {
	res, err = service.Bbs().BbsReply(ctx, req)
	return
}

//BbsGetLikeReq
func (*cBbs) BbsGetLike(ctx context.Context, req *desk.BbsGetLikeReq) (res *desk.BbsGetLikeRes, err error) {
	res, err = service.Bbs().BbsGetLike(ctx, req)
	return
}

//FeedbackPublishReq

func (*cBbs) FeedbackPublish(ctx context.Context, req *desk.FeedbackPublishReq) (res *desk.FeedbackPublishRes, err error) {
	res, err = service.Bbs().FeedbackPublish(ctx, req)
	return
}
