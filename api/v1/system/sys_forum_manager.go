package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model"
	"tender/internal/app/desk/model/entity"
	system_entity "tender/internal/app/system/model/entity"
)

// 论坛列表
type BbsGetListAllReq struct {
	g.Meta   `path:"/bbs/all/list" tags:"贴吧管理" method:"get" summary:"论坛列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}
type BbsList struct {
	Bbs *model.Bbs `json:"bbs"`
	// BbsContent *entity.BbsContent        `json:"bbsContent"`
	UserInfo *system_entity.MemberUser `json:"memberUser"`
}
type BbsGetListAllRes struct {
	BbsList []BbsList `json:"bbsList"`
	common.ListRes
}

// 单个论坛
type BbsGetReq struct {
	g.Meta `path:"/bbs/gets" tags:"贴吧管理" method:"get" summary:"获取单个论坛"`
	Id     int `p:"id" v:"required#论坛id不能为空"`
	common.Author
}

type BbsGetRes struct {
	Bbs *entity.Bbs `json:"bbs"`
	// BbsContent *entity.BbsContent `json:"bbsContent"`
}

// 删除论坛
type BbsDelReq struct {
	g.Meta `path:"/bbs/del" tags:"贴吧管理" method:"post" summary:"论坛删除"`
	Id     int `p:"id" v:"required#论坛id不能为空"`
	common.Author
}
type BbsDelRes struct {
}

// 恢复论坛
type BbsRestoreReq struct {
	g.Meta `path:"/bbs/restore" tags:"贴吧管理" method:"post" summary:"论坛恢复"`
	Id     int `p:"id" v:"required#论坛id不能为空"`
	common.Author
}
type BbsRestoreRes struct {
}

// 永久删除论坛
type BbsForeverReq struct {
	g.Meta `path:"/bbs/forever" tags:"贴吧管理" method:"post" summary:"论坛永久删除"`
	Id     int `p:"id" v:"required#论坛id不能为空"`
	common.Author
}

type BbsForeverRes struct {
}

// 论坛审查列表
type BbsGetReviewListReq struct {
	g.Meta   `path:"/bbs/review/list" tags:"贴吧管理" method:"get" summary:"论坛审查列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}
type BbsGetReviewListRes struct {
	BbsList []BbsList `json:"bbsList"`
	common.ListRes
}

// 论坛审批
type BbsReviewReq struct {
	g.Meta          `path:"/bbs/review" tags:"贴吧管理" method:"post" summary:"论坛审查"`
	Id              uint   `p:"id" v:"required#论坛id不能为空"`
	Status          int    `p:"status" v:"required#论坛审查状态不能为空"`
	OpReviewMessage string `p:"opreviewMessage" v:"max-length:400#字数超过400"`
	common.PageReq
	common.Author
}
type BbsReviewRes struct {
}

// 置顶论坛
type BbsTopReq struct {
	g.Meta `path:"/bbs/top" tags:"贴吧管理" method:"post" summary:"论坛置顶"`
	Ids    []string `p:"ids" v:"required#id列表不能为空"`
	Type   int      `p:"type" v:"required#类型不能为空"`
	common.Author
}
type BbsTopRes struct {
}
