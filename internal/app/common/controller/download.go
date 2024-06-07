package controller

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/jung-kurt/gofpdf"
	"github.com/klauspost/compress/zip"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	netUrl "net/url"
	"os"
	"strconv"
	"strings"
	commonsrvice "tender/internal/app/common/service"
	"tender/internal/app/desk/dao"
	desk_model "tender/internal/app/desk/model"
	desk_do "tender/internal/app/desk/model/do"
	desk_entity "tender/internal/app/desk/model/entity"
	"tender/internal/app/desk/service"
	system_dao "tender/internal/app/system/dao"
	"tender/internal/app/system/library/libPdfUtils"
	"tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/library/liberr"
)

// 招投标批量下载
var (
	Download = DownloadController{}
)

type DownloadController struct {
	BaseController
}

type DownRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GenUUID 生成一个随机的唯一ID
func GenUUID() string {
	return uuid.NewV4().String()
}

// 知库下载相关

func GetHisKnowledge(r *ghttp.Request, userid uint64, knowledgeid uint64) (isexit bool, err error) {
	count := 0
	err = g.Try(r.Context(), func(ctx context.Context) {
		m := dao.HisKnowledge.Ctx(r.Context())
		m = m.Where("user_id = ?", userid)
		m = m.Where("knowledge_id = ?", knowledgeid)
		count, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取his知识失败")
		if err != nil {
			return
		}
	})
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func InsertHisKnowledge(r *ghttp.Request, userid uint64, knowledgeid uint64) (err error) {
	data := desk_do.HisKnowledge{}
	data.UserId = userid
	data.KnowledgeId = knowledgeid
	err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.HisKnowledge.Ctx(r.Context()).Insert(data)
			liberr.ErrIsNil(ctx, err, "插入数据his知识失败")
			if err != nil {
				return
			}
		})
		return err
	})
	return

}

func GetUserInfo(r *ghttp.Request) (userinfo *entity.MemberUser, err error) {
	userinfo = (*entity.MemberUser)(nil)
	err = g.Try(r.Context(), func(ctx context.Context) {
		m := system_dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			system_dao.MemberUser.Columns().Id, service.Context().GetUserId(r.Context())))
		err = m.Limit(1).Scan(&userinfo)
	})
	return
}
func UpdateDownUser(r *ghttp.Request, userinfo *entity.MemberUser, knowledge *desk_entity.MemberKnowledge) (err error) {
	log.Println("积分-------------------->", userinfo.Integral, knowledge.IntegralSetting)
	if userinfo.Integral < knowledge.IntegralSetting {
		err = errors.New("用户积分不足")
		return err
	}
	err = g.Try(r.Context(), func(ctx context.Context) {
		m := system_dao.MemberUser.Ctx(ctx)
		m = m.Where(fmt.Sprintf("%s='%d'",
			system_dao.MemberUser.Columns().Id, service.Context().GetUserId(r.Context())))
		_, err = m.Update(&do.MemberUser{Integral: userinfo.Integral - knowledge.IntegralSetting})
		liberr.ErrIsNil(ctx, err, "更新用户积分失败")
	})

	return
}

