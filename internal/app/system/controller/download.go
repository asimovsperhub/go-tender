package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"strings"
	desk_dao "tender/internal/app/desk/dao"
	desk_entity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/entity"
)

// 模版下载
func (c *DownloadController) DownloadSystemTp(r *ghttp.Request) {
	f := excelize.NewFile()
	f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "知识标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "知识类型")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "一级分类")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "二级分类")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "阅读下载权限")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "积分设置")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "知识内容")
	f.SetCellValue("sheet1", "A"+strconv.Itoa(2), "测试标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(2), "普通")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(2), "招标采购")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(2), "采购制度")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(2), "所有人")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(2), "50")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(2), "知识内容")
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+"template.xlsx")
	r.Response.Write(buff.Bytes())
}

// 会员批量导出
func (c *DownloadController) MemberExport(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	res := []*entity.MemberUser{}
	m := dao.MemberUser.Ctx(r.Context())
	p := strings.Split(url, "ids=")
	var ids []string
	if len(p) > 1 {
		p = strings.Split(p[1], "&")
		ids = strings.Split(p[0], ",")
		log.Println("r.URL.RawQuery", p[0], ids, url)
		m = m.WhereIn("id", ids)
	}
	keywords := strings.Split(url, "keyWords=")
	if len(keywords) > 1 {
		key := strings.Split(keywords[1], "&")[0]
		if key != "" {
			m = m.Where("user_nickname like ?", "%"+key+"%")
		}
	}
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
			return
		}
	})
	if err != nil {
		// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "ID")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "昵称")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "电话")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "会员等级")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "有效期")

	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), strconv.FormatUint(res[i].Id, 10))
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].UserNickname)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Mobile)
		memberlevel := "-"
		switch res[i].MemberLevel {
		case 1:
			memberlevel = "月卡会员"
		case 2:
			memberlevel = "季卡会员"
		case 3:
			memberlevel = "年卡会员"
		}
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), memberlevel)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].MaturityAt)
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"user"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 企业批量导出
func (c *DownloadController) EnterpriseExport(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "ids=")
	m := dao.SysEnterprise.Ctx(r.Context())
	var ids []string
	if len(p) > 1 {
		p = strings.Split(p[1], "&")
		ids = strings.Split(p[0], ",")
		log.Println("r.URL.RawQuery", p[0], ids, url)
		m = m.WhereIn("id", ids)
	}
	keywords := strings.Split(url, "keyWords=")
	if len(keywords) > 1 {
		key := strings.Split(keywords[1], "&")[0]
		if key != "" {
			m = m.Where("name like ?", "%"+key+"%")
		}
	}
	res := []*entity.SysEnterprise{}
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
			return
		}
	})
	if err != nil {
		// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "企业名称")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "所在地")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "所属行业")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "联系电话")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "入驻时间")

	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Name)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Location)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Industry)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Contact)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].CreatedAt)
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"enterprise"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 未认证的企业批量导出
func (c *DownloadController) EnterpriseUncertifiedExport(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "ids=")
	m := dao.SysEnterprise.Ctx(r.Context())
	var ids []string
	if len(p) > 1 {
		p = strings.Split(p[1], "&")
		ids = strings.Split(p[0], ",")
		log.Println("r.URL.RawQuery", p[0], ids, url)
		m = m.WhereIn("id", ids)
	}
	keywords := strings.Split(url, "keyWords=")
	if len(keywords) > 1 {
		key := strings.Split(keywords[1], "&")[0]
		if key != "" {
			m = m.Where("name like ?", "%"+key+"%")
		}
	}
	m = m.Where("license_status <> 1")
	res := []*entity.SysEnterprise{}
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
			return
		}
	})
	if err != nil {
		// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "企业名称")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "所在地")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "所属行业")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "联系电话")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "入驻时间")

	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Name)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Location)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Industry)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Contact)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].CreatedAt)
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"enterprise"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 反馈导出

func (c *DownloadController) FeedbackExport(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "ids=")
	m := desk_dao.Feedback.Ctx(r.Context())
	var ids []string
	if len(p) > 1 {
		p = strings.Split(p[1], "&")
		ids = strings.Split(p[0], ",")
		log.Println("r.URL.RawQuery", p[0], ids, url)
		m = m.WhereIn("id", ids)
	}
	//keywords := strings.Split(url, "keyWords=")
	//if len(keywords) > 1 {
	//	key := strings.Split(keywords[1], "&")[0]
	//	if key != "" {
	//		m = m.Where("name like ?", "%"+key+"%")
	//	}
	//}
	res := []*desk_entity.Feedback{}
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
			return
		}
	})
	if err != nil {
		// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "公司名称")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "联系人")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "联系电话")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "备注")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "附件")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "上传时间")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "状态")

	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Company)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].ContactPerson)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].ContactInformation)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Remarks)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].Attachment)
		f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].CreatedAt)
		f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].Status)
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"feedback"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 财务汇总导出

func (c *DownloadController) FinanceExport(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "ids=")
	m := dao.PurchaseLog.Ctx(r.Context())
	var ids []string
	if len(p) > 1 {
		p = strings.Split(p[1], "&")
		ids = strings.Split(p[0], ",")
		log.Println("r.URL.RawQuery", p[0], ids, url)
		m = m.WhereIn("id", ids)
	}
	mobile := strings.Split(url, "mobile=")
	if len(mobile) > 1 {
		key := strings.Split(mobile[1], "&")[0]
		if key != "" {
			m = m.Where("mobile like ?", "%"+key+"%")
		}
	}
	start := strings.Split(url, "start=")
	if len(start) > 1 {
		key := strings.Split(start[1], "&")[0]
		if key != "" {
			m = m.Where("created_at >=?", key)
		}
	}
	end := strings.Split(url, "end=")
	if len(end) > 1 {
		key := strings.Split(end[1], "&")[0]
		if key != "" {
			m = m.Where("created_at <= ?", key)
		}
	}
	res := []*entity.PurchaseLog{}
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
			return
		}
	})
	if err != nil {
		// r.Response.Write(&DownRes{Code: 50, Message: "查询数据失败"})
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "账户id")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "名称")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "联系电话")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "产品")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "充值金额")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "支付时间")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "支付渠道")
	f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "流水号")
	f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "到期时间")

	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), strconv.Itoa(res[i].UserId))
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].NickName)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Mobile)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Memo)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), fmt.Sprintf("%.3f", float64(res[i].Amount)/float64(100)))
		f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].CreatedAt)
		f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].PayType)
		f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].OrderNo)
		f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), "")
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"finance"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}
