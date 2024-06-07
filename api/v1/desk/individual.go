package desk

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/api/v1/common"
	"tender/internal/app/system/model/entity"
)

// 个人中心

// 企业入驻
type EnterpriseAddReq struct {
	g.Meta             `path:"/enterprise/add" tags:"前台个人中心" method:"post" summary:"企业入驻"`
	Name               string `p:"name" v:"required#企业名称不能为空"`      //企业全称
	NickName           string `p:"nickName"`                        // 企业简称
	Contact            string `p:"contact" v:"phone#手机号格式错误"`       // 联系方式
	EstablishmentAt    string `p:"establishmentAt" v:"date#时间格式错误"` //注册时间
	Location           string `p:"location"`                        // 所在地
	Industry           string `p:"industry"`                        // 行业
	Icon               string `p:"icon"`                            // logo
	License            string `p:"license"`                         // 营业执照
	Certificate        string `p:"certificate"`                     //证书
	Introduction       string `p:"introduction"`                    // 简介
	Content            string `p:"content"`                         // 主页内容
	CertificateMessage string `p:"certificateMessage"`              // 执照审核留言
	LicenseMessage     string `p:"licenseMessage"`                  // 证书审核留言
	common.Author
}
type EnterpriseAddRes struct {
}

// 修改企业信息
type EnterpriseEditReq struct {
	g.Meta             `path:"/enterprise/edit" tags:"前台个人中心" method:"post" summary:"企业信息修改"`
	OldName            string `p:"oldName" v:"required#旧名称不能为空"`    // 老企业名称
	Name               string `p:"name" v:"required#企业名称不能为空"`      //企业全称
	NickName           string `p:"nickname"`                        // 企业简称
	Contact            string `p:"contact" v:"phone#手机号格式错误"`       // 手机号`                      // 联系方式
	EstablishmentAt    string `p:"establishmentAt" v:"date#时间格式错误"` //注册时间
	Location           string `p:"location"`                        // 所在地
	Industry           string `p:"industry"`                        // 行业
	Icon               string `p:"icon"`                            // logo
	License            string `p:"license"`                         // 营业执照
	Certificate        string `p:"certificate"`                     //证书
	Introduction       string `p:"introduction"`                    // 简介
	Content            string `p:"content"`                         // 主页内容
	CertificateMessage string `p:"certificateMessage"`              // 执照审核留言
	LicenseMessage     string `p:"licenseMessage"`                  // 证书审核留言
	common.Author
}
type EnterpriseEditRes struct {
	Result         string                `json:"result" dc:"修改结果"`
	EnterpriseInfo *entity.SysEnterprise `json:"enterpriseInfo"`
}

// 获取企业信息
type EnterpriseGetReq struct {
	g.Meta `path:"/enterprise/get" tags:"前台个人中心" method:"get" summary:"获取企业信息"`
	UserId *int `p:"userId"` // 用户id
	Id     *int `p:"id"`     // 企业id
	common.Author
}
type EnterpriseGetRes struct {
	// Result         string               `json:"result" dc:"修改结果"`
	EnterpriseInfo *entity.SysEnterprise `json:"enterpriseInfo"`
}
