package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var EnterpriseManger = EnterpriseMangerController{}

type EnterpriseMangerController struct {
	BaseController
}

func (c *EnterpriseMangerController) List(ctx context.Context, req *system.EnterpriseSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res, err = service.SysEnterpriseManger().List(ctx, req)
	return
}
func (c *EnterpriseMangerController) LcenseList(ctx context.Context, req *system.EnterpriselLcenseSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res, err = service.SysEnterpriseManger().LicenseList(ctx, req)
	return
}
func (c *EnterpriseMangerController) CertificateList(ctx context.Context, req *system.EnterpriseCertificateSearchReq) (res *system.EnterpriseSearchRes, err error) {
	res, err = service.SysEnterpriseManger().CertificateList(ctx, req)
	return
}
func (c *EnterpriseMangerController) LicenseReview(ctx context.Context, req *system.EnterpriselLcenseReviewReq) (res *system.EnterpriselLcenseReviewRes, err error) {
	res, err = service.SysEnterpriseManger().LicenseReview(ctx, req)
	return
}
func (c *EnterpriseMangerController) CertificateReview(ctx context.Context, req *system.EnterpriseCertificateReviewReq) (res *system.EnterpriseCertificateReviewRes, err error) {
	res, err = service.SysEnterpriseManger().CertificateReview(ctx, req)
	return
}
