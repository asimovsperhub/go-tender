package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	deskmodel "tender/internal/app/desk/model"
	deskentity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/model/entity"
)

type DataSearchReq struct {
	g.Meta `path:"/data/tender/list" tags:"数据中心" method:"get" summary:"数据列表-招标信息"`
	// 0没有附件，1有附件
	Attachment *int `p:"attachment"`
	// 原始公告类型
	OriginalBulletin string `p:"originalBulletin"`
	// 原始行业类型
	OriginalIndustry string `p:"originalIndustry"`
	// 站点
	Site         string `p:"site"`
	BulletinType string `p:"bulletinType"`
	IndustryType string `p:"industryType"`
	City         string `p:"city"`
	KeyWords     string `p:"keyWords" `
	common.PageReq
	common.Author
}

type DataSearchRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation []*deskentity.Bid `json:"biddingInformation"`
	Statistics         StatisticsRes     `p:"statistics" `
	common.ListRes
}

// 单个招标数据
type TenderGetReq struct {
	g.Meta `path:"/data/tender/get" tags:"数据中心" method:"get" summary:"数据列表-招标信息-单个招标信息"`
	Id     string `p:"id" v:"required#招标id不能为空"`
	common.Author
}

type TenderGetRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation *deskentity.Bid `json:"biddingInformation"`
}

// 编辑单个招标数据
type TenderEditReq struct {
	g.Meta       `path:"/data/tender/edit" tags:"数据中心" method:"post" summary:"数据列表-招标信息-修改单个招标信息"`
	Id           string `p:"id" v:"required#招标id不能为空"`
	BulletinType string `p:"bulletinType"`
	IndustryType string `p:"industryType"`
	common.Author
}

type TenderEditRes struct {
	g.Meta `mime:"application/json"`
}

// 招标数据原始类型
type TenderTypeGetReq struct {
	g.Meta `path:"/data/tender/type/get" tags:"数据中心" method:"get" summary:"数据列表-招标信息-一二级原始类型"`
}

type TenderTypeGetRes struct {
	g.Meta                 `mime:"application/json"`
	OriginalType           []string `p:"originalType" `
	OriginalClassification []string `p:"originalClassification" `
}

// 添加自定义排序
type TenderAddSortReq struct {
	g.Meta `path:"/data/tender/add/sort" tags:"自定义排序" method:"post" summary:"添加排序-招标数据"`
	Ids    []string `p:"ids" v:"required#id列表不能为空"`
	Type   int      `p:"type" v:"required#类型不能为空"`
	Height int      `p:"height" v:"required#高度不能为空"`
}

type TenderAddSortRes struct {
	g.Meta `mime:"application/json"`
}

// 	移除排序
type TenderDelSortReq struct {
	g.Meta `path:"/data/tender/del/sort" tags:"自定义排序" method:"post" summary:"移除排序-招标数据"`
	Type   int      `p:"type" v:"required#类型不能为空"`
	Ids    []string `p:"ids" v:"required#招标id列表不能为空"`
}

type TenderDelSortRes struct {
	g.Meta `mime:"application/json"`
}

// 拖拽 drag
type TenderDragReq struct {
	g.Meta `path:"/data/tender/drag" tags:"自定义排序" method:"post" summary:"拖拽-招标数据"`
	Id     string `p:"id" v:"required#招标id列表不能为空"`
	//Prev   string `p:"prev"`
	//Next   string `p:"next"`
	Type      int    `p:"type" v:"required#类型不能为空"`
	Rank      string `p:"rank" v:"required#当前rank不能为空"`
	Direction int    `p:"direction" v:"required#方向不能为空"`
}

type TenderDragRes struct {
	g.Meta `mime:"application/json"`
}

//删除招标数据
type TenderDelReq struct {
	g.Meta `path:"/data/tender/del" tags:"数据中心" method:"post" summary:"数据列表-招标信息-剪除招标信息"`
	Id     string `p:"id" v:"required#招标id不能为空"`
	common.Author
}

type TenderDelRes struct {
	g.Meta `mime:"application/json"`
}

type LawSearchReq struct {
	g.Meta   `path:"/data/lawsearch" tags:"数据中心" method:"get" summary:"法律数据列表"`
	KeyWords string `p:"keyWords" `
	common.PageReq
	common.Author
}

type LawSearchRes struct {
	g.Meta `mime:"application/json"`
	Law    []*deskmodel.Consultation `json:"Law"`
	common.ListRes
}

type AllLawSearchReq struct {
	g.Meta   `path:"/data/lawall" tags:"数据中心" method:"get" summary:"所有法律数据列表"`
	KeyWords string `p:"keyWords" `
	common.PageReq
	common.Author
}

