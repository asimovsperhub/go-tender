package service

import (
	"context"
	"tender/api/v1/desk"
)

// 接口层
type (
	IBbs interface {
		BbsPublish(ctx context.Context, req *desk.BbsPublishReq) (res *desk.BbsPublishRes, err error)
		BbsGetList(ctx context.Context, req *desk.BbsGetListReq) (res *desk.BbsGetListRes, err error)
		BbsGetListAll(ctx context.Context, req *desk.BbsGetListAllReq) (res *desk.BbsGetListAllRes, err error)
		BbsGet(ctx context.Context, req *desk.BbsGetReq) (res *desk.BbsGetRes, err error)
		BbsEdit(ctx context.Context, req *desk.BbsEditReq) (res *desk.BbsEditRes, err error)
		BbsDel(ctx context.Context, req *desk.BbsDelReq) (res *desk.BbsDelRes, err error)
		BbsBrowse(ctx context.Context, req *desk.BbsBrowseReq) (res *desk.BbsBrowseRes, err error)
		BbsLike(ctx context.Context, req *desk.BbsLikeReq) (res *desk.BbsLikeRes, err error)
		BbsComment(ctx context.Context, req *desk.BbsCommentReq) (res *desk.BbsCommentRes, err error)
		BbsCommentLike(ctx context.Context, req *desk.BbsCommentLikeReq) (res *desk.BbsCommentLikeRes, err error)
		BbsCommentDel(ctx context.Context, req *desk.BbsCommentDelReq) (res *desk.BbsCommentDelRes, err error)
		BbsReply(ctx context.Context, req *desk.BbsReplyReq) (res *desk.BbsReplyRes, err error)
		BbsGetLike(ctx context.Context, req *desk.BbsGetLikeReq) (res *desk.BbsGetLikeRes, err error)

		//FeedbackPublish
		FeedbackPublish(ctx context.Context, req *desk.FeedbackPublishReq) (res *desk.FeedbackPublishRes, err error)
	}
)

var (
	localBbs IBbs
)

func Bbs() IBbs {
	if localBbs == nil {
		panic("implement not found for interface IBbs, forgot register?")
	}
	return localBbs
}

// 注册发布接口
func RegisterBbs(i IBbs) {
	localBbs = i
}
