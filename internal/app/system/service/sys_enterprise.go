package service

import (
	"context"
	"tender/api/v1/system"
)

type IEnterpriseManger interface {
	List(ctx context.Context, req *system.EnterpriseSearchReq) (res *system.EnterpriseSearchRes, err error)
	LicenseList(ctx context.Context, req *system.EnterpriselLcenseSearchReq) (res *system.EnterpriseSearchRes, err error)
	CertificateList(ctx context.Context, req *system.EnterpriseCertificateSearchReq) (res *system.EnterpriseSearchRes, err error)
	LicenseReview(ctx context.Context, req *system.EnterpriselLcenseReviewReq) (res *system.EnterpriselLcenseReviewRes, err error)
	CertificateReview(ctx context.Context, req *system.EnterpriseCertificateReviewReq) (res *system.EnterpriseCertificateReviewRes, err error)
}

var localEnterpriseManger IEnterpriseManger

func SysEnterpriseManger() IEnterpriseManger {
	if localEnterpriseManger == nil {
		panic("implement not found for interface EnterpriseManger, forgot register?")
	}
	return localEnterpriseManger
}

func RegisterEnterpriseManger(i IEnterpriseManger) {
	localEnterpriseManger = i
}