type AllLawSearchRes struct {
	g.Meta `mime:"application/json"`
	Law    []*entity.Law `json:"Law"`
	common.ListRes
}

type LawAddReq struct {
	g.Meta `path:"/data/lawadd" tags:"数据中心" method:"post" summary:"添加经常使用法律"`
	Id     int `p:"id" v:"required#法律id不能为空"`
	common.Author
}

type LawAddRes struct {
	g.Meta `mime:"application/json"`
}

type LawDelReq struct {
	g.Meta `path:"/data/lawdel" tags:"数据中心" method:"post" summary:"删除经常使用法律"`
	Id     int `p:"id" v:"required#法律id不能为空"`
	common.Author
}

type LawDelRes struct {
	g.Meta `mime:"application/json"`
}

//	Id           string //
//	Consultation string // 咨询服务非营利类目
//	Purchase     string // 采购咨询分类
//	Bid          string // 投标咨询分类
//	Industry     string // 行业咨询分类
//	Market       string // 市场研究分类
//	OperationId  string // 操作管理员id
//	CreatedAt    string // 创建日期
//	UpdatedAt    string // 修改日期
type SettingReq struct {
	g.Meta          `path:"/data/setting" tags:"数据中心" method:"post" summary:"其他设置"`
	Consultation    *string `p:"consultation"`    //咨询服务非营利类目
	DelConsultation *string `p:"delConsultation"` //要删除的 咨询服务非营利类目
	Bid             *string `p:"bid"`             //投标咨询分类
	Purchase        *string `p:"purchase"`        //采购咨询
	PurchaseType    *string `p:"purchaseType"`    //采购咨询分类
	Industry        *string `p:"industry"`        //行业咨询
	IndustryType    *string `p:"industryType"`    //行业咨询分类
	Market          *string `p:"market"`          //市场调研
	MarketType      *string `p:"marketType"`      //市场调研分类
	KeyWords        *string `p:"keyWords"`        //关键词
	common.Author
}

type SettingRes struct {
	g.Meta `mime:"application/json"`
}

type SettingGetReq struct {
	g.Meta `path:"/data/setting" tags:"数据中心" method:"get" summary:"其他设置"`
	common.Author
}

type SettingGetRes struct {
	g.Meta `mime:"application/json"`
	Result *entity.SysDataset `json:"result"`
}

// 企业信息
type EnterpriseListReq struct {
	g.Meta   `path:"/data/enterprise/list" tags:"数据中心" method:"get" summary:"数据列表-企业信息"`
	KeyWords string `p:"keyWords" `
	common.PageReq
	common.Author
}
type EnterpriseListRes struct {
	g.Meta                `mime:"application/json"`
	EnterpriseInformation []*entity.SysEnterprise `json:"enterprise"`
	Statistics            StatisticsRes           `p:"statistics" `
	common.ListRes
}

type EnterpriseGetReq struct {
	g.Meta `path:"/data/enterprise/get" tags:"数据中心" method:"get" summary:"数据列表-企业信息-单个企业信息"`
	Id     string `p:"id" v:"required#企业id不能为空"`
	common.Author
}
type EnterpriseGetRes struct {
	g.Meta                `mime:"application/json"`
	EnterpriseInformation *entity.SysEnterprise `json:"enterprise"`
}

//EnterpriseDel

type EnterpriseDelReq struct {
	g.Meta `path:"/data/enterprise/del" tags:"数据中心" method:"post" summary:"数据列表-企业信息-剪除单个企业信息"`
	Id     string `p:"id" v:"required#企业id不能为空"`
	common.Author
}
type EnterpriseDelRes struct {
	g.Meta `mime:"application/json"`
}

//KnowledgeList
type KnowledgeListReq struct {
	g.Meta    `path:"/data/knowledge/list" tags:"数据中心" method:"get" summary:"数据列表-在线知库"`
	Primary   string `p:"primary"`
	Secondary string `p:"secondary"`
	KeyWords  string `p:"keyWords" `
	// 0没有附件，1有附件
	Attachment *int `p:"attachment"`
	//// 原始公告类型
	//OriginalBulletin string `p:"originalBulletin"`
	// 原始行业类型
	OriginalIndustry string `p:"originalIndustry"`
	// 站点
	Site string `p:"site"`
	common.PageReq
	common.Author
}
type KnowledgeListRes struct {
	g.Meta          `mime:"application/json"`
	MemberKnowledge []*deskentity.MemberKnowledge `json:"knowledge"`
	Statistics      StatisticsRes                 `p:"statistics" `
	common.ListRes
}

