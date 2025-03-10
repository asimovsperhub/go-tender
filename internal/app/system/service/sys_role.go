package service

import (
	"context"
	"tender/api/v1/system"
	"tender/internal/app/system/model/entity"
)

type (
	ISysRole interface {
		GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
		GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
		AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error)
		DelRoleRule(ctx context.Context, roleId int64) (err error)
		AddRole(ctx context.Context, req *system.RoleAddReq) (err error)
		Get(ctx context.Context, id uint) (res *entity.SysRole, err error)
		GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error)
		EditRole(ctx context.Context, req *system.RoleEditReq) (err error)
		DeleteByIds(ctx context.Context, ids []int64) (err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