// word xml
func makeDocumentXMLFile(txt string) string {
	return `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas" xmlns:cx="http://schemas.microsoft.com/office/drawing/2014/chartex" xmlns:cx1="http://schemas.microsoft.com/office/drawing/2015/9/8/chartex" xmlns:cx2="http://schemas.microsoft.com/office/drawing/2015/10/21/chartex" xmlns:cx3="http://schemas.microsoft.com/office/drawing/2016/5/9/chartex" xmlns:cx4="http://schemas.microsoft.com/office/drawing/2016/5/10/chartex" xmlns:cx5="http://schemas.microsoft.com/office/drawing/2016/5/11/chartex" xmlns:cx6="http://schemas.microsoft.com/office/drawing/2016/5/12/chartex" xmlns:cx7="http://schemas.microsoft.com/office/drawing/2016/5/13/chartex" xmlns:cx8="http://schemas.microsoft.com/office/drawing/2016/5/14/chartex" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:aink="http://schemas.microsoft.com/office/drawing/2016/ink" xmlns:am3d="http://schemas.microsoft.com/office/drawing/2017/model3d" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex" xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid" xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk" xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" mc:Ignorable="w14 w15 w16se w16cid w16 w16cex wp14">
    <w:body>
        <w:p w14:paraId="396385E8" w14:textId="6BF766E2" w:rsidR="00947F74" w:rsidRDefault="006640DB">
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" />
                </w:rPr>
                <w:t>` + txt + `</w:t>
            </w:r>
            <w:r>
                <w:t xml:space="preserve"></w:t>
            </w:r>
        </w:p>
        <w:sectPr w:rsidR="00947F74">
            <w:pgSz w:w="11906" w:h="16838" />
            <w:pgMar w:top="1440" w:right="1800" w:bottom="1440" w:left="1800" w:header="851" w:footer="992" w:gutter="0" />
            <w:cols w:space="425" />
            <w:docGrid w:type="lines" w:linePitch="312" />
        </w:sectPr>
    </w:body>
</w:document>
	`
}

// word
func (c *DownloadController) WriteWord(content string) []byte {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()
	partFile, _ := zipWriter.Create("word/document.xml")
	partFile.Write([]byte(makeDocumentXMLFile(content)))
	// Make sure to check the error on Close.
	zipWriter.Close()
	return buf.Bytes()
}

