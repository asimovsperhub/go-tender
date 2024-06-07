package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/desk/model"
	deskentity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/model/entity"
)

// 查企业
type EnterpriseSearchReq struct {
	g.Meta   `path:"/index/search/enterprise" tags:"首页" method:"get" summary:"首页企业搜索"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}

type EnterpriseSearchRes struct {
	g.Meta                `mime:"application/json"`
	EnterpriseInformation interface{} `json:"EnterpriseInformation"`
	common.ListRes
}

// 工商信息和主要人员
type EnterpriseCommerceReq struct {
	g.Meta `path:"/index/search/enterprise/commerce" tags:"首页-企业搜索" method:"get" summary:"首页企业搜索-详情-工商信息和主要人员"`
	Name   *string `p:"name"`
	common.Author
}

type EnterpriseCommerceRes struct {
	g.Meta   `mime:"application/json"`
	Commerce interface{} `json:"commerce"`
}

// 经营风险
type EnterprisePunishmentReq struct {
	g.Meta `path:"/index/search/enterprise/punishment" tags:"首页-企业搜索" method:"get" summary:"首页企业搜索-详情-经营风险"`
	Name   *string `p:"name"`
	common.PageReq
	common.Author
}

type EnterprisePunishmentRes struct {
	g.Meta     `mime:"application/json"`
	Punishment interface{} `json:"punishment"`
}

// 建筑资质
type EnterpriseQualificationReq struct {
	g.Meta `path:"/index/search/enterprise/qualification" tags:"首页-企业搜索" method:"get" summary:"首页企业搜索-详情-资质"`
	Name   *string `p:"name"`
	common.PageReq
	common.Author
}

type EnterpriseQualificationRes struct {
	g.Meta        `mime:"application/json"`
	Qualification interface{} `json:"qualification"`
}

//司法风险

type EnterpriseLawSuitReq struct {
	g.Meta `path:"/index/search/enterprise/lawsuit" tags:"首页-企业搜索" method:"get" summary:"首页企业搜索-详情-司法风险"`
	Name   *string `p:"name"`
	common.PageReq
	common.Author
}

type EnterpriseLawSuitRes struct {
	g.Meta  `mime:"application/json"`
	LawSuit interface{} `json:"lawSuit"`
}

//中标

type EnterpriseBiddingReq struct {
	g.Meta `path:"/index/search/enterprise/bidding" tags:"首页-企业搜索" method:"get" summary:"首页企业搜索-详情-中标"`
	Name   *string `p:"name"`
	common.PageReq
	common.Author
}

type EnterpriseBiddingRes struct {
	g.Meta  `mime:"application/json"`
	Bidding []*model.BiddingEnterprise `json:"bidding"`
	common.ListRes
}

// 查招标
type BidSearchReq struct {
	g.Meta   `path:"/index/search/bidding" tags:"首页" method:"get" summary:"首页招标搜索"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}

type BidSearchRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation []*deskentity.Bid `json:"biddingInformation"`
	common.ListRes
}

// 	最新公告
type AnnouncementListReq struct {
	g.Meta   `path:"/index/announcement" tags:"首页" method:"get" summary:"最新公告"`
	KeyWords string `p:"keyWords"`
	Type     string `p:"type"`
	Industry string `p:"industry"`
	City     string `p:"city"`
	common.PageReq
	common.Author
}

type AnnouncementListRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation []*deskentity.Bid `json:"announcement"`
	common.ListRes
}

// 单个招标数据
type AnnouncementReq struct {
	g.Meta `path:"/index/announcement/get" tags:"首页" method:"get" summary:"获取单个招标公告/单个自营咨询服务"`
	Id     string `p:"id"`
	common.Author
}

type AnnouncementRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation *deskentity.Bid `json:"announcement"`
}

// 咨询服务
type ConsulListReq struct {
	g.Meta   `path:"/index/consultation" tags:"首页" method:"get" summary:"咨询服务"`
	KeyWords string `p:"keyWords"`
	Type     string `p:"type"`
	Industry string `p:"industry"`
	City     string `p:"city"`
	common.PageReq
	common.Author
}

type ConsulListRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation []*model.Bidding `json:"consultation"`
	common.ListRes
}

// 后台数据中心的setting
type SettingGetReq struct {
	g.Meta `path:"/index/consultation/setting" tags:"首页" method:"get" summary:"咨询服务-后台数据中心setting"`
	common.Author
}

type SettingGetRes struct {
	g.Meta `mime:"application/json"`
	Result *entity.SysDataset `json:"result"`
}

// 单个自营咨询服务
//type ConsulSinReq struct {
//	g.Meta `path:"/index/consultation/get" tags:"首页" method:"get" summary:"获取单个自营咨询服务"`
//	Id     string `p:"id"`
//	common.Author
//}
//
//type ConsulSinRes struct {
//	g.Meta             `mime:"application/json"`
//	BiddingInformation *deskentity.Bid `json:"consultation"`
//}

