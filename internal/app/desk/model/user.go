package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"tender/internal/app/system/model/entity"
)

// UserLoginInput 用户登录
type UserLoginInput struct {
	Name     string // 账号
	Password string // 密码(明文)
}

// UserRegisterInput 用户注册
//UserName:     req.Name,
//		UserPassword: req.Password,
//		UserNickname: req.Nickename,
type UserRegisterInput struct {
	UserName     string // 账号
	UserPassword string //昵称
	UserNickname string // 密码(明文)
}

// LoginUserRes 登录返回
type LoginUserRes struct {
	Id           uint64                `orm:"id,primary"       json:"id"`           //
	UserName     string                `orm:"user_name,unique" json:"userName"`     // 用户名
	UserNickname string                `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string                `orm:"user_password"    json:"userPassword"` // 登录密码;cmf_password加密
	UserSalt     string                `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint                  `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	Avatar       string                `orm:"avatar" json:"avatar"`                 //头像
	Mobile       string                `orm:"mobile" json:"mobile"`                 //手机号
	Describe     string                `orm:"describe" json:"describe"`             //个性签名
	Address      string                `orm:"address" json:"address"`               //地址
	MemberLevel  int                   `orm:"member_level" json:"memberLevel"`      //会员等级
	Integral     int                   `orm:"integral" json:"integral"`             //积分
	ReleaseAt    *gtime.Time           `orm:"release_at" json:"releaseAt"`          //积分
	Enterprise   *entity.SysEnterprise `json:"enterprise"`
}

// 企业信息
type EnterpriseRes struct {
	Id           uint64 `orm:"id,primary"       json:"id"`
	Name         string `orm:"name,unique" json:"name"`          // 企业名称
	Location     string `orm:"location" json:"location"`         // 所在地
	Industry     string `orm:"industry" json:"industry"`         // 所属行业
	Contact      string `orm:"contact" json:"contact"`           // 联系电话
	Icon         string `orm:"icon" json:"icon"`                 // 图标
	Introduction string `orm:"introduction" json:"introduction"` // 简介
	//Remark             interface{} // 其他
	//License            interface{} // 营业执照
	//LicenseStatus      interface{} // 审核状态 0待审核 1审核通过 2审核未通过
	//Certificate        interface{} // 证明书
	//CertificateStatus  interface{} // 审核状态 0待审核 1审核通过 2审核未通过
	//OperationId        interface{} // 操作管理员id
	EstablishmentAt string `orm:"establishment_at" json:"establishmentAt"` // 成立时间
	//CreatedAt          *gtime.Time // 创建日期
	//UpdatedAt          *gtime.Time // 修改日期
	//LicenseMessage     interface{} // 营业执照审核留言
	//CertificateMessage interface{} // 证明书审核留言
	UserId int `orm:"user_id" json:"user_id"` // 用户id
}
