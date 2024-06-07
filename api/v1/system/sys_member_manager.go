package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/system/model/entity"
)

// MemberUserSearchReq 会员搜索请求参数
type MemberUserSearchReq struct {
	g.Meta  `path:"/member/list" tags:"会员管理" method:"get" summary:"会员列表"`
	Name    string `p:"name"`    // 企业名称
	Contact string `p:"contact"` // 联系方式
	Level   *int   `p:"level"`   // 会员等级
	Status  *int   `p:"status"`  // 用户状态
	common.PageReq
	common.Author
}

type MemberUserSearchRes struct {
	g.Meta         `mime:"application/json"`
	MemberUserList []*entity.MemberUser `json:"MemberUserList"`
	common.ListRes
}

type MemberUserEditReq struct {
	g.Meta   `path:"/member/edit" tags:"会员管理" method:"post" summary:"会员积分修改"`
	Id       int  `p:"id" v:"required#会员id不能为空"`
	Integral *int `p:"integral"` //积分
	Level    *int `p:"level" `   //等级
	common.Author
}

type MemberUserEditRes struct {
	g.Meta `mime:"application/json"`
}
type DisableMemberUserReq struct {
	g.Meta       `path:"/member/disable" tags:"会员管理" method:"post" summary:"会员禁用"`
	MemberUserId uint `p:"member_user_id"  v:"required#会员id不能为空"` //会员id
	Day          int  `p:"day"  v:"required#禁用天数不能为空"`            //禁用天数
	common.Author
}
type DisableMemberUserRes struct {
	Day int `p:"day"` //禁用天数
}

type MemberFee struct {
	MonthlycardOriginal string `p:"monthlycardOriginal"`
	MonthlycardCurrent  string `p:"monthlycardCurrent"`
	QuartercardOriginal string `p:"quartercardOriginal"`
	QuartercardCurrent  string `p:"quartercardCurrent"`
	AnnualcardOriginal  string `p:"annualcardOriginal"`
	AnnualcardCurrent   string `p:"annualcardCurrent"`
	DownloadKnowledge   string `p:"downloadKnowledge"`
	DownloadVideo       string `p:"downloadVideo"`
}

// 会费管理
type MemberFeeReq struct {
	g.Meta `path:"/member/fee/edit" tags:"会员管理" method:"post" summary:"修改会费"`
	//Id     string `p:"id" v:"required#修改id不能为空"`
	MemberFee
	common.Author
}

//type MemberFeeAddReq struct {
//	g.Meta `path:"/member/fee/add" tags:"会员管理" method:"post" summary:"设置会费"`
//	MemberFee
//	common.Author
//}
type MemberFeeRes struct {
}

type MemberIntegral struct {
	MonthlycardIntegral *int `p:"monthlycardIntegral"`
	QuartercardIntegral *int `p:"quartercardIntegral"`
	AnnualcardIntegral  *int `p:"annualcardIntegral"`
	KnowledgeIntegral   *int `p:"knowledgeIntegral"`
	VideoIntegral       *int `p:"videoIntegral"`
	// issue
	IssueIntegral    *int `p:"issueIntegral"`
	Ordinary         *int `p:"ordinary"`
	Select           *int `p:"select"`
	Ratio            *int `p:"ratio"`
	MonthlycardRatio *int `p:"monthlycardRatio"`
	QuartercardRatio *int `p:"quartercardRatio"`
	AnnualcardRatio  *int `p:"annualcardRatio"`
}
type MemberIntegralReq struct {
	g.Meta `path:"/member/integral/edit" tags:"会员管理" method:"post" summary:"修改积分"`
	// Id     string `p:"id" v:"required#修改id不能为空"`
	MemberIntegral
	common.Author
}

//type MemberIntegralAddReq struct {
//	g.Meta `path:"/member/integral/add" tags:"会员管理" method:"post" summary:"设置积分"`
//	MemberIntegral
//	common.Author
//}
type MemberIntegralRes struct {
}

type MemberSubscription struct {
	MonthlycardSubscription      *int   `p:"monthlycardSubscription"`
	QuartercardSubscription      *int   `p:"quartercardSubscription"`
	AnnualcardSubscription       *int   `p:"annualcardSubscription"`
	MonthlycardSubscriptionPrice string `p:"monthlycardSubscriptionPrice"`
	QuartercardSubscriptionPrice string `p:"quartercardSubscriptionPrice"`
	AnnualcardSubscriptionPrice  string `p:"annualcardSubscriptionPrice"`
}
type MemberSubscriptionReq struct {
	g.Meta `path:"/member/subscription/edit" tags:"会员管理" method:"post" summary:"修改订阅"`
	// Id     string `p:"id" v:"required#修改id不能为空"`
	MemberSubscription
	common.Author
}

//type MemberSubscriptionAddReq struct {
//	g.Meta `path:"/member/subscription/add" tags:"会员管理" method:"post" summary:"设置订阅"`
//	MemberSubscription
//	common.Author
//}
type MemberSubscriptionRes struct {
}

type MemberFeeFindReq struct {
	g.Meta `path:"/member/fee" tags:"会员管理" method:"get" summary:"会费设置查询"`
	common.Author
}
type MemberFeeFindRes struct {
	g.Meta            `mime:"application/json"`
	*entity.MemberFee `json:"member_fee"`
}

type MemberInFindReq struct {
	g.Meta `path:"/member/integral" tags:"会员管理" method:"get" summary:"积分设置查询"`
	common.Author
}
type MemberInFindRes struct {
	g.Meta                 `mime:"application/json"`
	*entity.MemberIntegral `json:"member_in"`
}

type MemberSuFindReq struct {
	g.Meta `path:"/member/subscription" tags:"会员管理" method:"get" summary:"订阅设置查询"`
	common.Author
}
type MemberSuFindRes struct {
	g.Meta                     `mime:"application/json"`
	*entity.MemberSubscription `json:"member_su"`
}
