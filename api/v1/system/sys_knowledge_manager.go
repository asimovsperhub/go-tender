package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model/entity"
)

// PrimaryClassification   string      `json:"primaryClassification"   description:"一级分类"`
// SecondaryClassification string      `json:"secondaryClassification" description:"二级分类"`

type KnowledgeSearchReq struct {
	g.Meta                  `path:"/knowledge/list" tags:"知库管理" method:"get" summary:"知库列表"`
	Title                   string `p:"title"`
	PrimaryClassification   string `p:"primaryClassification"`
	SecondaryClassification string `p:"secondaryClassification"`
	// 0没有附件，1有附件
	Attachment *int `p:"attachment"`
	// 原始行业类型
	OriginalIndustry string `p:"originalIndustry"`
	// 知识类型  普通/精选
	KnowledgeType string `p:"knowledgeType"`
	Authority     *int   `p:"authority"`
	Display       *int   `p:"display"`
	// 站点
	Site string `p:"site"`
	common.PageReq
	common.Author
}
type KnowledgeSearchRes struct {
	g.Meta        `mime:"application/json"`
	KnowledgeList []*entity.MemberKnowledge `json:"knowledgeList"`
	common.ListRes
}

// 知识导入列表
type KnowledgeReq struct {
	g.Meta   `path:"/knowledge/import/list" tags:"知库管理" method:"get" summary:"知库导入列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}
type KnowledgeRes struct {
	g.Meta        `mime:"application/json"`
	KnowledgeList []*entity.MemberKnowledge `json:"knowledgeList"`
	common.ListRes
}
type KnowledgeReviewSearchReq struct {
	g.Meta                  `path:"/knowledge/review/list" tags:"知库管理" method:"get" summary:"知库审查列表"`
	KeyWords                string `p:"keyWords"`
	PrimaryClassification   string `p:"primaryClassification"`
	SecondaryClassification string `p:"secondaryClassification"`
	common.PageReq
	common.Author
}

// 删除知库
type KnowledgeDelReq struct {
	g.Meta `path:"/knowledge/delete" tags:"知库管理" method:"post" summary:"知库删除"`
	Id     uint `p:"id" v:"required#知识id不能为空"`
	// Status int  `p:"status" v:"required#知库审查状态不能为空"`
	common.Author
}
type KnowledgeDelRes struct {
}
type KnowledgeReviewSearchRes struct {
	g.Meta        `mime:"application/json"`
	KnowledgeList []*entity.MemberKnowledge `json:"knowledgeList"`
	common.ListRes
}

type KnowledgeReviewReq struct {
	g.Meta          `path:"/knowledge/review" tags:"知库管理" method:"post" summary:"知库审查"`
	Id              uint   `p:"id" v:"required#知识id不能为空"`
	Status          int    `p:"status" v:"required#知库审查状态不能为空"`
	OpReviewMessage string `p:"opreviewMessage" v:"max-length:400#字数超过400"`
	common.Author
}
type KnowledgeReviewRes struct {
}

type KnowledgeProcessVideoReq struct {
	g.Meta `path:"/knowledge/processVideo" tags:"知库管理" method:"post" summary:"知库视频处理"`
	common.Author
}

type KnowledgeProcessVideoRes struct {
}

type KnowledgeProcessPdfReq struct {
	g.Meta `path:"/knowledge/processPdf" tags:"知库管理" method:"post" summary:"知库pdf处理"`
	common.Author
}

type KnowledgeProcessPdfRes struct {
}
