package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"tender/api/v1/common"
)

type CreateRechargeOrderReq struct {
	g.Meta  `path:"/rechargeorder/create" tags:"积分购买" method:"post" summary:"创建订单"`
	Type    string  `p:"type"`
	PayType string  `p:"payType"`
	Memo    string  `p:"memo"`
	Amount  float64 `p:"amount"`
	common.Author
}
type CreateRechargeOrderRes struct {
	OrderId int64  `json:"id"`
	OrderNo string `json:"orderNo"`
}

type QueryRechargeOrderPaymentReq struct {
	g.Meta  `path:"/rechargeorder/payment" tags:"积分购买" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryRechargeOrderPaymentRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryRechargeOrderPaymentAlipayReq struct {
	g.Meta  `path:"/rechargeorder/payment/alipay" tags:"积分购买" method:"get" summary:"查询支付信息"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryRechargeOrderPaymentAlipayRes struct {
	CodeUrl string `json:"code_url"`
}

type QueryRechargeOrderInfoReq struct {
	g.Meta  `path:"/rechargeorder/query" tags:"积分购买" method:"get" summary:"查询订单状态"`
	OrderNo string `p:"orderNo"`
	common.Author
}

type QueryRechargeOrderInfoRes struct {
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
