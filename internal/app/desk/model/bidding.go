package model

import "github.com/gogf/gf/v2/os/gtime"

type BiddingEnterprise struct {
	Id           uint64      `json:"id"                     description:""`
	BulletinType string      `json:"bulletinType"           description:"公告类型"`
	City         string      `json:"city"                   description:"所属城市"`
	ReleaseTime  string      `json:"releaseTime"            description:"发布时间"`
	Title        string      `json:"title"                  description:"标题"`
	Enterprise   string      `json:"enterprise"             description:"企业信息"`
	Link         string      `json:"link"                   description:"源链接"`
	CreatedAt    *gtime.Time `json:"createdAt"              description:"爬取日期"`
	BidAmount    string      `json:"bidAmount"              description:"中标金额"`
	TenderName   string      `json:"tenderName"             description:"招标企业"`
}

type Bidding struct {
	Id                     int64  `json:"id"                     description:"UID"`
	BulletinType           string `json:"bulletinType"           description:"公告类型"`
	NoticeNature           string `json:"noticeNature"           description:"公告性质"`
	City                   string `json:"city"                   description:"所属城市"`
	IndustryClassification string `json:"industryClassification" description:"行业分类"`
	ReleaseTime            string `json:"releaseTime"            description:"发布时间"`
	TenderDeadline         string `json:"tenderDeadline"         description:"投标截止时间"`
	BidopeningTime         string `json:"bidopeningTime"         description:"开标时间"`
	Title                  string `json:"title"                  description:"标题"`
	//AnnouncementContent    string `json:"announcementContent"    description:"公告内容"`
	Abstract           string `json:"abstract"                  description:"摘要"`
	Attachment         string `json:"attachment"             description:"招标附件"`
	Amount             string `json:"amount"                 description:"金额"`
	ContactPerson      string `json:"contactPerson"          description:"联系人"`
	ContactInformation string `json:"contactInformation"     description:"联系方式"`
	ContactContent     string `json:"contactContent"         description:"联系方式内容"`
	Link               string `json:"link"                   description:"源链接"`
}
