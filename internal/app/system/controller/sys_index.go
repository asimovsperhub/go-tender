package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var SysIndexManger = SysIndexMangerController{}

type SysIndexMangerController struct {
	BaseController
}

// 最新入驻企业列表
func (c *SysIndexMangerController) EnterpriseList(ctx context.Context, req *system.EnterpriseReq) (res *system.EnterpriseRes, err error) {
	res, err = service.SysIndexManger().EnterpriseList(ctx, req)
	return
}

// 最新入驻会员列表
func (c *SysIndexMangerController) MemberList(ctx context.Context, req *system.MemberUserReq) (res *system.MemberUserRes, err error) {
	res, err = service.SysIndexManger().MemberList(ctx, req)
	return
}

func (c *SysIndexMangerController) UserCount(ctx context.Context, req *system.UserReq) (res *system.UserRes, err error) {
	res, err = service.SysIndexManger().UserCount(ctx, req)
	return
}

func (c *SysIndexMangerController) NumberCount(ctx context.Context, req *system.UserNumberReq) (res *system.UserNumberRes, err error) {
	res, err = service.SysIndexManger().NumberCount(ctx, req)
	return
}

//TrendingReq
func (c *SysIndexMangerController) Trending(ctx context.Context, req *system.TrendingReq) (res *system.TrendingRes, err error) {
	res, err = service.SysIndexManger().Trending(ctx, req)
	return
}
