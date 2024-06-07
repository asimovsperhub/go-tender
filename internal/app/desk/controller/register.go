package controller

import (
	"context"
	"fmt"
	"log"
	"tender/api/v1/desk"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/service"
	"tender/library/libUtils"
	"tender/library/libmessage"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Register = cRegister{}
)

type cRegister struct {
}

func (*cRegister) SendCode(ctx context.Context, req *desk.SendCodeReq) (res *desk.SendCodeRes, err error) {

	//finduser := (*entity.MemberUser)(nil)
	//err = g.Try(ctx, func(ctx context.Context) {
	//	m := dao.MemberUser.Ctx(ctx)
	//	m = m.Where(fmt.Sprintf("%s='%s'",
	//		dao.MemberUser.Columns().Mobile,
	//		req.Mobile))
	//	err := m.Limit(1).Scan(&finduser)
	//	liberr.ErrIsNil(ctx, err, "手机号已注册")
	//	log.Println("--------------------------->", finduser.Mobile)
	//})
	//if finduser != nil {
	//	res = &desk.SendCodeRes{
	//		Result: "手机号已注册",
	//	}
	//	return res, nil
	//}
	send := libmessage.GetEntity()
	err_sc := send.SendCode(req.Mobile)
	if err_sc != nil {
		log.Println("SendCode err ------------------>", err_sc)
		res = &desk.SendCodeRes{
			Result: "发送失败",
		}
		return res, err_sc
	}
	res = &desk.SendCodeRes{
		Result: "验证码发送成功",
	}
	return res, nil
}

func (*cRegister) BindWx(ctx context.Context, req *desk.BindWxReq) (res *desk.RegisterDoRes, err error) {
	send := libmessage.GetEntity()
	err = send.CheckVerificationCode(req.Mobile, req.Code)
	if err != nil {
		return
	}

	// 注册 service
	if err = service.DeskUser().RegisterBind(ctx, req.Password, req.Mobile, req.Nickname, req.Name, req.OpenId); err != nil {
		return &desk.RegisterDoRes{
			Result:   fmt.Sprintf("register failed:%s", err),
			UserInfo: nil,
			Token:    ""}, nil
	} else {
		// 自动登录
		var (
			user  *model.LoginUserRes
			token string
		)
		ip := libUtils.GetClientIp(ctx)
		userAgent := libUtils.GetUserAgent(ctx)
		user, err = service.DeskUser().GetUserByUsernamePassword(ctx, &desk.UserLoginReq{Mobile: req.Mobile, Password: req.Password})
		if err != nil {
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
		return &desk.RegisterDoRes{
			Result:   "register success login success",
			UserInfo: user,
			Token:    token}, nil
	}
}

func (*cRegister) RegisterBind(ctx context.Context, req *desk.RegisterDoReq) (res *desk.RegisterDoRes, err error) {
	passwd := req.Password
	send := libmessage.GetEntity()
	err = send.CheckVerificationCode(req.Mobile, req.Code)
	if err != nil {
		return
	}
	// 注册 service
	if err = service.DeskUser().RegisterBind(ctx, req.Password, req.Mobile, req.Nickname, req.Nickname, ""); err != nil {
		return
	} else {
		// 自动登录
		var (
			user  *model.LoginUserRes
			token string
		)
		ip := libUtils.GetClientIp(ctx)
		userAgent := libUtils.GetUserAgent(ctx)
		user, err = service.DeskUser().GetUserByUsernamePassword(ctx, &desk.UserLoginReq{Mobile: req.Mobile, Password: passwd})
		if err != nil {
			// g.Log().Error(ctx, err)
			// err = gerror.New("账号或密码错误")
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
		return &desk.RegisterDoRes{fmt.Sprintf("register success login success"), user, token}, nil
	}
}
