package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"tender/api/v1/system"
	"tender/internal/app/system/model"
	"tender/internal/app/system/model/entity"
)

// 接口层
type (
	ISysUser interface {
		GetCasBinUserPrefix() string
		NotCheckAuthAdminIds(ctx context.Context) *gset.Set
		GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error)
		GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error)
		// LoginLog(ctx context.Context, params *model.LoginLogParams)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		GetAdminRules(ctx context.Context, userId uint64) (menuList []*model.UserMenus, permissions []string, roleids []uint, err error)
		GetAdminRole(ctx context.Context, userId uint64, allRoleList []*entity.SysRole) (roles []*entity.SysRole, err error)
		GetAdminRoleIds(ctx context.Context, userId uint64) (roleIds []uint, err error)
		GetAllMenus(ctx context.Context) (menus []*model.UserMenus, err error)
		GetAdminMenusByRoleIds(ctx context.Context, roleIds []uint) (menus []*model.UserMenus, err error)
		GetMenusTree(menus []*model.UserMenus, pid uint) []*model.UserMenus
		GetPermissions(ctx context.Context, roleIds []uint) (userButtons []string, err error)
		List(ctx context.Context, req *system.UserSearchReq) (total interface{}, userList []*entity.SysUser, err error)
		GetUsersRoleDept(ctx context.Context, userList []*entity.SysUser) (users []*model.SysUserRoleDeptRes, err error)
		Add(ctx context.Context, req *system.UserAddReq) (err error)
		Edit(ctx context.Context, req *system.UserEditReq) (err error)
		AddUserPost(ctx context.Context, tx gdb.TX, postIds []int64, userId int64) (err error)
		EditUserRole(ctx context.Context, roleIds []int64, userId int64) (err error)
		UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error
		GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error)
		GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error)
		GetUserPostIds(ctx context.Context, userId uint64) (postIds []int64, err error)
		ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)
		ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error)
		Delete(ctx context.Context, ids []int) (err error)
		GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimpleRes, err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

// 用户注册
func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
