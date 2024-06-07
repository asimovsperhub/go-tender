// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// KnowledgeOrder is the golang structure for table knowledge_order.
type KnowledgeOrder struct {
	Id             int         `json:"id"             description:""`
	Type           string      `json:"type"           description:""`
	UserId         int         `json:"userId"         description:""`
	OriginalAmount int         `json:"originalAmount" description:""`
	PayAmount      int         `json:"payAmount"      description:""`
	Status         string      `json:"status"         description:""`
	PayType        string      `json:"payType"        description:""`
	Memo           string      `json:"memo"           description:""`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:""`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:""`
	PayTime        int         `json:"payTime"        description:""`
	OrderNo        string      `json:"orderNo"        description:""`
	KnowledgeId    int         `json:"knowledgeId"    description:""`
}
