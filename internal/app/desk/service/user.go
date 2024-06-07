package service

import (
	"context"
	"tender/api/v1/desk"
	"tender/internal/app/desk/model"
	"tender/internal/app/system/model/entity"
)

// 接口层
type (
	IUser interface {
		RegisterBind(ctx context.Context, pwd, mobile, nickname, name, openId string) (err error)
		GetUserByUsernamePassword(ctx context.Context, req *desk.UserLoginReq) (res *model.LoginUserRes, err error)
		GetUserByUsername(ctx context.Context, Mobile string) (user *model.LoginUserRes, err error)
		GetEnterpriseByUserId(ctx context.Context, UserId int) (enterprise *entity.SysEnterprise, err error)
	}
)

var (
	localUser IUser
)

func DeskUser() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// 用户注册
func RegisterUser(i IUser) {
	localUser = i
}
