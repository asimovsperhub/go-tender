package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model/entity"
)

/*
收藏
*/

// 收藏
type CollectAddReq struct {
	g.Meta   `path:"/collect/add" tags:"前台个人中心" method:"post" summary:"收藏"`
	Title    string `p:"title"`
	Type     string `p:"type" v:"required#类型不能为空"`
	Location string `p:"location"`
	Industry string `p:"industry"`
	// UserId    int    `p:"userId" v:"required#用户id不能为空"`
	ArticleId *int   `p:"articleId" v:"required#文章id不能为空"` //文章id
	Url       string `p:"url"`                             //文章url
	common.Author
}
type CollectAddRes struct {
}

// 删除收藏
type CollectDelReq struct {
	g.Meta    `path:"/collect/del" tags:"前台个人中心" method:"post" summary:"删除收藏"`
	UserId    int    `p:"userId" v:"required#用户id不能为空"`
	ArticleId int    `p:"articleId" v:"required#文章id不能为空"`
	Type      string `p:"type" v:"required#类型不能为空"`
	common.Author
}
type CollectDelRes struct {
	// Result         string               `json:"result" dc:"修改结果"`
}

// 我的收藏夹
type CollectGetListReq struct {
	g.Meta `path:"/collect/list" tags:"前台个人中心" method:"get" summary:"我的收藏夹"`
	UserId int    `p:"userId" v:"required#用户id不能为空"`
	Type   string `p:"type" `
	common.PageReq
	common.Author
}
type CollectGetListRes struct {
	MemberCollectList []*entity.MemberCollect `json:"MemberCollectList"`
	common.ListRes
}

// 用户是否收藏接口
type CollectGetReq struct {
	g.Meta    `path:"/collect/user/get" tags:"前台个人中心" method:"get" summary:"用户是否收藏"`
	UserId    int    `p:"userId" v:"required#用户id不能为空"`
	ArticleId int    `p:"articleId" v:"required#文章id不能为空"`
	Type      string `p:"type" v:"required#类型不能为空"`
	common.Author
}
type CollectGetRes struct {
	IsCollect bool `json:"isCollect" dc:"用户是否收藏"`
}

//	Id        uint64      `json:"id"        description:""`
//	Name      string      `json:"name"      description:"方案名称"`
//	Type      string      `json:"type"      description:"信息类型"`
//	Location  string      `json:"location"  description:"所在地"`
//	Keywords  string      `json:"keywords"  description:"关键字"`
//	UserId    int         `json:"userId"    description:"用户id"`

// 订阅
type SubscribeAddReq struct {
	g.Meta   `path:"/subscribe/add" tags:"订阅" method:"post" summary:"订阅"`
	Name     string `p:"title"`
	Type     string `p:"type"`
	Industry string `p:"industry"`
	Location string `p:"location"`
	Keywords string `p:"keywords"`
	UserId   int    `p:"userId"`
	common.Author
}
type SubscribeAddRes struct {
}

// 修改订阅
type SubscribeEditReq struct {
	g.Meta      `path:"/subscribe/edit" tags:"订阅" method:"post" summary:"修改订阅"`
	SubscribeId int    `p:"subscribeId" v:"required#订阅id不能为空"`
	Name        string `p:"title"`
	Type        string `p:"type"`
	Industry    string `p:"industry"`
	Location    string `p:"location"`
	Keywords    string `p:"keywords"`
	UserId      int    `p:"userId"`
	common.Author
}
type SubscribeEditRes struct {
}

// 删除订阅
type SubscribeDelReq struct {
	g.Meta      `path:"/subscribe/del" tags:"订阅" method:"post" summary:"删除订阅"`
	SubscribeId int `p:"subscribeId" v:"required#订阅id不能为空"`
	common.Author
}
type SubscribeDelRes struct {
	// Result         string               `json:"result" dc:"修改结果"`
}

// 我的订阅列表
type SubscribeListReq struct {
	g.Meta `path:"/subscribe/list" tags:"订阅" method:"get" summary:"我的订阅列表"`
	UserId int `p:"userId" v:"required#用户id不能为空"`
	common.PageReq
	common.Author
}
type SubscribeListRes struct {
	MemberSubscribeList []*entity.MemberSubscribe `json:"MemberSubscribeList"`
	common.ListRes
}

// 单个订阅
type SubscribeGetReq struct {
	g.Meta      `path:"/subscribe/get" tags:"订阅" method:"get" summary:"获取单个订阅"`
	SubscribeId int `p:"subscribeId" v:"required#用户id不能为空"`
	common.Author
}
type SubscribeGetRes struct {
	MemberSubscribe *entity.MemberSubscribe `json:"MemberSubscribe"`
}

// 当前用户剩余订阅数 can
type SubscribeCanReq struct {
	g.Meta `path:"/subscribe/can" tags:"订阅" method:"get" summary:"当前用户剩余订阅数/可订阅数"`
	common.Author
}
type SubscribeCanRes struct {
	MemberSubscribe *entity.MemberSubscribe `json:"MemberSubscribe"`
	Count           int                     `json:"count"`
}
