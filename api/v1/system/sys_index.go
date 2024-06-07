package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/system/model/entity"
)

//最新入驻企业
type EnterpriseReq struct {
	g.Meta `path:"/index/enterprise" tags:"后台首页" method:"get" summary:"最新入驻企业列表"`
	common.PageReq
	common.Author
}
type EnterpriseRes struct {
	g.Meta         `mime:"application/json"`
	EnterpriseList []*entity.SysEnterprise `json:"EnterpriseList"`
	common.ListRes
}

//最新入驻会员
type MemberUserReq struct {
	g.Meta `path:"/index/member" tags:"后台首页" method:"get" summary:"最新入驻会员列表"`
	common.PageReq
	common.Author
}

type MemberUserRes struct {
	g.Meta         `mime:"application/json"`
	MemberUserList []*entity.MemberUser `json:"MemberUserList"`
	common.ListRes
}

//总数统计
type UserNumberReq struct {
	g.Meta `path:"/index/number" tags:"后台首页" method:"get" summary:"用户总数-会员数-企业数"`
	common.Author
}

type UserNumberRes struct {
	g.Meta          `mime:"application/json"`
	UserCount       int `json:"userCount"`
	VipCount        int `json:"vipCount"`
	EnterpriseCount int `json:"enterpriseCount"`
}

//用户总数
type UserReq struct {
	g.Meta `path:"/index/usercount" tags:"后台首页" method:"get" summary:"用户总数"`
	common.Author
}

type UserRes struct {
	g.Meta `mime:"application/json"`
	Count  int `json:"count"`
}

// 销售额

// 会员走势
type TrendingReq struct {
	g.Meta `path:"/index/trending" tags:"后台首页" method:"get" summary:"会员走势图"`
	Start  string `p:"start"`
	End    string `p:"end"`
	Type   string `p:"type"`
	common.Author
}

type Trending struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}
type TrendingRes struct {
	g.Meta           `mime:"application/json"`
	TrendingItemList []*Trending `json:"trendingItemList"`
}
