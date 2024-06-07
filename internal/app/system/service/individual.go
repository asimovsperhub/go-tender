package service

import (
	"context"
	"tender/api/v1/desk"
)

type IEnterprise interface {
	Add(ctx context.Context, req *desk.EnterpriseAddReq) (res *desk.EnterpriseAddRes, err error)
	Edit(ctx context.Context, req *desk.EnterpriseEditReq) (res *desk.EnterpriseEditRes, err error)
	Get(ctx context.Context, req *desk.EnterpriseGetReq) (res *desk.EnterpriseGetRes, err error)
}

var localEnterprise IEnterprise

func Enterprise() IEnterprise {
	if localEnterprise == nil {
		panic("implement not found for interface Enterprise, forgot register?")
	}
	return localEnterprise
}

func RegisterEnterprise(i IEnterprise) {
	localEnterprise = i
}
