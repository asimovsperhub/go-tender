package service

import (
	"context"
	"tender/api/v1/system"
)

type ISysMemberManger interface {
	List(ctx context.Context, req *system.MemberUserSearchReq) (res *system.MemberUserSearchRes, err error)
	MemberUserEdit(ctx context.Context, req *system.MemberUserEditReq) (res *system.MemberUserEditRes, err error)
	Disable(ctx context.Context, req *system.DisableMemberUserReq) (res *system.DisableMemberUserRes, err error)
	EditFee(ctx context.Context, req *system.MemberFeeReq) (res *system.MemberFeeRes, err error)
	EditIn(ctx context.Context, req *system.MemberIntegralReq) (res *system.MemberIntegralRes, err error)
	EditSu(ctx context.Context, req *system.MemberSubscriptionReq) (res *system.MemberSubscriptionRes, err error)
	FindFee(ctx context.Context, req *system.MemberFeeFindReq) (res *system.MemberFeeFindRes, err error)
	FindIn(ctx context.Context, req *system.MemberInFindReq) (res *system.MemberInFindRes, err error)
	FindSu(ctx context.Context, req *system.MemberSuFindReq) (res *system.MemberSuFindRes, err error)
}

var localMemberManger ISysMemberManger

func SysMemberManger() ISysMemberManger {
	if localMemberManger == nil {
		panic("implement not found for interface SysMemberManger, forgot register?")
	}
	return localMemberManger
}

func RegisterMemberManger(i ISysMemberManger) {
	localMemberManger = i
}
