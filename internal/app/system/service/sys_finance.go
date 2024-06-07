package service

import (
	"context"
	"tender/api/v1/system"
)

type IFinance interface {
	PurchaseList(ctx context.Context, req *system.FinanceReq) (res *system.FinanceRes, err error)
	PurchaseInfo(ctx context.Context, req *system.PurchaseInfoReq) (res *system.PurchaseInfoRes, err error)
	TrendingInfo(ctx context.Context, req *system.TrendingInfoReq) (res *system.TrendingInfoRes, err error)
	QueryPaySettings(ctx context.Context, req *system.PaySettingsReq) (res *system.PaySettingsRes, err error)
	UpdatePaySettings(ctx context.Context, req *system.UpdatePaySettingsReq) (res *system.UpdatePaySettingsRes, err error)
}

var localFinance IFinance

func SysFinance() IFinance {
	if localFinance == nil {
		panic("implement not found for interface localFinance, forgot register?")
	}
	return localFinance
}

func RegisterFinance(i IFinance) {
	localFinance = i
}