type KnowledgeGetReq struct {
	g.Meta `path:"/data/knowledge/get" tags:"数据中心" method:"get" summary:"数据列表-在线知库-单个知识/删除用知库管理里的删除"`
	Id     string `p:"id" v:"required#知识id不能为空"`
	common.Author
}
type KnowledgeGetRes struct {
	g.Meta          `mime:"application/json"`
	MemberKnowledge *deskentity.MemberKnowledge `json:"knowledge"`
}

// 行业标准原始类型
type KnowledgeTypeGetReq struct {
	g.Meta `path:"/data/knowledge/type/get" tags:"数据中心" method:"get" summary:"在线知库-行业标准-二级原始类型"`
}

type KnowledgeTypeGetRes struct {
	g.Meta       `mime:"application/json"`
	OriginalType []string `p:"originalType" `
}

//政策资讯
type InformationListReq struct {
	g.Meta         `path:"/data/information/list" tags:"数据中心" method:"get" summary:"数据列表-政策资讯"`
	Classification string `p:"classification"`
	KeyWords       string `p:"keyWords" `
	// 0没有附件，1有附件
	Attachment *int `p:"attachment"`
	// 原始公告类型
	OriginalBulletin string `p:"originalBulletin"`
	// 原始行业类型
	OriginalIndustry string `p:"originalIndustry"`
	// 站点
	Site string `p:"site"`
	common.PageReq
	common.Author
}

type InformationListRes struct {
	g.Meta                  `mime:"application/json"`
	ConsultationInformation []*deskmodel.Consultation `json:"information"`
	Statistics              StatisticsRes             `p:"statistics" `
	common.ListRes
}
type InformationGetReq struct {
	g.Meta         `path:"/data/information/get" tags:"数据中心" method:"get" summary:"数据列表-政策资讯-单个政策资讯"`
	Id             string `p:"id" v:"required#政策资讯id不能为空"`
	Classification string `p:"classification"`
	common.Author
}
type InformationGetRes struct {
	g.Meta      `mime:"application/json"`
	Information *deskmodel.Consultation `json:"information"`
}

type InformationEditReq struct {
	g.Meta `path:"/data/information/edit" tags:"数据中心" method:"post" summary:"数据列表-政策资讯-修改单个政策资讯的咨询类型"`
	Id     string `p:"id" v:"required#政策资讯id不能为空"`
	Type   string `p:"type"`
	common.Author
}
type InformationEditRes struct {
	g.Meta `mime:"application/json"`
}

type InformationDisplayReq struct {
	g.Meta `path:"/data/information/display" tags:"政策咨询在前端显示数据中心" method:"post" summary:"数据列表-政策咨询在前端显示"`
	Ids    []string `p:"ids" v:"required#政策资讯id不能为空"`
	Switch int      `p:"switch" v:"required#开关不能为空"`
	common.Author
}
type InformationDisplayRes struct {
	g.Meta `mime:"application/json"`
}

type InformationDelReq struct {
	g.Meta         `path:"/data/information/del" tags:"数据中心" method:"post" summary:"数据列表-政策资讯-剪除单个政策资讯"`
	Id             string `p:"id" v:"required#政策资讯id不能为空"`
	Classification string `p:"classification"`
	common.Author
}
type InformationDelRes struct {
	g.Meta `mime:"application/json"`
}

// search city statistics
type StatisticsReq struct {
	g.Meta `path:"/data/statistics" tags:"数据中心" method:"get" summary:"数据列表-搜索-城市统计"`
	// 0没有附件，1有附件
	Attachment *int `p:"attachment"`
	// 原始公告类型
	OriginalBulletin string `p:"originalBulletin"`
	// 原始行业类型
	OriginalIndustry string `p:"originalIndustry"`
	// 站点
	Site           string `p:"site"`
	BulletinType   string `p:"bulletinType"`
	IndustryType   string `p:"industryType"`
	City           string `p:"city"`
	KeyWords       string `p:"keyWords"`
	Classification string `p:"classification"`
	Primary        string `p:"primary"`
	Secondary      string `p:"secondary"`
	Type           int    `p:"type" v:"required#类型不能为空"`
	common.Author
}
type StatisticsRes struct {
	g.Meta `mime:"application/json"`
	City   map[string]int64 `json:"city"`
	Plate  map[int]int      `json:"plate"`
}

