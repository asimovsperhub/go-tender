package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
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
	"tender/library/liberr"
	"tender/library/libmessage"
)

var (
	Login = loginController{}
)

type loginController struct {
	BaseController
}

func (c *loginController) Login(ctx context.Context, req *desk.UserLoginReq) (res *desk.UserLoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	finduser := (*entity.MemberUser)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%s'",
			//dao.MemberUser.Columns().UserName,
			//req.Username,
			dao.MemberUser.Columns().Mobile,
			req.Mobile))
		err := m.Limit(1).Scan(&finduser)
		liberr.ErrIsNil(ctx, err, "用户名或手机号不存在")
		if finduser == nil {
			return
		}
	})
	if err != nil {
		return
	}
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.DeskUser().GetUserByUsernamePassword(ctx, req)
	if err != nil {
		// g.Log().Error(ctx, err)
		err = gerror.New("账号或密码错误")
		return
	}
	if user.ReleaseAt != nil {
		if gtime.Now().Unix() >= user.ReleaseAt.Unix() {
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.MemberUser.Ctx(ctx).WherePri(user.Id).Update(do.MemberUser{
					// 用户状态
					UserStatus: 1,
				})
				liberr.ErrIsNil(ctx, err, "会员解禁更新失败")
				return
			})
		}
	}
	//账号状态 用户状态;0:禁用,1:正常,2:未验证
	if user.UserStatus == 0 {
		err = gerror.New("账号已被冻结")
		return
	}
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	// 不能返回用户密码
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	res = &desk.UserLoginRes{
		UserInfo: user,
		Token:    token,
	}
	return
}

func (c *loginController) MsgLogin(ctx context.Context, req *desk.UserLoginMessageReq) (res *desk.UserLoginMessageRes, err error) {
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
		liberr.ErrIsNil(ctx, err, "手机号不存在")
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
	if err != nil {
		return
	}

	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.DeskUser().GetUserByUsername(ctx, finduser.Mobile)
	if err != nil {
		// g.Log().Error(ctx, err)
		err = gerror.New("账号或密码错误")
		return
	}
	if user.ReleaseAt != nil {
		if gtime.Now().Unix() >= user.ReleaseAt.Unix() {
			err = g.Try(ctx, func(ctx context.Context) {
				_, err = dao.MemberUser.Ctx(ctx).WherePri(user.Id).Update(do.MemberUser{
					// 用户状态
					UserStatus: 1,
				})
				liberr.ErrIsNil(ctx, err, "会员解禁更新失败")
			})
		}
	}
	//账号状态 用户状态;0:禁用,1:正常,2:未验证
	if user.UserStatus == 0 {
		err = gerror.New("账号已被冻结")
		return
	}
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.Mobile) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.Mobile) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	// 不能返回用户密码
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	res = &desk.UserLoginMessageRes{
		UserInfo: user,
		Token:    token,
	}
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *desk.UserLoginOutReq) (res *desk.UserLoginOutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}

func (c *loginController) UserPassEdit(ctx context.Context, req *desk.UserForgetPassReq) (res *desk.UserForgetPassRes, err error) {
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
	res = &desk.UserForgetPassRes{
		Result:   "修改密码成功",
		Token:    token,
		UserInfo: user,
	}
	return
}
