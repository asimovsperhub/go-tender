// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HisKnowledge is the golang structure of table his_knowledge for DAO operations like Where/Data.
type HisKnowledge struct {
	g.Meta      `orm:"table:his_knowledge, do:true"`
	Id          interface{} //
	KnowledgeId interface{} // 知识id
	UserId      interface{} // 用户id
	CreatedAt   *gtime.Time // 发布时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