// 	非自营 咨询服务
// industry
type ConsulReq struct {
	g.Meta   `path:"/index/consultation/enterprise" tags:"首页" method:"get" summary:"咨询服务-非自营-入驻企业信息"`
	KeyWords string `p:"keyWords"`
	Industry string `p:"industry"`
	City     string `p:"city"`
	common.PageReq
	common.Author
}

type ConsulRes struct {
	g.Meta                `mime:"application/json"`
	EnterpriseInformation []*entity.SysEnterprise `json:"enterprise"`
	common.ListRes
}

// 政策资讯
type ConsultationListReq struct {
	g.Meta         `path:"/index/information" tags:"首页" method:"get" summary:"政策资讯"`
	KeyWords       string `p:"keyWords"`
	Type           string `p:"type"`
	Classification string `p:"classification"`
	common.PageReq
	common.Author
}

type ConsultationListRes struct {
	g.Meta                  `mime:"application/json"`
	ConsultationInformation []*model.Consultation `json:"information"`
	common.ListRes
}

// 单个政策资讯
type ConsultationReq struct {
	g.Meta         `path:"/index/information/get" tags:"首页" method:"get" summary:"单个政策资讯"`
	Id             string `p:"id"`
	Classification string `p:"classification"`
	common.Author
}

type ConsultationRes struct {
	g.Meta                  `mime:"application/json"`
	ConsultationInformation *model.Consultation `json:"information"`
}

// 在线知识库
type KnowledgeListReq struct {
	g.Meta         `path:"/index/knowledge" tags:"首页" method:"get" summary:"在线知库"`
	KeyWords       string `p:"keyWords"`
	Type           string `p:"type"`
	ArticleType    string `p:"articleType"`
	Classification string `p:"classification"`
	common.PageReq
	common.Author
}

type Knowledge struct {
	Knowledge *deskentity.MemberKnowledge `json:"knowledge"`
	Count     int                         `json:"count"`
}
type KnowledgeListRes struct {
	g.Meta          `mime:"application/json"`
	MemberKnowledge []Knowledge `json:"memberKnowledge"`
	common.ListRes
}

type KnowledgeBrowseReq struct {
	g.Meta      `path:"/index/knowledge/browse" tags:"首页" method:"post" summary:"知识浏览量增加接口"`
	KnowledgeId int `p:"knowledgeId" v:"required#知识id不能为空"`
	common.Author
}
type KnowledgeBrowseRes struct {
}

type CollectCountGetReq struct {
	g.Meta `path:"/index/collect/count" tags:"首页" method:"get" summary:"文章收藏量"`
	// UserId    int    `p:"userId" v:"required#用户id不能为空"`
	ArticleId int    `p:"articleId" v:"required#文章id不能为空"`
	Type      string `p:"type" v:"required#类型不能为空"`
	common.Author
}
type CollectCountGetRes struct {
	Count int `json:"count" dc:"收藏数"`
}

// 统计接口

type StatisticsBrowseReq struct {
	g.Meta `path:"/index/statistics/browse" tags:"首页" method:"post" summary:"咨询量/下载量-增加接口（0咨询，1下载量）"`
	Type   int `p:"type" v:"required#类型不能为空"`
	common.Author
}
type StatisticsBrowseRes struct {
}

type StatisticsGetReq struct {
	g.Meta `path:"/index/statistics" tags:"首页" method:"get" summary:"统计"`
	common.Author
}
type StatisticsGetRes struct {
	HisTenderCount       int `json:"hisTender" dc:"历史招标数/招标总量"`
	TodayTenderCount     int `json:"todayTender" dc:"今日招标数"`
	HisConsultationCount int `json:"hisConsultation" dc:"历史咨询"`
	HisDownload          int `json:"hisDownload" dc:"历史下载"`
	TodayEnterprise      int `json:"todayEnterprise" dc:"今日企业"`
	Website              int `json:"website" dc:"网站数"`
	InformationCount     int `json:"information" dc:"资讯数量"`
}

// 文章积分查询

type MemberInFindReq struct {
	g.Meta `path:"/index/member/integral" tags:"首页" method:"get" summary:"积分设置查询"`
	common.Author
}
type MemberInFindRes struct {
	g.Meta                 `mime:"application/json"`
	*entity.MemberIntegral `json:"member_in"`
}

// 搜索
type SearchReq struct {
	g.Meta   `path:"/index/search" tags:"首页" method:"get" summary:"搜招标-咨询"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}
type SearchRes struct {
	g.Meta `mime:"application/json"`
	Result interface{} `json:"result"`
}
