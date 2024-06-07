package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/api/v1/common"
)

// 发布知识
type CreateVipOrderReq struct {
	g.Meta  `path:"/viporder/create" tags:"会员订阅" method:"post" summary:"创建订单"`
	Type    string `p:"type"`
	PayType string `p:"payType"`
	Memo    string `p:"memo"`
	common.Author
}
type CreateVipOrderRes struct {
	OrderId int64  `json:"id"`
	OrderNo string `json:"orderNo"`
}

type QueryVipOrderPaymentReq struct {
	g.Meta  `path:"/viporder/payment" tags:"会员订阅" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryVipOrderPaymentRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryVipOrderPaymentAlipayReq struct {
	g.Meta  `path:"/viporder/payment/alipay" tags:"会员订阅" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryVipOrderPaymentAlipayRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryVipOrderInfoReq struct {
	g.Meta  `path:"/viporder/query" tags:"会员订阅" method:"get" summary:"查询订单状态"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryVipOrderInfoRes struct {
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
	PayTime        *int        `json:"payTime"        description:""`
	OrderNo        string      `json:"orderNo"        description:""`
}
