package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/api/v1/common"
)

type CreateKnowledgeOrderReq struct {
	g.Meta      `path:"/knowledgeorder/create" tags:"知识库购买" method:"post" summary:"创建订单"`
	PayType     string `p:"payType"`
	Memo        string `p:"memo"`
	KnowledgeId int64  `p:"knowledgeId"`
	common.Author
}
type CreateKnowledgeOrderRes struct {
	OrderId int64  `json:"id"`
	OrderNo string `json:"orderNo"`
}

type QueryKnowledgeOrderPaymentReq struct {
	g.Meta  `path:"/knowledgeorder/payment" tags:"知识库购买" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryKnowledgeOrderPaymentRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryKnowledgeOrderPaymentAlipayReq struct {
	g.Meta  `path:"/knowledgeorder/payment/alipay" tags:"知识库购买" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryKnowledgeOrderPaymentAlipayRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryKnowledgeOrderInfoReq struct {
	g.Meta  `path:"/knowledgeorder/query" tags:"知识库购买" method:"get" summary:"查询订单状态"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryKnowledgeOrderInfoRes struct {
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
