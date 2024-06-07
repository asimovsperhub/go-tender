package controller

import (
	"context"
	"tender/api/v1/common"
	"tender/api/v1/system"
	"tender/internal/app/system/model"
	"tender/internal/app/system/service"
)

var Menu = menuController{}

type menuController struct {
	BaseController
}

func (c *menuController) List(ctx context.Context, req *system.RuleSearchReq) (res *system.RuleListRes, err error) {
	var list []*model.SysAuthRuleInfoRes
	data := make([]*model.SysAuthRuleTreeRes, 0)
	res = &system.RuleListRes{
		Result: common.Response{Data: data, Code: 0, Msg: "sussess"},
	}
	list, err = service.SysAuthRule().GetMenuListSearch(ctx, req)
	if req.Title != "" || req.Component != "" {
		for _, menu := range list {
			data = append(data, &model.SysAuthRuleTreeRes{
				SysAuthRuleInfoRes: menu,
			})
		}
		res.Result = common.Response{Data: data, Code: 0, Msg: "sussess"}
	} else {
		res.Result = common.Response{Data: service.SysAuthRule().GetMenuListTree(0, list), Code: 0, Msg: "sussess"}
	}
	return
}

func (c *menuController) Add(ctx context.Context, req *system.RuleAddReq) (res *system.RuleAddRes, err error) {
	err = service.SysAuthRule().Add(ctx, req)
	//res.Result={}
	return
}

// GetAddParams 获取菜单添加及编辑相关参数
func (c *menuController) GetAddParams(ctx context.Context, req *system.RuleGetParamsReq) (res *system.RuleGetParamsRes, err error) {
	// 获取角色列表
	res = new(system.RuleGetParamsRes)
	res.Roles, err = service.SysRole().GetRoleList(ctx)
	if err != nil {
		return
	}
	res.Menus, err = service.SysAuthRule().GetIsMenuList(ctx)
	return
}

// Get 获取菜单信息
func (c *menuController) Get(ctx context.Context, req *system.RuleInfoReq) (res *system.RuleInfoRes, err error) {
	res = new(system.RuleInfoRes)
	res.Rule, err = service.SysAuthRule().Get(ctx, req.Id)
	if err != nil {
		return
	}
	res.RoleIds, err = service.SysAuthRule().GetMenuRoles(ctx, req.Id)
	return
}

// Update 菜单修改
func (c *menuController) Update(ctx context.Context, req *system.RuleUpdateReq) (res *system.RuleUpdateRes, err error) {
	err = service.SysAuthRule().Update(ctx, req)
	return
}

// Delete 删除菜单
func (c *menuController) Delete(ctx context.Context, req *system.RuleDeleteReq) (res *system.RuleDeleteRes, err error) {
	err = service.SysAuthRule().DeleteMenuByIds(ctx, req.Ids)
	return
}
