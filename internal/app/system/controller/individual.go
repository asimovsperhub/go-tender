package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/system/service"
)

var Enterprise = EnterpriseController{}

type EnterpriseController struct {
	BaseController
}

// 企业入驻
func (c *EnterpriseController) EnterpriseAdd(ctx context.Context, req *desk.EnterpriseAddReq) (res *desk.EnterpriseAddRes, err error) {
	res, err = service.Enterprise().Add(ctx, req)
	return
}

// 修改企业信息
func (c *EnterpriseController) EnterpriseEdit(ctx context.Context, req *desk.EnterpriseEditReq) (res *desk.EnterpriseEditRes, err error) {
	res, err = service.Enterprise().Edit(ctx, req)
	return
}

//EnterpriseGetReq
func (c *EnterpriseController) EnterpriseGet(ctx context.Context, req *desk.EnterpriseGetReq) (res *desk.EnterpriseGetRes, err error) {
	res, err = service.Enterprise().Get(ctx, req)
	return
}
