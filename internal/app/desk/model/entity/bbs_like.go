// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BbsLike is the golang structure for table bbs_like.
type BbsLike struct {
	Id        uint64      `json:"id"        description:""`
	BbsId     int64       `json:"bbsId"     description:"论坛id"`
	ReplyId   int64       `json:"replyId"   description:"回复id"`
	UserId    int64       `json:"userId"    description:"用户id"`
	CreatedAt *gtime.Time `json:"createdAt" description:"发布时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
