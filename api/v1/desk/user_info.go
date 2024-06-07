package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model"
	"tender/internal/app/system/model/entity"
)

type UserReq struct {
	g.Meta `path:"/user/get" tags:"前台个人中心" method:"get" summary:"获取个人信息"`
	UserId string `p:"userId" v:"required#用户id不能为空"` // 用户id
	common.Author
}
type UserRes struct {
	// Result   string              `json:"result" dc:"修改结果"`
	UserInfo *entity.MemberUser `json:"userInfo"`
	// Token    string              `json:"token"`
	WxInfo *entity.SysWxUser `json:"wxInfo"`
}

type UserEditReq struct {
	g.Meta `path:"/user/edit" tags:"前台个人中心" method:"post" summary:"个人信息修改"`
	// Name            string `p:"name" v:"required#企业名称不能为空"`
	NickName *string `p:"userNickname"`                              // 昵称
	Email    *string `p:"userEmail" v:"email#邮箱格式不正确"`               // 邮箱
	Mobile   *string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"` // 手机号
	Avatar   []byte  `p:"avatar"`                                    // 头像
	Address  *string `p:"address"`                                   // 地址
	Describe *string `p:"describe" v:"max-length:200#字数超过200"`       // 个性签名
	common.Author
}
type UserEditRes struct {
	Result   string              `json:"result" dc:"修改结果"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	// Token    string              `json:"token"`
}

type UserPassEditReq struct {
	g.Meta   `path:"/user/resetpass" tags:"前台个人中心" method:"post" summary:"密码修改"`
	Password string `p:"passWord" v:"required|password#密码不能为空|密码格式错误"`
	Mobile   string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Code     string `p:"code" v:"required#验证码不能为空"`
	common.Author
}
type UserPassEditRes struct {
	Result   string              `json:"result" dc:"修改结果"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}
