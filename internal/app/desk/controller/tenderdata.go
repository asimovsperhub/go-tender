package controller

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/service"
)

var TnderData = deskTenderDataController{}

type deskTenderDataController struct {
	BaseController
}

// List 招标数据
func (c *deskTenderDataController) List(ctx context.Context, req *desk.DataSearchReq) (res *desk.DataSearchRes, err error) {
	res, err = service.DeskTenderDataManger().List(ctx, req)
	return
}
