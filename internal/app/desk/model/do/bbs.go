// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Bbs is the golang structure of table bbs for DAO operations like Where/Data.
type Bbs struct {
	g.Meta         `orm:"table:bbs, do:true"`
	Id             interface{} //
	Title          interface{} // 标题
	Abstract       interface{} // 摘要
	ReviewMessage  interface{} // 审核留言
	ReviewStatus   interface{} // 审核状态;0:待审核,1:通过，2:未通过
	Views          interface{} // 浏览量
	ReplyCount     interface{} // 回复量
	LikeCount      interface{} // 点赞量
	UserId         interface{} // 用户id
	CreatedAt      *gtime.Time // 发布时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	Classification interface{} // 所属分类
	Status         interface{} // 审核状态;0:软删除,1:正常，2:永久删除
	Rank           interface{} // 排名
	Content        interface{} // 内容
}