// 数据中心-数据是否隐藏
type KnowledgeDisplayReq struct {
	g.Meta `path:"/data/knowledge/display" tags:"数据中心-数据是否隐藏" method:"post" summary:"数据是否隐藏"`
	Ids    []string `p:"ids" v:"required#知识id不能为空"`
	Switch int      `p:"switch" v:"required#开关不能为空"`
	Type   int      `p:"type" v:"required#类型不能为空"`
	common.Author
}
type KnowledgeDisplayRes struct {
	g.Meta `mime:"application/json"`
}

// -----------------------------------------调研公告------------------------------------------------------

// 发布调研公告
type ReleaseReq struct {
	g.Meta       `path:"/research/add" tags:"公告管理" method:"post" summary:"发布调研公告"`
	Title        string `p:"title" v:"max-length:80#字数超过80"` //标题
	Content      string `p:"content"`                        // 正文
	Attachment   string `p:"attachment" `                    //附件
	IndustryType string `p:"industryType"`                   //行业类型
	ReleaseTime  string `p:"releaseTime" `                   //发布时间
	common.Author
}
type ReleaseRes struct {
	g.Meta `mime:"application/json"`
}

// 编辑调研公告
type ReleaseEditReq struct {
	g.Meta       `path:"/research/edit" tags:"公告管理" method:"post" summary:"编辑调研公告"`
	Id           int     `p:"id" v:"required#公告id不能为空"`
	Title        *string `p:"title" v:"max-length:80#字数超过80"` //标题
	Content      *string `p:"content"`                        // 正文
	Attachment   *string `p:"attachment" `                    //附件
	IndustryType *string `p:"industryType"`                   //行业类型
	ReleaseTime  *string `p:"releaseTime" `                   //发布时间
	common.Author
}
type ReleaseEditRes struct {
	g.Meta `mime:"application/json"`
}

type ReleaseListReq struct {
	g.Meta        `path:"/research/list" tags:"公告管理" method:"get" summary:"发布调研列表"`
	Title         string `p:"title"`         //标题
	ContactPerson string `p:"contactPerson"` // 正文
	common.Author
	common.PageReq
}
type LReleaseList struct {
	BiddingInformation *deskentity.Bid `json:"biddingInformation"`
	Count              int             `p:"count"` // 正文
}
type ReleaseListRes struct {
	g.Meta        `mime:"application/json"`
	LResearchList []LReleaseList `json:"researchList"`
	common.ListRes
}

type ReleaseGetReq struct {
	g.Meta `path:"/research/get" tags:"公告管理" method:"get" summary:"单个调研公告"`
	Id     int `p:"id" v:"required#公告id不能为空"`
	common.Author
}
type ReleaseGetRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation *deskentity.Bid `json:"biddingInformation"`
}

// 删除调研公告
type ReleaseDelReq struct {
	g.Meta `path:"/research/del" tags:"公告管理" method:"post" summary:"删除调研公告"`
	Id     int `p:"id" v:"required#公告id不能为空"`
	common.Author
}
type ReleaseDelRes struct {
	g.Meta `mime:"application/json"`
}

//---------------------------------------------------调研公告反馈页-----------------------------------------------------

type FeedbackListReq struct {
	g.Meta     `path:"/feedback/list" tags:"公告管理" method:"get" summary:"反馈页数据"`
	BulletinId int `p:"bulletinId" v:"required#公告id不能为空"`
	common.Author
	common.PageReq
}
type FeedbackListRes struct {
	g.Meta       `mime:"application/json"`
	FeedbackList []*deskentity.Feedback `json:"feedbackList"`
	common.ListRes
}

type FeedbackGetReq struct {
	g.Meta `path:"/feedback/get" tags:"公告管理" method:"get" summary:"反馈页查看"`
	Id     int `p:"id" v:"required#反馈id不能为空"`
	common.Author
}
type FeedbackGetRes struct {
	g.Meta   `mime:"application/json"`
	Feedback *deskentity.Feedback `json:"feedback"`
}

type FeedbackDelReq struct {
	g.Meta `path:"/feedback/del" tags:"公告管理" method:"post" summary:"反馈删除"`
	Id     int `p:"id" v:"required#反馈id不能为空"`
	common.Author
}
type FeedbackDelRes struct {
	g.Meta `mime:"application/json"`
}

type FeedbackReviewReq struct {
	g.Meta        `path:"/feedback/review" tags:"公告管理" method:"post" summary:"反馈审查"`
	Id            int    `p:"id" v:"required#反馈id不能为空"`
	Status        int    `p:"status" v:"required#审查状态不能为空"`
	ReviewMessage string `p:"reviewMessage"`
	common.Author
}
type FeedbackReviewRes struct {
	g.Meta `mime:"application/json"`
}
