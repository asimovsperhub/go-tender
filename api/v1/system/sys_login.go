package system

import (
	"tender/api/v1/common"
	dModel "tender/internal/app/desk/model"
	"tender/internal/app/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
}

type UserLoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	MenuList    []*model.UserMenus  `json:"menuList"`
	Permissions []string            `json:"permissions"`
	RoleIds     []uint              `json:"roleids"`
}

type UserLoginOutReq struct {
	g.Meta `path:"/logout" tags:"登录" method:"get" summary:"退出登录"`
	common.Author
}

type QrcodeReq struct {
	g.Meta `path:"/qrcode" tags:"登录" method:"get" summary:"扫码登陆"`
	common.Author
}

type UserLoginOutRes struct {
}

type QrcodeRes struct {
	g.Meta  `mime:"application/json"`
	Qrcode  string `json:"qrCode"`
	SenceId string `json:"senceId"`
}

type CheckScanQrcodeReq struct {
	g.Meta  `path:"/checkToken" tags:"登录" method:"get" summary:"扫码登陆轮询"`
	SenceId string `p:"senceId" v:"required#senceId不能为空"`
	Type    int    `p:"type"`
}

type CheckScanQrcodeRes struct {
	g.Meta      `mime:"application/json"`
	IsScan      bool                 `json:"isScan"`
	IsBind      bool                 `json:"isBind"`
	OpenId      string               `json:"openId"`
	UserInfo    *dModel.LoginUserRes `json:"userInfo"`
	Token       string               `json:"token"`
	MenuList    []*model.UserMenus   `json:"menuList"`
	Permissions []string             `json:"permissions"`
	RoleIds     []uint               `json:"roleids"`
}
