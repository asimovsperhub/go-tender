package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"log"
	"tender/api/v1/desk"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/library/libUtils"
	"tender/library/libmessage"
)

/*
	个人信息
*/
var (
	UserInfo = UserInfoController{}
)

type UserInfoController struct {
	BaseController
}

func (c *UserInfoController) UserPassEdit(ctx context.Context, req *desk.UserPassEditReq) (res *desk.UserPassEditRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	finduser := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%s'",
			dao.MemberUser.Columns().Mobile,
			req.Mobile))
		err = m.Limit(1).Scan(&finduser)
	})
	if finduser == nil {
		err = errors.New("手机号不存在")
		return
	}
	send := libmessage.GetEntity()
	err = send.CheckVerificationCode(req.Mobile, req.Code)
	if err != nil {
		return
	}

	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.DeskUser().GetUserByUsername(ctx, finduser.Mobile)
	if err != nil {
		// g.Log().Error(ctx, err)
		err = errors.New("获取用户信息失败")
		return
	}
	//账号状态 用户状态;0:禁用,1:正常,2:未验证
	if user.UserStatus == 0 {
		err = errors.New("账号已被冻结")
	}
	UserSalt := grand.S(10)
	req.Password = libUtils.EncryptPassword(req.Password, UserSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberUser.Ctx(ctx).TX(tx).Where("mobile = ?", req.Mobile).Update(do.MemberUser{
				UserPassword: req.Password,
				UserSalt:     UserSalt,
				UpdatedAt:    gtime.Now(),
			})
			if e != nil {
				err = errors.New("更改用户密码失败")
			}
		})
		return err
	})
	log.Println("Auth---------------------------->", req.Authorization)
	// 删除token
	// service.GfToken().RemoveToken(ctx,req.Authorization)
	// 用新密码生成token
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.Mobile) + gmd5.MustEncryptString(req.Password)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.Mobile) + gmd5.MustEncryptString(req.Password+ip+userAgent)
	}
	// 不能返回用户密码
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = errors.New("登录失败，后端服务出现错误")
		return
	}
	res = &desk.UserPassEditRes{
		Result:   "修改密码成功",
		Token:    token,
		UserInfo: user,
	}
	return
}

func (c *UserInfoController) UserEdit(ctx context.Context, req *desk.UserEditReq) (res *desk.UserEditRes, err error) {
	var (
		user *model.LoginUserRes
		// token string
	)
	user_ := service.Context().GetLoginUser(ctx)
	if user_ == nil {
		err = errors.New("用户未登陆")
		return
	}
	finduser := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%s'",
			dao.MemberUser.Columns().Mobile,
			*req.Mobile))
		err = m.Limit(1).Scan(&finduser)
	})
	user_id := service.Context().GetUserId(ctx)
	user_info := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			dao.MemberUser.Columns().Id,
			user_id))
		err = m.Limit(1).Scan(&user_info)
	})

	if finduser != nil {
		if user_info.Mobile != finduser.Mobile {
			err = errors.New("手机号已注册")
			return
		}
	}
	//账号状态 用户状态;0:禁用,1:正常,2:未验证
	if user_info.UserStatus == 0 {
		err = errors.New("账号已被冻结")
	}
	data := do.MemberUser{}
	if req.Mobile != nil {
		data.Mobile = *req.Mobile
	}
	if req.NickName != nil {
		data.UserNickname = *req.NickName
	}
	if req.Email != nil {
		data.UserEmail = *req.Email
	}
	if req.Avatar != nil {
		data.Avatar = req.Avatar
	}
	if req.Address != nil {
		data.Address = *req.Address
	}
	if req.Describe != nil {
		data.Describe = *req.Describe
	}
	data.UpdatedAt = gtime.Now()
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.MemberUser.Ctx(ctx).TX(tx).Where("mobile = ?", user_info.Mobile).Update(data)
			if e != nil {
				err = errors.New("编辑用户失败")
			}
		})
		return err
	})
	user, err = service.DeskUser().GetUserByUsername(ctx, *req.Mobile)
	if err != nil {
		// g.Log().Error(ctx, err)
		err = errors.New("获取用户信息失败")
		return
	}
	// 不能返回用户密码
	user.UserPassword = ""
	res = &desk.UserEditRes{
		Result: "修改用户信息成功",
		// Token:    token,
		UserInfo: user,
	}
	// 修改的手机号不是当前的手机清除token
	if *req.Mobile != user_info.Mobile {
		err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	}
	return
}

func (c *UserInfoController) User(ctx context.Context, req *desk.UserReq) (res *desk.UserRes, err error) {
	user_id := service.Context().GetUserId(ctx)
	finduser := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			dao.MemberUser.Columns().Id,
			user_id))
		err = m.Limit(1).Scan(&finduser)
	})
	if finduser == nil {
		err = errors.New("用户不存在")
		return
	}
	//账号状态 用户状态;0:禁用,1:正常,2:未验证
	if finduser.UserStatus == 0 {
		err = errors.New("账号已被冻结")
	}
	// 不能返回用户密码
	finduser.UserPassword = ""
	wxuser := (*entity.SysWxUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysWxUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			dao.SysWxUser.Columns().MainId,
			user_id))
		err = m.Limit(1).Scan(&wxuser)
	})
	res = &desk.UserRes{
		UserInfo: finduser,
		WxInfo:   wxuser,
	}
	return
}
