package desk

import (
	"tender/internal/app/desk/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SendCodeReq struct {
	g.Meta `path:"/sendcode" method:"post" summary:"发送验证码" tags:"前台用户"`
	Mobile string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	// Code string `p:"code" v:"required|验证码不能为空"`
}
type SendCodeRes struct {
	Result string `json:"result" dc:"发送结果"`
	// Token    string              `json:"token"`
}

type RegisterDoReq struct {
	g.Meta `path:"/registerbind" method:"post" summary:"用户绑定注册" tags:"前台用户"`
	// Name     string `p:"name" v:"required"`
	Password string `p:"passWord" v:"required#密码不能为空"`
	Nickname string `p:"nickName"`
	// UserSalt string `p:"userSalt"`
	Mobile string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Code   string `p:"code" v:"required#验证码不能为空"`
}
type RegisterDoRes struct {
	Result   string              `json:"result" dc:"注册结果"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type BindWxReq struct {
	g.Meta   `path:"/bindwx" method:"post" summary:"用户绑定微信" tags:"前台用户"`
	Mobile   string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Code     string `p:"code" v:"required#验证码不能为空"`
	OpenId   string `p:"openId" v:"required#openId不能为空"`
	Password string `p:"passWord" v:"required#密码不能为空"`
	Nickname string `p:"nickName" v:"required#昵称不能为空"`
	Name     string `p:"name" v:"required#姓名不能为空"`
}

type BindWxRes struct {
	Result string `json:"result" dc:"绑定结果"`
}
