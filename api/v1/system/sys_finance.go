package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/system/model/entity"
)

type FinanceReq struct {
	g.Meta `path:"/index/finance" tags:"后台财务中心" method:"get" summary:"购买明细"`
	Mobile string `p:"mobile"`
	common.PageReq
	common.Author
}

type FinanceRes struct {
	g.Meta          `mime:"application/json"`
	PurchaseLogList []*entity.PurchaseLog `json:"purchaseLogList"`
	common.ListRes
}

type PurchaseInfoReq struct {
	g.Meta `path:"/index/purchase" tags:"后台财务中心" method:"get" summary:"销售额总览"`
	common.Author
}

type PurchaseInfoRes struct {
	g.Meta          `mime:"application/json"`
	TodayAmount     float64 `json:"todayAmount"`
	ThisWeekAmount  float64 `json:"thisWeekAmount"`
	ThisMonthAmount float64 `json:"thisMonthAmount"`
}

type TrendingInfoReq struct {
	g.Meta `path:"/index/finance/trending" tags:"后台财务中心" method:"get" summary:"趋势"`
	Start  string `p:"start"`
	End    string `p:"end"`
	Type   string `p:"type"`
	common.Author
}

type TrendingItem struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
type TrendingInfoRes struct {
	g.Meta           `mime:"application/json"`
	TrendingItemList []*TrendingItem `json:"trendingItemList"`
	common.ListRes
}

type PaySettingsReq struct {
	g.Meta `path:"/index/finance/paySettings" tags:"后台财务中心" method:"get" summary:"查询支付设置"`
	common.Author
}

type PaySettingsRes struct {
	g.Meta   `mime:"application/json"`
	Settings *entity.PaySettings `json:"settings"`
}

type UpdatePaySettingsReq struct {
	g.Meta                 `path:"/index/finance/updatePaySettings" tags:"后台财务中心" method:"post" summary:"更新支付设置"`
	Id                     int     `json:"id"                     description:""`
	WeixinMchid            *string `json:"weixinMchid"            description:""`
	WeixinAppid            *string `json:"weixinAppid"            description:""`
	WeixinApikey           *string `json:"weixinApikey"           description:""`
	WeixinSerialno         *string `json:"weixinSerialno"         description:""`
	WeixinPrivatekey       *string `json:"weixinPrivatekey"       description:""`
	AlipayAppid            *string `json:"alipayAppid"            description:""`
	AlipayPrivatekey       *string `json:"alipayPrivatekey"       description:""`
	AlipayAppCertPublicKey *string `json:"alipayAppCertPublicKey" description:""`
	AlipayRootCert         *string `json:"alipayRootCert"         description:""`
	AlipayPublicCert       *string `json:"alipayPublicCert"       description:""`
	common.Author
}

type UpdatePaySettingsRes struct {
	g.Meta `mime:"application/json"`
}
