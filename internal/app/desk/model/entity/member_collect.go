// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberCollect is the golang structure for table member_collect.
type MemberCollect struct {
	Id        uint64      `json:"id"        description:""`
	Title     string      `json:"title"     description:"标题"`
	Type      string      `json:"type"      description:"类型"`
	Location  string      `json:"location"  description:"所在地"`
	Industry  string      `json:"industry"  description:"行业"`
	UserId    int         `json:"userId"    description:"用户id"`
	ArticleId int         `json:"articleId" description:"文章id"`
	CreatedAt *gtime.Time `json:"createdAt" description:"收藏时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
	Url       string      `json:"url"       description:"收藏url"`
}
