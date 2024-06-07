package service

import (
	"context"
	"tender/api/v1/system"
)

type IIndexManger interface {
	EnterpriseList(ctx context.Context, req *system.EnterpriseReq) (res *system.EnterpriseRes, err error)
	MemberList(ctx context.Context, req *system.MemberUserReq) (res *system.MemberUserRes, err error)
	UserCount(ctx context.Context, req *system.UserReq) (res *system.UserRes, err error)
	NumberCount(ctx context.Context, req *system.UserNumberReq) (res *system.UserNumberRes, err error)
	Trending(ctx context.Context, req *system.TrendingReq) (res *system.TrendingRes, err error)
}

var localIndexManger IIndexManger

func SysIndexManger() IIndexManger {
	if localIndexManger == nil {
		panic("implement not found for interface SysIndexManger, forgot register?")
	}
	return localIndexManger
}

func RegisterIndexManger(i IIndexManger) {
	localIndexManger = i
}
