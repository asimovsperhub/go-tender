package model

import "github.com/gogf/gf/v2/os/gtime"

type Consultation struct {
	Id         uint64      `json:"id"         description:""`
	Type       string      `json:"type"       description:"类型"`
	Title      string      `json:"title"      description:"标题"`
	Content    string      `json:"content"    description:"内容"`
	Publish    string      `json:"publish"    description:"发布日期"`
	Expiry     string      `json:"expiry"     description:"施行日期"`
	Office     string      `json:"office"     description:"制定机关"`
	Url        string      `json:"url"        description:"外部链接"`
	Attachment string      `json:"attachment" description:"附件"`
	Word       string      `json:"word" description:"word"`
	Pdf        string      `json:"pdf" description:"pdf"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	Source     string      `json:"source"     description:"来源"`
	Status     string      `json:"status"     description:"时效性"`
	Display    int         `json:"display"    description:"是否展示"`
	Height     string      `json:"height"     description:"高度"`
	Rank       string      `json:"rank"       description:"排序"`
	Crawler    int         `json:"crawler"       description:"是否爬取"`
}
