// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Bid is the golang structure for table bid.
type Bid struct {
	Id                     uint64      `json:"id"                     description:""`
	BulletinType           string      `json:"bulletinType"           description:"公告类型"`
	OriginalType           string      `json:"originalType"           description:"原公告类型"`
	Province               string      `json:"province"               description:"所属省份"`
	City                   string      `json:"city"                   description:"所属城市"`
	IndustryClassification string      `json:"industryClassification" description:"行业分类"`
	OriginalClassification string      `json:"originalClassification" description:"原行业分类"`
	ReleaseTime            string      `json:"releaseTime"            description:"发布时间"`
	BidopeningTime         string      `json:"bidopeningTime"         description:"开标时间"`
	Title                  string      `json:"title"                  description:"标题"`
	Abstract               string      `json:"abstract"               description:"摘要"`
	AnnouncementContent    string      `json:"announcementContent"    description:"公告内容"`
	OriginalContent        string      `json:"originalContent"        description:"未清洗的公告内容"`
	Link                   string      `json:"link"                   description:"源链接"`
	OriginalLink           string      `json:"originalLink"           description:"源源链接"`
	Attachment             string      `json:"attachment"             description:"招标附件"`
	Amount                 string      `json:"amount"                 description:"金额"`
	ContactPerson          string      `json:"contactPerson"          description:"联系人"`
	ContactInformation     string      `json:"contactInformation"     description:"联系方式"`
	BidAmount              string      `json:"bidAmount"              description:"中标金额"`
	TenderName             string      `json:"tenderName"             description:"招标企业"`
	WinName                string      `json:"winName"                description:"中标企业"`
	Site                   string      `json:"site"                   description:"站点"`
	CreatedAt              *gtime.Time `json:"createdAt"              description:"爬取日期"`
	Height                 string      `json:"height"                 description:"1高2中3低"`
	Rank                   string      `json:"rank"                   description:"排序"`
	Area                   string      `json:"area"                   description:"地区排序"`
	Display                uint        `json:"display"                description:"是否展示,0不展示1展示"`
	Crawler                uint        `json:"crawler"                description:"是否爬取,0 非爬取 1 爬取"`
	ReleaseDay             string      `json:"releaseDay"             description:"发布时间day排序用的"`
}
