package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var MemberManger = MemberMangerController{}

type MemberMangerController struct {
	BaseController
}

func (c *MemberMangerController) List(ctx context.Context, req *system.MemberUserSearchReq) (res *system.MemberUserSearchRes, err error) {
	res, err = service.SysMemberManger().List(ctx, req)
	return
}

//MemberUserEditReq
func (c *MemberMangerController) MemberUserEdit(ctx context.Context, req *system.MemberUserEditReq) (res *system.MemberUserEditRes, err error) {
	res, err = service.SysMemberManger().MemberUserEdit(ctx, req)
	return
}

// Disable
func (c *MemberMangerController) Disable(ctx context.Context, req *system.DisableMemberUserReq) (res *system.DisableMemberUserRes, err error) {
	res, err = service.SysMemberManger().Disable(ctx, req)
	return
}

func (c *MemberMangerController) EditFee(ctx context.Context, req *system.MemberFeeReq) (res *system.MemberFeeRes, err error) {
	res, err = service.SysMemberManger().EditFee(ctx, req)
	return
}

func (c *MemberMangerController) EditIn(ctx context.Context, req *system.MemberIntegralReq) (res *system.MemberIntegralRes, err error) {
	res, err = service.SysMemberManger().EditIn(ctx, req)
	return
}

func (c *MemberMangerController) EditSu(ctx context.Context, req *system.MemberSubscriptionReq) (res *system.MemberSubscriptionRes, err error) {
	res, err = service.SysMemberManger().EditSu(ctx, req)
	return
}

func (c *MemberMangerController) FindFee(ctx context.Context, req *system.MemberFeeFindReq) (res *system.MemberFeeFindRes, err error) {
	res, err = service.SysMemberManger().FindFee(ctx, req)
	return
}

func (c *MemberMangerController) FindIn(ctx context.Context, req *system.MemberInFindReq) (res *system.MemberInFindRes, err error) {
	res, err = service.SysMemberManger().FindIn(ctx, req)
	return
}

func (c *MemberMangerController) FindSu(ctx context.Context, req *system.MemberSuFindReq) (res *system.MemberSuFindRes, err error) {
	res, err = service.SysMemberManger().FindSu(ctx, req)
	return
}
