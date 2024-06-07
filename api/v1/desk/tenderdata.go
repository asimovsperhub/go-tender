package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	deskentity "tender/internal/app/desk/model/entity"
)

type DataSearchReq struct {
	g.Meta   `path:"/bidding/search" tags:"信息中心" method:"get" summary:"信息中心搜索"`
	KeyWords string `p:"keyWords"`
	// DataType     string `p:"dataType"`
	BulletinType string `p:"bulletinType"`
	// IndustryType string `p:"industryType"`
	// City         string `p:"city"`
	common.PageReq
	common.Author
}

type DataSearchRes struct {
	g.Meta             `mime:"application/json"`
	BiddingInformation []*deskentity.Bid `json:"BiddingInformation"`
	common.ListRes
}
