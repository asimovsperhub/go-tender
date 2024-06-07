package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/system/model/entity"
)

type EnterpriseAddReq struct {
	g.Meta          `path:"/enterprise/add" tags:"企业管理" method:"post" summary:"企业入驻"`
	Name            string `p:"name" v:"required#企业名称不能为空"`
	Location        string `p:"location"`
	Contact         string `p:"location"`
	Icon            string `p:"icon"`
	Introduction    string `p:"introduction"`
	Remark          string `p:"remark"`
	License         string `p:"license"`
	Certificate     string `p:"certificate"`
	EstablishmentAt string `p:"establishment"`
	common.Author
}
type EnterpriseAddRes struct {
}
type EnterpriseSearchReq struct {
	g.Meta   `path:"/enterprise/list" tags:"企业管理" method:"get" summary:"企业列表"`
	Name     string `p:"name"`     // 企业名称
	City     string `p:"city"`     // 所在地
	Industry string `p:"industry"` // 行业
	Contact  string `p:"contact"`  // 联系方式
	common.PageReq
	common.Author
}
type EnterpriselLcenseSearchReq struct {
	g.Meta   `path:"/enterprise/license/list" tags:"企业管理" method:"get" summary:"入驻审查列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}

// 	LicenseMessage:     "license_message",
//	CertificateMessage: "certificate_message",
type EnterpriselLcenseReviewReq struct {
	g.Meta           `path:"/enterprise/license/review" tags:"企业管理" method:"post" summary:"入驻审查"`
	Id               uint   `p:"id" v:"required#企业id不能为空"`
	Status           int    `p:"status" v:"required#入驻审查状态不能为空"`
	OpLicenseMessage string `p:"oplicenseMessage" v:"max-length:400#字数超过400"`
	common.Author
}
type EnterpriselLcenseReviewRes struct {
}
type EnterpriseCertificateSearchReq struct {
	g.Meta   `path:"/enterprise/certificate/list" tags:"企业管理" method:"get" summary:"证明书审查列表"`
	KeyWords string `p:"keyWords"`
	common.PageReq
	common.Author
}
type EnterpriseCertificateReviewReq struct {
	g.Meta               `path:"/enterprise/certificate/review" tags:"企业管理" method:"post" summary:"证明书审查"`
	Id                   uint   `p:"id" v:"required#企业id不能为空"`
	Status               int    `p:"status" v:"required#证明书审查状态不能为空"`
	OpCertificateMessage string `p:"opcertificateMessage" v:"max-length:400#字数超过400"`
	common.Author
}

type EnterpriseCertificateReviewRes struct {
}
type EnterpriseSearchRes struct {
	g.Meta         `mime:"application/json"`
	EnterpriseList []*entity.SysEnterprise `json:"EnterpriseList"`
	common.ListRes
}