// 写 pdf  这个只有在word下载里边用
func (c *DownloadController) WritePdf(title string, date string, content string, page int) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "./resource/ttf/")
	// titleStr := title
	// pdf.SetFont("dejavu", "", 14)
	pdf.AddUTF8Font("FANGSONG", "", "FANGSONG.ttf")
	// pdf.SetTitle(titleStr, false)
	//pdf.SetAuthor("Jules Verne", false)
	//pdf.SetHeaderFunc(func() {
	//	// Arial bold 15
	//	pdf.SetFont("PingFang", "", 15)
	//	// Calculate width of title and position
	//	wd := pdf.GetStringWidth(titleStr) + 6
	//	pdf.SetX((210 - wd) / 2)
	//	// Colors of frame, background and text
	//	pdf.SetDrawColor(0, 80, 180)
	//	pdf.SetFillColor(230, 230, 0)
	//	pdf.SetTextColor(220, 50, 50)
	//	// Thickness of frame (1 mm)
	//	pdf.SetLineWidth(1)
	//	// Title
	//	pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
	//	// Line break
	//	pdf.Ln(10)
	//})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("FANGSONG", "", 8)
		// Text color in gray
		// pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	//chapterTitle := func(chapNum int, titleStr string) {
	//	// 	// Arial 12
	//	pdf.SetFont("FANGSONG", "", 12)
	//	// Background color
	//	pdf.SetFillColor(200, 220, 255)
	//	// Title
	//	pdf.CellFormat(0, 6, fmt.Sprintf("发布时间 : %s", titleStr),
	//		"", 1, "L", true, 0, "")
	//	// Line break
	//	pdf.Ln(4)
	//}
	chapterBody := func(txtStr string) {
		// Read text file
		//txtStr, err := ioutil.ReadFile(fileStr)
		//if err != nil {
		//	pdf.SetError(err)
		//}
		// Times 12
		pdf.SetFont("FANGSONG", "", 0)
		// Output justified text
		pdf.MultiCell(0, 5, string(txtStr), "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("FANGSONG", "", 0)
		// pdf.Cell(0, 5, "(end of excerpt)")
	}
	printChapter := func(chapNum int, titleStr, fileStr string) {
		// 去掉新增页
		pdf.AddPage()
		// chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}
	printChapter(page, date, content)
	// err := pdf.OutputFileAndClose("./test.pdf")
	var buff bytes.Buffer
	err := pdf.Output(&buff)
	//log.Println(buf.Bytes())
	if err != nil {
		log.Println("Error generating PDF: ", err)
	}
	return buff.Bytes()
}

// 招标 批量/单个 excel
func (c *DownloadController) DownloadFile(r *ghttp.Request) {
	// GetJson解析当前请求内容为JSON格式，并返回JSON对象。
	//注意，请求内容是从request BODY读取的，而不是从FORM的任何字段读取的。
	//reqdata, _ := r.GetJson()
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := []*desk_entity.Bid{}
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.WhereIn("id", ids)
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
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "序号")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "标题")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "项目名称")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "采购单位")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "联系人")
	f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系电话")
	f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "城市")
	f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "类型")
	f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "预算金额")
	f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "发布时间")
	f.SetCellValue("sheet1", "K"+strconv.Itoa(1), "具体内容")

	userinfo, _ := GetUserInfo(r)
	if userinfo != nil {
		// 付费会员
		if userinfo.MemberLevel > 0 {
			for i := 0; i < len(res); i++ {
				log.Println(res[i].Title)
				f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), strconv.Itoa(i+1))
				f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
				f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Title)
				f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].TenderName)
				f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].ContactPerson)
				f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactInformation)
				f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].City)
				f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].BulletinType)
				f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].Amount)
				f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].ReleaseTime)
				f.SetCellValue("sheet1", "K"+strconv.Itoa(i+2), res[i].Link)
			}
		} else {
			r.Response.Write(&DownRes{Code: 50, Message: "非会员不能批量下载"})
			return
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "请先注册/登陆"})
		return
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"tender"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 招标单个pdf
func (c *DownloadController) DownloadFilePdf(r *ghttp.Request) {
	// GetJson解析当前请求内容为JSON格式，并返回JSON对象。
	//注意，请求内容是从request BODY读取的，而不是从FORM的任何字段读取的。
	//reqdata, _ := r.GetJson()
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_entity.Bid)(nil)
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
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
	var buff []byte
	userinfo, _ := GetUserInfo(r)
	if res != nil {
		if userinfo != nil {
			// buff = c.WritePdf(res.Title, res.ReleaseTime, res.AnnouncementContent, 1)
			//if userinfo.MemberLevel > 0 {
			//}
			strings.Replace(res.Title, " ", "", -1)
			PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
			isFile := FileExit(PdfPath)
			if !isFile {
				buff = HtmlToPdf([]byte(res.AnnouncementContent))
				distFile, err := gfile.Create(PdfPath)
				if err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
				if _, err = io.Copy(distFile, bytes.NewReader(buff)); err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
			}
			buff, err = os.ReadFile(PdfPath)
			if err != nil {
				fmt.Println(err)
			}

		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", netUrl.QueryEscape(res.Title)))

	r.Response.Write(buff)
}

// 招标单个word
func (c *DownloadController) DownloadFileWord(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_entity.Bid)(nil)
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
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
	var buff []byte
	var wordbuff []byte
	userinfo, _ := GetUserInfo(r)
	if res != nil {
		if userinfo != nil {
			strings.Replace(res.Title, " ", "", -1)
			PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
			WordPath := "/mnt/tender/download/" + res.Title + ".docx"
			isFile := FileExit(PdfPath)
			if !isFile {
				buff = c.WritePdf(res.Title, res.ReleaseTime, res.AnnouncementContent, 1)
				// buff = HtmlToPdf([]byte(res.Content))
				distFile, err := gfile.Create(PdfPath)
				if err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
				if _, err := io.Copy(distFile, bytes.NewReader(buff)); err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
			}
			result, err := libPdfUtils.ConvertPdfToWord(PdfPath, "/mnt/tender/download")
			fmt.Printf(result)
			wordbuff, err = os.ReadFile(WordPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", netUrl.QueryEscape(res.Title)))
	r.Response.Write(wordbuff)
}

// 单个知识 execl
func (c *DownloadController) DownloadKnowledgeFile(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()

	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_entity.MemberKnowledge)(nil)
	m := dao.MemberKnowledge.Ctx(r.Context())
	m = m.WhereIn("id", ids)
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
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "类型")
	f.SetCellStr("sheet1", "C"+strconv.Itoa(1), "时间")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "内容")
	userinfo, _ := GetUserInfo(r)
	var buff []byte
	if res != nil {
		if userinfo != nil {
			if res.Authority == 1 {
				if userinfo.Id != uint64(res.UserId) {
					is_, _ := GetHisKnowledge(r, userinfo.Id, res.Id)
					if !is_ {
						if userinfo.MemberLevel > 0 {
							err = UpdateDownUser(r, userinfo, res)
							if err != nil {
								r.Response.Write(&DownRes{Code: 50, Message: "用户积分不足"})
								return
							}
							err = InsertHisKnowledge(r, userinfo.Id, res.Id)
							if err != nil {
								return
							}
						} else {
							r.Response.Write(&DownRes{Code: 50, Message: "用户非会员并未购买"})
							return
						}
					}
				}
			}
			log.Println(res.CreatedAt, res.CreatedAt.Format("Y-m-d H:i:s.u"))
			f.SetCellStr("sheet1", "A"+strconv.Itoa(2), res.Title)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(2), res.SecondaryClassification)
			f.SetCellStr("sheet1", "C"+strconv.Itoa(2), res.CreatedAt.Format("Y-m-d H:i:s.u"))
			f.SetCellValue("sheet1", "D"+strconv.Itoa(2), res.Content)
			// 写内存 文件大的容易崩
			buff_, _ := f.WriteToBuffer()
			if buff_ != nil {
				buff = buff_.Bytes()
			}
		} else {
			r.Response.Write(&DownRes{Code: 50, Message: "用户未登陆"})
			return
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	strings.Replace(res.Title, " ", "", -1)
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", netUrl.QueryEscape(res.Title)))
	r.Response.Write(buff)
}

// 单个知识pdf / 附件
func (c *DownloadController) DownloadKnowledgeFilePdf(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()

	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_entity.MemberKnowledge)(nil)
	m := dao.MemberKnowledge.Ctx(r.Context())
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	var buff []byte
	userinfo, _ := GetUserInfo(r)
	dwfilename := ""
	if res != nil {
		if userinfo != nil {
			if res.Authority == 1 {
				if userinfo.Id != uint64(res.UserId) {
					// 下载过/买过
					is_, _ := GetHisKnowledge(r, userinfo.Id, res.Id)
					if !is_ {
						if userinfo.MemberLevel > 0 {
							err = UpdateDownUser(r, userinfo, res)
							if err != nil {
								r.Response.Write(&DownRes{Code: 50, Message: "用户积分不足"})
								return
							}
							err = InsertHisKnowledge(r, userinfo.Id, res.Id)
							if err != nil {
								return
							}
						} else {
							r.Response.Write(&DownRes{Code: 50, Message: "用户非会员并未购买"})
							return
						}
					}
				}
			} else {
				if userinfo.Id != uint64(res.UserId) {
					// 下载过/买过
					is_, _ := GetHisKnowledge(r, userinfo.Id, res.Id)
					if !is_ {
						err = UpdateDownUser(r, userinfo, res)
						if err != nil {
							r.Response.Write(&DownRes{Code: 50, Message: "用户积分不足"})
							return
						}
						err = InsertHisKnowledge(r, userinfo.Id, res.Id)
						if err != nil {
							return
						}
					}
				}

			}
			// 有附件下载附件没有转word
			if res.AttachmentUrl != "" {
				path := strings.Split(res.AttachmentUrl, "tender")
				if len(path) > 1 {
					object := commonsrvice.S3.GetObject("tender", path[1])
					defer object.Close()
					buff, err = io.ReadAll(object)
					if err != nil {
						fmt.Println(err)
					}
					dwfilename = strings.Replace(path[1], "/", "", -1)
				}
			} else {
				strings.Replace(res.Title, " ", "", -1)
				PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
				// buff = c.WritePdf(res.Title, res.CreatedAt.Format("Y-m-d H:i:s.u"), res.Content, 1)
				log.Println(res.Content)
				//buff = HtmlToPdf([]byte(res.Content))
				isFile := FileExit(PdfPath)
				log.Println(isFile)
				if !isFile {
					buff = HtmlToPdf([]byte(res.Content))
					distFile, err := gfile.Create(PdfPath)
					if err != nil {
						log.Println(err)
						r.Response.Write(err)
						return
					}
					if _, err = io.Copy(distFile, bytes.NewReader(buff)); err != nil {
						log.Println(err)
						r.Response.Write(err)
						return
					}
				}
				buff, err = os.ReadFile(PdfPath)
				if err != nil {
					fmt.Println(err)
				}
				dwfilename = fmt.Sprintf("%s.pdf", res.Title)
			}
		} else {
			r.Response.Write(&DownRes{Code: 50, Message: "用户未登陆"})
			return
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", netUrl.QueryEscape(dwfilename)))
	r.Response.Write(buff)
}

// 单个知识word
func (c *DownloadController) DownloadKnowledgeFileWord(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()

	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_entity.MemberKnowledge)(nil)
	m := dao.MemberKnowledge.Ctx(r.Context())
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	var buff []byte
	var wordbuff []byte
	userinfo, _ := GetUserInfo(r)
	if res != nil {
		if userinfo != nil {
			if res.Authority == 1 {
				if userinfo.Id != uint64(res.UserId) {
					is_, _ := GetHisKnowledge(r, userinfo.Id, res.Id)
					if !is_ {
						if userinfo.MemberLevel > 0 {
							err = UpdateDownUser(r, userinfo, res)
							if err != nil {
								r.Response.Write(&DownRes{Code: 50, Message: "用户积分不足"})
								return
							}
							err = InsertHisKnowledge(r, userinfo.Id, res.Id)
							if err != nil {
								return
							}
						} else {
							r.Response.Write(&DownRes{Code: 50, Message: "用户非会员并未购买"})
							return
						}
					}
				}
			}
			strings.Replace(res.Title, " ", "", -1)
			PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
			WordPath := "/mnt/tender/download/" + res.Title + ".docx"
			isFile := FileExit(PdfPath)
			if !isFile {
				buff = c.WritePdf(res.Title, res.CreatedAt.Format("Y-m-d H:i:s.u"), res.Content, 1)
				// buff = HtmlToPdf([]byte(res.Content))
				distFile, err := gfile.Create(PdfPath)
				if err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
				if _, err := io.Copy(distFile, bytes.NewReader(buff)); err != nil {
					log.Println(err)
					r.Response.Write(err)
					return
				}
				result, err := libPdfUtils.ConvertPdfToWord(PdfPath, "/mnt/tender/download")
				fmt.Printf(result)
			}
			wordbuff, err = os.ReadFile(WordPath)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			r.Response.Write(&DownRes{Code: 50, Message: "用户未登陆"})
			return
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", netUrl.QueryEscape(res.Title)))
	r.Response.Write(wordbuff)
}

// 政策资讯 批量/单个 execl下载
func (c *DownloadController) DownloadConsultationFile(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := ([]*desk_entity.Consultation)(nil)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	m = m.WhereIn("id", ids)
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
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "标题")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "来源")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "定制机关")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "发布时间")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "内容")
	if res != nil {
		for i := 0; i < len(res); i++ {
			log.Println(res[i].Title)
			f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Url)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Office)
			f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Publish)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].Content)
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	// 写内存 文件大的容易崩
	buff, _ := f.WriteToBuffer()
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", "attachment;filename="+"consultation"+GenUUID()+".xlsx")
	r.Response.Write(buff.Bytes())
}

