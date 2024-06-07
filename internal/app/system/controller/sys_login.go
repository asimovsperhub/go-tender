package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tender/api/v1/system"
	commonService "tender/internal/app/common/service"
	dModel "tender/internal/app/desk/model"
	deskService "tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/library/libUtils"
	"tender/internal/app/system/model"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"

	"github.com/chanxuehong/wechat/mp/qrcode"
)

var (
	Login = loginController{}
)

type loginController struct {
	BaseController
}

func (c *loginController) Login(ctx context.Context, req *system.UserLoginReq) (res *system.UserLoginRes, err error) {
	var (
		user        *model.LoginUserRes
		token       string
		permissions []string
		menuList    []*model.UserMenus
		roleids     []uint
	)
	//判断验证码是否正确
	//debug := gmode.IsDevelop()
	//if !debug {
	//	if !commonService.Captcha().VerifyString(req.VerifyKey, req.VerifyCode) {
	//		err = gerror.New("验证码输入错误")
	//		return
	//	}
	//}
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	user, err = service.SysUser().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		//// 保存登录失败的日志信息
		//service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
		//	Status:    0,
		//	Username:  req.Username,
		//	Ip:        ip,
		//	UserAgent: userAgent,
		//	Msg:       err.Error(),
		//	Module:    "系统后台",
		//})
		return
	}
	//err = service.SysUser().UpdateLoginInfo(ctx, user.Id, ip)
	//if err != nil {
	//	return
	//}
	// 报存登录成功的日志信息
	//service.SysLoginLog().Invoke(ctx, &model.LoginLogParams{
	//	Status:    1,
	//	Username:  req.Username,
	//	Ip:        ip,
	//	UserAgent: userAgent,
	//	Msg:       "登录成功",
	//	Module:    "系统后台",
	//})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败，后端服务出现错误")
		return
	}
	//获取用户菜单数据
	menuList, permissions, roleids, err = service.SysUser().GetAdminRules(ctx, user.Id)
	if err != nil {
		return
	}
	res = &system.UserLoginRes{
		UserInfo:    user,
		Token:       token,
		MenuList:    menuList,
		Permissions: permissions,
		RoleIds:     roleids,
	}
	//用户在线状态保存
	//service.SysUserOnline().Invoke(ctx, &model.SysUserOnlineParams{
	//	UserAgent: userAgent,
	//	Uuid:      gmd5.MustEncrypt(token),
	//	Token:     token,
	//	Username:  user.UserName,
	//	Ip:        ip,
	//})
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *system.UserLoginOutReq) (res *system.UserLoginOutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}

func (c *loginController) Qrcode(ctx context.Context, req *system.QrcodeReq) (res *system.QrcodeRes, err error) {

	//获取二维码
	res = &system.QrcodeRes{
		Qrcode: "",
	}
	senceId := uuid.New().String()
	resp, err := qrcode.CreateStrSceneTempQrcode(commonService.Mp(), senceId, 24*2600)
	if err != nil {
		return
	}
	url := qrcode.QrcodePicURL(resp.Ticket)
	res.Qrcode = url
	res.SenceId = senceId
	if err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysWxQrcode.Ctx(ctx).Insert(do.SysWxQrcode{
			SenceId:   senceId,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
	}); err != nil {
		return
	}
	return
}

func (c *loginController) CheckScanQrcode(ctx context.Context, req *system.CheckScanQrcodeReq) (res *system.CheckScanQrcodeRes, err error) {
	res = &system.CheckScanQrcodeRes{
		IsScan: false,
	}
	var user *dModel.LoginUserRes
	// 如果已登陆，绑定微信到用户
	context := new(model.Context)
	if data, err := service.GfToken().ParseToken(g.RequestFromCtx(ctx)); err == nil {
		if err = gconv.Struct(data.Data, &context.User); err == nil {
			//获取用户菜单数据

			model := &entity.SysWxQrcode{}
			if err := dao.SysWxQrcode.Ctx(ctx).Where(dao.SysWxQrcode.Columns().SenceId, req.SenceId).Scan(&model); err != nil {
				return nil, err
			}
			if model.IsScan == 1 {
				res.Token = model.Token
				res.IsScan = true

				wxUserModel := &entity.SysWxUser{}
				if err := dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, model.OpenId).Scan(&wxUserModel); err != nil && !errors.Is(err, sql.ErrNoRows) {
					return nil, err
				}
				// 没有绑定用户信息
				if wxUserModel.MainId == 0 {
					if _, err := dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, model.OpenId).Data(dao.SysWxUser.Columns().MainId, context.User.Id).Update(); err != nil {
						return nil, err
					}
					res.IsBind = true
				}
			}
		}
	}
	// 处理扫码登录
	model := &entity.SysWxQrcode{}
	if err := dao.SysWxQrcode.Ctx(ctx).Where(dao.SysWxQrcode.Columns().SenceId, req.SenceId).Scan(&model); err != nil {
		return nil, fmt.Errorf("SysWxQrcode not found: %w", err)
	}
	// 已经扫码
	if model.IsScan == 1 {
		res.Token = model.Token
		res.IsScan = true
		res.OpenId = model.OpenId

		wxUserModel := &entity.SysWxUser{}
		if err := dao.SysWxUser.Ctx(ctx).Where(dao.SysWxUser.Columns().OpenId, model.OpenId).Scan(&wxUserModel); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		// 没有绑定用户信息
		if wxUserModel.MainId == 0 {

		} else {
			member := &entity.MemberUser{}
			if err := dao.MemberUser.Ctx(ctx).Where(dao.MemberUser.Columns().Id, wxUserModel.MainId).Scan(&member); err != nil {
				return nil, err
			}
			res.IsBind = true
			user, err = deskService.DeskUser().GetUserByUsername(ctx, member.Mobile)
			res.UserInfo = user
		}
	}
	return
}
