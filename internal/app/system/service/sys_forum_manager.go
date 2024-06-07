package service

import (
	"context"
	"tender/api/v1/system"
)

type IForumManger interface {
	BbsGetListAll(ctx context.Context, req *system.BbsGetListAllReq) (res *system.BbsGetListAllRes, err error)
	BbsGet(ctx context.Context, req *system.BbsGetReq) (res *system.BbsGetRes, err error)
	BbsDel(ctx context.Context, req *system.BbsDelReq) (res *system.BbsDelRes, err error)
	BbsRestore(ctx context.Context, req *system.BbsRestoreReq) (res *system.BbsRestoreRes, err error)
	BbsForever(ctx context.Context, req *system.BbsForeverReq) (res *system.BbsForeverRes, err error)
	BbsGetReviewList(ctx context.Context, req *system.BbsGetReviewListReq) (res *system.BbsGetReviewListRes, err error)
	BbsReview(ctx context.Context, req *system.BbsReviewReq) (res *system.BbsReviewRes, err error)
	BbsTop(ctx context.Context, req *system.BbsTopReq) (res *system.BbsTopRes, err error)
}

var localForumManger IForumManger

func SysForumManger() IForumManger {
	if localForumManger == nil {
		panic("implement not found for interface SysForumManger, forgot register?")
	}
	return localForumManger
}

func RegisterForumManger(i IForumManger) {
	localForumManger = i
}
