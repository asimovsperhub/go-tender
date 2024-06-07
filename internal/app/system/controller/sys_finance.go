package controller

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/service"
)

var SysFinanceManger = SysFinanceController{}

type SysFinanceController struct {
	BaseController
}

func (c *SysFinanceController) PurchaseList(ctx context.Context, req *system.FinanceReq) (res *system.FinanceRes, err error) {
	res, err = service.SysFinance().PurchaseList(ctx, req)
	return
}

func (c *SysFinanceController) PurchaseInfo(ctx context.Context, req *system.PurchaseInfoReq) (res *system.PurchaseInfoRes, err error) {
	res, err = service.SysFinance().PurchaseInfo(ctx, req)
	return
}

func (c *SysFinanceController) TrendingInfo(ctx context.Context, req *system.TrendingInfoReq) (res *system.TrendingInfoRes, err error) {
	res, err = service.SysFinance().TrendingInfo(ctx, req)
	return
}

func (c *SysFinanceController) QueryPaySettings(ctx context.Context, req *system.PaySettingsReq) (res *system.PaySettingsRes, err error) {
	res, err = service.SysFinance().QueryPaySettings(ctx, req)
	return
}

func (c *SysFinanceController) UpdatePaySettings(ctx context.Context, req *system.UpdatePaySettingsReq) (res *system.UpdatePaySettingsRes, err error) {
	res, err = service.SysFinance().UpdatePaySettings(ctx, req)
	return
}
