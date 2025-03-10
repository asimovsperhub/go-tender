package controller

import (
	"context"
	"log"
	"tender/api/v1/system"
	"tender/internal/app/system/model"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
)

var (
	User = userController{}
)

type userController struct {
	BaseController
}

// GetUserMenus 获取用户菜单及按钮权限
func (c *userController) GetUserMenus(ctx context.Context, req *system.UserMenusReq) (res *system.UserMenusRes, err error) {
	var (
		permissions []string
		menuList    []*model.UserMenus
		roleids     []uint
	)
	userId := service.Context().GetUserId(ctx)
	menuList, permissions, roleids, err = service.SysUser().GetAdminRules(ctx, userId)
	log.Println("GetUserMenus------------>", menuList)
	res = &system.UserMenusRes{
		MenuList:    menuList,
		Permissions: permissions,
		RoleIds:     roleids,
	}
	return
}

// List 用户列表
func (c *userController) List(ctx context.Context, req *system.UserSearchReq) (res *system.UserSearchRes, err error) {
	var (
		total    interface{}
		userList []*entity.SysUser
	)
	res = new(system.UserSearchRes)
	total, userList, err = service.SysUser().List(ctx, req)
	if err != nil || total == 0 {
		return
	}
	res.Total = total
	res.UserList, err = service.SysUser().GetUsersRoleDept(ctx, userList)
	return
}

// Add 添加用户
func (c *userController) Add(ctx context.Context, req *system.UserAddReq) (res *system.UserAddRes, err error) {
	err = service.SysUser().Add(ctx, req)
	return
}

// GetEditUser 获取修改用户信息
func (c *userController) GetEditUser(ctx context.Context, req *system.UserGetEditReq) (res *system.UserGetEditRes, err error) {
	res, err = service.SysUser().GetEditUser(ctx, req.Id)
	return
}

// Edit 修改用户
func (c *userController) Edit(ctx context.Context, req *system.UserEditReq) (res *system.UserEditRes, err error) {
	err = service.SysUser().Edit(ctx, req)
	return
}

// ResetPwd 重置密码
func (c *userController) ResetPwd(ctx context.Context, req *system.UserResetPwdReq) (res *system.UserResetPwdRes, err error) {
	err = service.SysUser().ResetUserPwd(ctx, req)
	return
}

// SetStatus 修改用户状态
func (c *userController) SetStatus(ctx context.Context, req *system.UserStatusReq) (res *system.UserStatusRes, err error) {
	err = service.SysUser().ChangeUserStatus(ctx, req)
	return
}

// Delete 删除用户
func (c *userController) Delete(ctx context.Context, req *system.UserDeleteReq) (res *system.UserDeleteRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}

// GetUsers 通过用户id批量获取用户信息
func (c *userController) GetUsers(ctx context.Context, req *system.UserGetByIdsReq) (res *system.UserGetByIdsRes, err error) {
	res = new(system.UserGetByIdsRes)
	res.List, err = service.SysUser().GetUsers(ctx, req.Ids)
	return
}
