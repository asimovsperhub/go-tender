// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberUser is the golang structure for table member_user.
type MemberUser struct {
	Id            uint64      `json:"id"            description:""`
	UserName      string      `json:"userName"      description:"用户名"`
	Mobile        string      `json:"mobile"        description:"中国手机不带国家代码，国际手机号格式为：国家代码-手机号"`
	UserNickname  string      `json:"userNickname"  description:"用户昵称"`
	Birthday      int         `json:"birthday"      description:"生日"`
	UserPassword  string      `json:"userPassword"  description:"登录密码;cmf_password加密"`
	UserSalt      string      `json:"userSalt"      description:"加密盐"`
	UserStatus    uint        `json:"userStatus"    description:"用户状态;0:禁用,1:正常,2:未验证"`
	UserEmail     string      `json:"userEmail"     description:"用户登录邮箱"`
	Sex           int         `json:"sex"           description:"性别;0:保密,1:男,2:女"`
	Avatar        string      `json:"avatar"        description:"用户头像"`
	DeptId        uint64      `json:"deptId"        description:"部门id"`
	Remark        string      `json:"remark"        description:"备注"`
	MemberLevel   int         `json:"memberLevel"   description:"1月卡 2季卡 3 年卡   0 否"`
	Integral      int         `json:"integral"      description:"积分"`
	Address       string      `json:"address"       description:"联系地址"`
	Describe      string      `json:"describe"      description:"描述信息"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录时间"`
	OperationId   int64       `json:"operationId"   description:"操作管理员id"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	DisabledAt    *gtime.Time `json:"disabledAt"    description:"禁用时间"`
	ReleaseAt     *gtime.Time `json:"releaseAt"     description:"解禁时间"`
	MaturityAt    *gtime.Time `json:"maturityAt"    description:"到期时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	DisabledDay   uint        `json:"disabledDay"   description:"禁用天数"`
	Subscribe     int         `json:"subscribe"     description:"可订阅数"`
	BuySubscribe  int         `json:"buySubscribe"  description:"购买的订阅数"`
	Subscribed    int         `json:"subscribed"    description:"已订阅数"`
}
