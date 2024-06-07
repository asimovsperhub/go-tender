package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model"
)

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"前台用户" method:"post" summary:"用户登录"`
	Mobile   string `p:"mobile" v:"required#手机号不能为空|手机号格式错误"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type UserLoginMessageReq struct {
	g.Meta `path:"/msglogin" tags:"前台用户" method:"post" summary:"短信登录"`
	Mobile string `p:"mobile" v:"required#手机号不能为空|手机号格式错误"`
	Code   string `p:"code" v:"required#验证码不能为空"`
}

type UserLoginMessageRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"前台用户" method:"get" summary:"退出登录"`
	common.Author
}

type UserLoginOutRes struct {
}

type UserForgetPassReq struct {
	g.Meta   `path:"/forgetPass" tags:"前台用户" method:"post" summary:"忘记密码"`
	Password string `p:"passWord" v:"required|password#密码不能为空|密码格式错误"`
	Mobile   string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Code     string `p:"code" v:"required#验证码不能为空"`
	common.Author
}
type UserForgetPassRes struct {
	Result   string              `json:"result" dc:"修改结果"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}