// 政策咨询不需要登陆和vip pdf下载
func (c *DownloadController) DownloadConsultationFilePdf(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()

	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_model.Consultation)(nil)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	var buff []byte
	// userinfo, _ := GetUserInfo(r)
	if res != nil {
		strings.Replace(res.Title, " ", "", -1)
		PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
		isFile := FileExit(PdfPath)
		if !isFile {
			buff = HtmlToPdf([]byte(res.Content))
			distFile, err := gfile.Create(PdfPath)
			if err != nil {
				log.Println(err)
				r.Response.Write(err)
				return
			}
			if _, err = io.Copy(distFile, bytes.NewReader(buff)); err != nil {
				log.Println(err)
				r.Response.Write(err)
				return
			}
		}
		buff, err = os.ReadFile(PdfPath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", netUrl.QueryEscape(res.Title)))
	r.Response.Write(buff)
}

// 政策资讯 单个 word下载
func (c *DownloadController) DownloadConsultationFileWord(r *ghttp.Request) {
	// reqdata, _ := r.GetJson()

	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := (*desk_model.Consultation)(nil)
	m := g.DB("data").Schema("crawldata").Model("consultation")
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Limit(1).Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	var buff []byte
	var wordbuff []byte
	// userinfo, _ := GetUserInfo(r)
	if res != nil {
		strings.Replace(res.Title, " ", "", -1)
		PdfPath := "/mnt/tender/download/" + res.Title + ".pdf"
		WordPath := "/mnt/tender/download/" + res.Title + ".docx"
		isFile := FileExit(PdfPath)
		if !isFile {
			buff = c.WritePdf(res.Title, res.Publish, res.Content, 1)
			// buff = HtmlToPdf([]byte(res.Content))
			distFile, err := gfile.Create(PdfPath)
			if err != nil {
				log.Println(err)
				r.Response.Write(err)
				return
			}
			// reader.Read(buf)
			if _, err := io.Copy(distFile, bytes.NewReader(buff)); err != nil {
				log.Println(err)
				r.Response.Write(err)
				return
			}
		}
		result, err := libPdfUtils.ConvertPdfToWord(PdfPath, "/mnt/tender/download")
		fmt.Printf(result)
		wordbuff, err = os.ReadFile(WordPath)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		r.Response.Write(&DownRes{Code: 50, Message: "内容不存在"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream;charset=utf-8")
	header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", netUrl.QueryEscape(res.Title)))
	r.Response.Write(wordbuff)
}

func HtmlToPdf(html []byte) []byte {
	// Initialize library.
	if err := pdf.Init(); err != nil {
		log.Fatal(err)
	}
	defer pdf.Destroy()

	// Create object from file.
	//object, err := pdf.NewObject("sample1.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//object.Header.ContentCenter = "[title]"
	//object.Header.DisplaySeparator = true

	// Create object from URL.
	//object2, err := pdf.NewObject("https://google.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//object2.Footer.ContentLeft = "[date]"
	//object2.Footer.ContentCenter = "Sample footer information"
	//object2.Footer.ContentRight = "[page]"
	//object2.Footer.DisplaySeparator = true

	// Create object from reader.
	//inFile, err := os.Open("sample2.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer inFile.Close()

	object3, err := pdf.NewObjectFromReader(bytes.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	object3.Zoom = 1.5
	object3.TOC.Title = "Table of Contents"

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()
	converter.Add(object3)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	//outFile, err := os.Create("out.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer outFile.Close()

	var buff bytes.Buffer
	// Run converter.
	if err := converter.Run(&buff); err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}
func FileExit(file string) bool {
	_, err := os.Stat(file)
	if err == nil {
		fmt.Println("File exist")
		return true
	}
	if os.IsNotExist(err) {
		fmt.Println("File not exist")
		return false
	}
	fmt.Println("File error")
	return false
}
