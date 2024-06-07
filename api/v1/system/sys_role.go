package system

import (
	"tender/api/v1/common"
	"tender/internal/app/system/model"
	"tender/internal/app/system/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleListReq struct {
	g.Meta    `path:"/role/list" tags:"角色管理" method:"get" summary:"角色列表"`
	RoleName  string `p:"roleName"`   //参数名称
	Status    string `p:"roleStatus"` //状态
	RoleAlias string `p:"rolealias"`
	common.PageReq
}

type RoleListRes struct {
	g.Meta `mime:"application/json"`
	common.ListRes
	List []*entity.SysRole `json:"list"`
}

type RoleGetParamsReq struct {
	g.Meta `path:"/role/getParams" tags:"角色管理" method:"get" summary:"角色编辑参数"`
}

type RoleGetParamsRes struct {
	g.Meta `mime:"application/json"`
	Menu   []*model.SysAuthRuleInfoRes `json:"menu"`
}

type RoleAddReq struct {
	g.Meta    `path:"/role/add" tags:"角色管理" method:"post" summary:"添加角色"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    uint   `p:"status"    `
	ListOrder uint   `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
	ButtonIds []int  `p:"buttonIds"`
	RoleAlias string `p:"rolealias"`
}

type RoleAddRes struct {
}

type RoleGetReq struct {
	g.Meta `path:"/role/get" tags:"角色管理" method:"get" summary:"获取角色信息"`
	Id     uint `p:"id" v:"required#角色id不能为空"`
}

type RoleGetRes struct {
	g.Meta    `mime:"application/json"`
	Role      *entity.SysRole `json:"role"`
	MenuIds   []int           `json:"menuIds"`
	ButtonIds []int           `json:"buttonIds"`
}

type RoleEditReq struct {
	g.Meta    `path:"/role/edit" tags:"角色管理" method:"put" summary:"修改角色"`
	Id        int64  `p:"id" v:"required#角色id必须"`
	Name      string `p:"name" v:"required#角色名称不能为空"`
	Status    *uint  `p:"status"`
	ListOrder *uint  `p:"listOrder" `
	Remark    string `p:"remark"    `
	MenuIds   []uint `p:"menuIds"`
	ButtonIds []int  `p:"buttonIds"`
	RoleAlias string `p:"rolealias"`
}

type RoleEditRes struct {
}

type RoleDeleteReq struct {
	g.Meta `path:"/role/delete" tags:"角色管理" method:"delete" summary:"删除角色"`
	Ids    []int64 `p:"ids" v:"required#角色id不能为空"`
}

type RoleDeleteRes struct {
}

type SendTemplateReq struct {
	g.Meta     `path:"/msg/sendTemplate" tags:"消息管理" method:"post" summary:"模版消息"`
	TemplateId string            `p:"templateId" v:"required#模版id不能为空"`
	Params     map[string]string `p:"params" v:"required#模版参数不能为空"`
	ToUser     string            `p:"toUser" v:"required#用户openid不能为空"`
	URL        string            `p:"url"`
}
type SendTemplateRes struct {
	g.Meta `mime:"application/json"`
}

type SendWsReq struct {
	g.Meta  `path:"/msg/sendWs" tags:"消息管理" method:"post" summary:"websocket消息"`
	UserId  uint64 `p:"userId" v:"required#用户id不能为空"`
	Content string `p:"content" v:"required#消息内容不能为空"`
}

type SendWsRes struct {
	g.Meta `mime:"application/json"`
}

type WsMsgListReq struct {
	g.Meta `path:"/msg/wsMsgList" tags:"消息管理" method:"get" summary:"websocket消息列表"`
	common.PageReq
	UserId uint64 `p:"userId" v:"required#用户id不能为空"`
}

type WsMsgListRes struct {
	g.Meta    `mime:"application/json"`
	List      []entity.SysWsMsg `json:"list"`
	Total     int               `json:"total"`
	TotalPage int               `json:"totalPage"`
}
type MarkAsReadReq struct {
	g.Meta    `path:"/msg/markAsRead" tags:"消息管理" method:"post" summary:"已读消息"`
	MessageId string `p:"messageId" v:"required#消息id不能为空"`
}

type MarkAsReadRes struct {
	g.Meta `mime:"application/json"`
}

type MarkAsAllReadReq struct {
	g.Meta `path:"/msg/markAsReadAll" tags:"消息管理" method:"post" summary:"已读全部消息"`
	UserId uint64 `p:"userId" v:"required#用户id不能为空"`
}

type MarkAsAllReadRes struct {
	g.Meta `mime:"application/json"`
}
type DeleteMsgReq struct {
	g.Meta    `path:"/msg/deleteMsg" tags:"消息管理" method:"post" summary:"删除消息"`
	MessageId string `p:"messageId" v:"required#消息id不能为空"`
}
type DeleteMsgRes struct {
	g.Meta `mime:"application/json"`
}

type ClearMsgReq struct {
	g.Meta `path:"/msg/clearMsg" tags:"消息管理" method:"post" summary:"清空消息"`
	UserId uint64 `p:"userId" v:"required#用户id不能为空"`
}

type ClearMsgRes struct {
	g.Meta `mime:"application/json"`
}

type UnreadCountReq struct {
	g.Meta `path:"/msg/unreadCount" tags:"消息管理" method:"get" summary:"未读消息数量"`
	UserId uint64 `p:"userId" v:"required#用户id不能为空"`
}

type UnreadCountRes struct {
	g.Meta `mime:"application/json"`
	Count  int `json:"count"`
}
