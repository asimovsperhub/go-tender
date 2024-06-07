package controller

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/klauspost/compress/zip"
	"github.com/xuri/excelize/v2"
	"io"
	"log"
	netUrl "net/url"
	"strconv"
	"strings"
	"tender/internal/app/common/qbucket"
	commonsrvice "tender/internal/app/common/service"
	"tender/internal/app/desk/dao"
	desk_entity "tender/internal/app/desk/model/entity"
	system_dao "tender/internal/app/system/dao"
	"tender/internal/app/system/library/libDownloadAtt"
	"tender/internal/app/system/model/entity"
	"tender/library/libUtils"
	"time"
)

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

func ZipFiles(files []string, buffs [][]byte, name string, content []byte) (bf []byte, err error) {
	buf := new(bytes.Buffer)
	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)
	work := libUtils.NewPool(10)
	for i, file := range files {
		log.Println("zip--filename----->", file)
		work.Add(1)
		go func(i int, file string, wg *libUtils.WaitGroup) {
			defer wg.Done()
			if buffs[i] != nil {
				zipFile, err := zipWriter.Create("/attachmen/" + file)
				if err != nil {
					log.Println("zip--附件-create err---->", file, err)
				}
				//bytes, _ := os.ReadFile(file)
				log.Println("buff  len ----------->", len(buffs[i]))
				_, err = zipFile.Write(buffs[i])
				if err != nil {
					log.Println("zip--附件-write err---->", file, err)
				}
			}
		}(i, file, work)
		log.Println("waiting...")
		work.Wait()
		log.Println("done")
		defer func() {}()
	}
	// 批量的execl 文件
	zipFile, err := zipWriter.Create(name)
	if err != nil {
		log.Println("zip--批量execl-create---->", err)
	}
	_, err = zipFile.Write(content)
	if err != nil {
		log.Println("zip--批量execl-write---->", err)
	}

	// Make sure to check the error on Close.
	err = zipWriter.Close()
	if err != nil {
		return nil, err
	}

	//write the zipped file to the disk
	return buf.Bytes(), err
}

//  知识批量下载
func (c *DownloadController) DownloadKnowledgeFile(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := []*desk_entity.MemberKnowledge{}
	m := dao.MemberKnowledge.Ctx(r.Context())
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	//files := []string{}
	//if len(files) < 1 {
	//	r.Response.Write(&DownRes{Code: 50, Message: "没有附件"})
	//	return
	//}
	log.Println("res-------------------->", res)
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "标题")
	f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "类型")
	f.SetCellStr("sheet1", "C"+strconv.Itoa(1), "时间")
	f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "内容")
	var filenames []string
	var buffs [][]byte
	for i := 0; i < len(res); i++ {
		// 添加附件
		if res[i].AttachmentUrl != "" {
			log.Println(res[i].AttachmentUrl)
			path := strings.Split(res[i].AttachmentUrl, "tender")
			var buff []byte
			if len(path) > 1 {
				object := commonsrvice.S3.GetObject("tender", path[1])
				defer object.Close()
				buff, err = io.ReadAll(object)
				if err != nil {
					fmt.Println(err)
				}
				filenames = append(filenames, strings.Replace(path[1], "/", "", -1))
				buffs = append(buffs, buff)
			}
			//if len(path) > 1 {
			//	// minio 位置
			//	log.Println("/mnt/tender" + path[1])
			//	files = append(files, "/mnt/tender"+path[1])
			//}
		}
		curl := ""
		if res[i].Url != "" {
			curl = res[i].Url
		} else {
			curl = fmt.Sprintf("https://www.biaoziku.com/onlineKnowledgeBase/lookMore/%d", res[i].Id)
		}
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Title)
		f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].SecondaryClassification)
		f.SetCellStr("sheet1", "C"+strconv.Itoa(i+2), res[i].CreatedAt.Format("Y-m-d H:i:s.u"))
		f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), curl)
	}
	execl_bf, _ := f.WriteToBuffer()
	var buff []byte
	if buff, err = ZipFiles(filenames, buffs, "/knowledge/xlsx/knowledge.xlsx", execl_bf.Bytes()); err != nil {
		log.Println(err)
		r.Response.Write(&DownRes{Code: 50, Message: "压缩文件失败"})
		return
	}
	// 写内存 文件大的容易崩
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fmt.Sprintf("%s.zip", fmt.Sprintf("%s", gtime.Now())+netUrl.QueryEscape("-在线知库")))
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}

// 招标批量下载
// noticeFileBOList
func (c *DownloadController) DownloadTenderFile(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	var butype string
	typeSp := strings.Split(url, "type=")
	if len(typeSp) > 1 {
		butype = strings.Split(typeSp[1], "&")[0]
	}
	var ids []string
	idsSp := strings.Split(url, "ids=")
	if len(idsSp) > 1 {
		idsL := strings.Split(idsSp[1], "&")[0]
		ids = strings.Split(idsL, ",")
	}
	res := []*desk_entity.Bid{}
	m := g.DB("data").Schema("crawldata").Model("bid")
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	var files []string
	var buffs [][]byte
	work := libUtils.NewPool(10)
	f := excelize.NewFile()
	title := "-招标信息"
	switch butype {
	case "招标公告":
		title += "-招标公告"
		//SetCellStr
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "标题")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "城市")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "类型")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "采购单位名称")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系人")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "联系电话")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "预算金额")
		f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "发布时间")
		f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].City)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].IndustryClassification)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactPerson)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].Amount)
			f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].ReleaseTime)
			f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].Link)
		}
	case "在建公告":
		title += "-在建公告"
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "标题")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "城市")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "类型")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "采购单位名称")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系人")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "联系电话")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "预算金额")
		f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "开标时间")
		f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].City)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].IndustryClassification)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactPerson)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].Amount)
			f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].BidopeningTime)
			f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].Link)
		}
	case "意向公开":
		title += "-意向公开"
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "标题")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "城市")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "类型")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "采购单位名称")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系人")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "联系电话")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "预算金额")
		f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "发布时间")
		f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].City)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].IndustryClassification)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactPerson)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].Amount)
			f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].ReleaseTime)
			f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].Link)
		}
	case "中标公告":
		title += "-中标公告"
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "标题")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "城市")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "类型")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "采购单位名称")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系人")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "联系电话")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "预算金额")
		f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "中标金额")
		f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "中标单位名称")
		f.SetCellValue("sheet1", "K"+strconv.Itoa(1), "发布时间")
		f.SetCellValue("sheet1", "L"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].City)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].IndustryClassification)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactPerson)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].Amount)
			f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].WinName)
			f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].BidAmount)
			f.SetCellValue("sheet1", "K"+strconv.Itoa(i+2), res[i].ReleaseTime)
			f.SetCellValue("sheet1", "L"+strconv.Itoa(i+2), res[i].Link)
		}
	case "合同公告":
		title += "-合同公告"
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "合同名称")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "甲方")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "甲方联系电话")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "乙方")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "乙方联系电话")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "合同金额")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "合同签订时间")
		f.SetCellValue("sheet1", "I"+strconv.Itoa(1), "合同公告时间")
		f.SetCellValue("sheet1", "J"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].WinName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].Amount)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].ReleaseTime)
			f.SetCellValue("sheet1", "I"+strconv.Itoa(i+2), res[i].ReleaseTime)
			f.SetCellValue("sheet1", "J"+strconv.Itoa(i+2), res[i].Link)
		}
	case "其他公告":
		title += "-其他公告"
		f.SetCellValue("sheet1", "A"+strconv.Itoa(1), "序号")
		f.SetCellValue("sheet1", "B"+strconv.Itoa(1), "标题")
		f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "城市")
		f.SetCellValue("sheet1", "D"+strconv.Itoa(1), "类型")
		f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "采购单位名称")
		f.SetCellValue("sheet1", "F"+strconv.Itoa(1), "联系人")
		f.SetCellValue("sheet1", "G"+strconv.Itoa(1), "联系电话")
		f.SetCellValue("sheet1", "H"+strconv.Itoa(1), "具体内容")
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			f.SetCellValue("sheet1", "A"+strconv.Itoa(i+2), i+1)
			f.SetCellValue("sheet1", "B"+strconv.Itoa(i+2), res[i].Title)
			f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].City)
			f.SetCellValue("sheet1", "D"+strconv.Itoa(i+2), res[i].IndustryClassification)
			f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].TenderName)
			f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].ContactPerson)
			f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].ContactInformation)
			f.SetCellValue("sheet1", "H"+strconv.Itoa(i+2), res[i].Link)
		}
	default:
		f.DeleteSheet("Sheet1")
		title += "-综合"
		zb := 0
		zj := 0
		win := 0
		ht := 0
		yx := 0
		qt := 0
		for i := 0; i < len(res); i++ {
			if res[i].Attachment != "" {
				for _, url_ := range strings.Split(res[i].Attachment, ",") {
					work.Add(1)
					go func(url string, wg *libUtils.WaitGroup) {
						defer wg.Done()
						path := strings.Split(url, ".com")
						if len(path) > 1 {
							buff := qbucket.NewQbucketService("announcement-1317075511").GetObject(path[1])
							name := strings.Replace(path[1], "/", "", -1)
							files = append(files, name)
							buffs = append(buffs, buff)
						}
					}(url_, work)
					log.Println("waiting...")
					work.Wait()
					log.Println("done")
					defer func() {}()
				}
			}
			switch res[i].BulletinType {
			case "招标公告":
				tm2, _ := time.Parse("2006-01-02 15:04:05", res[i].BidopeningTime)
				if tm2.Unix() > time.Now().Unix() {
					f.NewSheet("在建公告")
					f.SetCellValue("在建公告", "A"+strconv.Itoa(1), "序号")
					f.SetCellValue("在建公告", "B"+strconv.Itoa(1), "标题")
					f.SetCellValue("在建公告", "C"+strconv.Itoa(1), "城市")
					f.SetCellValue("在建公告", "D"+strconv.Itoa(1), "类型")
					f.SetCellValue("在建公告", "E"+strconv.Itoa(1), "采购单位名称")
					f.SetCellValue("在建公告", "F"+strconv.Itoa(1), "联系人")
					f.SetCellValue("在建公告", "G"+strconv.Itoa(1), "联系电话")
					f.SetCellValue("在建公告", "H"+strconv.Itoa(1), "预算金额")
					f.SetCellValue("在建公告", "I"+strconv.Itoa(1), "开标时间")
					f.SetCellValue("在建公告", "J"+strconv.Itoa(1), "具体内容")
					f.SetCellValue("在建公告", "A"+strconv.Itoa(zj+2), zj+1)
					f.SetCellValue("在建公告", "B"+strconv.Itoa(zj+2), res[i].Title)
					f.SetCellValue("在建公告", "C"+strconv.Itoa(zj+2), res[i].City)
					f.SetCellValue("在建公告", "D"+strconv.Itoa(zj+2), res[i].IndustryClassification)
					f.SetCellValue("在建公告", "E"+strconv.Itoa(zj+2), res[i].TenderName)
					f.SetCellValue("在建公告", "F"+strconv.Itoa(zj+2), res[i].ContactPerson)
					f.SetCellValue("在建公告", "G"+strconv.Itoa(zj+2), res[i].ContactInformation)
					f.SetCellValue("在建公告", "H"+strconv.Itoa(zj+2), res[i].Amount)
					f.SetCellValue("在建公告", "I"+strconv.Itoa(zj+2), res[i].BidopeningTime)
					f.SetCellValue("在建公告", "J"+strconv.Itoa(zj+2), res[i].Link)
					zj += 1

				} else {
					f.NewSheet("招标公告")
					f.SetCellValue("招标公告", "A"+strconv.Itoa(1), "序号")
					f.SetCellValue("招标公告", "B"+strconv.Itoa(1), "标题")
					f.SetCellValue("招标公告", "C"+strconv.Itoa(1), "城市")
					f.SetCellValue("招标公告", "D"+strconv.Itoa(1), "类型")
					f.SetCellValue("招标公告", "E"+strconv.Itoa(1), "采购单位名称")
					f.SetCellValue("招标公告", "F"+strconv.Itoa(1), "联系人")
					f.SetCellValue("招标公告", "G"+strconv.Itoa(1), "联系电话")
					f.SetCellValue("招标公告", "H"+strconv.Itoa(1), "预算金额")
					f.SetCellValue("招标公告", "I"+strconv.Itoa(1), "发布时间")
					f.SetCellValue("招标公告", "J"+strconv.Itoa(1), "具体内容")
					f.SetCellValue("招标公告", "A"+strconv.Itoa(zb+2), zb+1)
					f.SetCellValue("招标公告", "B"+strconv.Itoa(zb+2), res[i].Title)
					f.SetCellValue("招标公告", "C"+strconv.Itoa(zb+2), res[i].City)
					f.SetCellValue("招标公告", "D"+strconv.Itoa(zb+2), res[i].IndustryClassification)
					f.SetCellValue("招标公告", "E"+strconv.Itoa(zb+2), res[i].TenderName)
					f.SetCellValue("招标公告", "F"+strconv.Itoa(zb+2), res[i].ContactPerson)
					f.SetCellValue("招标公告", "G"+strconv.Itoa(zb+2), res[i].ContactInformation)
					f.SetCellValue("招标公告", "H"+strconv.Itoa(zb+2), res[i].Amount)
					f.SetCellValue("招标公告", "I"+strconv.Itoa(zb+2), res[i].ReleaseTime)
					f.SetCellValue("招标公告", "J"+strconv.Itoa(zb+2), res[i].Link)
					zb += 1
				}
			case "意向公开":
				f.NewSheet("意向公开")
				f.SetCellValue("意向公开", "A"+strconv.Itoa(1), "序号")
				f.SetCellValue("意向公开", "B"+strconv.Itoa(1), "标题")
				f.SetCellValue("意向公开", "C"+strconv.Itoa(1), "城市")
				f.SetCellValue("意向公开", "D"+strconv.Itoa(1), "类型")
				f.SetCellValue("意向公开", "E"+strconv.Itoa(1), "采购单位名称")
				f.SetCellValue("意向公开", "F"+strconv.Itoa(1), "联系人")
				f.SetCellValue("意向公开", "G"+strconv.Itoa(1), "联系电话")
				f.SetCellValue("意向公开", "H"+strconv.Itoa(1), "预算金额")
				f.SetCellValue("意向公开", "I"+strconv.Itoa(1), "发布时间")
				f.SetCellValue("意向公开", "J"+strconv.Itoa(1), "具体内容")
				f.SetCellValue("意向公开", "A"+strconv.Itoa(yx+2), yx+1)
				f.SetCellValue("意向公开", "B"+strconv.Itoa(yx+2), res[i].Title)
				f.SetCellValue("意向公开", "C"+strconv.Itoa(yx+2), res[i].City)
				f.SetCellValue("意向公开", "D"+strconv.Itoa(yx+2), res[i].IndustryClassification)
				f.SetCellValue("意向公开", "E"+strconv.Itoa(yx+2), res[i].TenderName)
				f.SetCellValue("意向公开", "F"+strconv.Itoa(yx+2), res[i].ContactPerson)
				f.SetCellValue("意向公开", "G"+strconv.Itoa(yx+2), res[i].ContactInformation)
				f.SetCellValue("意向公开", "H"+strconv.Itoa(yx+2), res[i].Amount)
				f.SetCellValue("意向公开", "I"+strconv.Itoa(yx+2), res[i].ReleaseTime)
				f.SetCellValue("意向公开", "J"+strconv.Itoa(yx+2), res[i].Link)
				yx += 1
			case "中标公告":
				f.NewSheet("中标公告")
				f.SetCellValue("中标公告", "A"+strconv.Itoa(1), "序号")
				f.SetCellValue("中标公告", "B"+strconv.Itoa(1), "标题")
				f.SetCellValue("中标公告", "C"+strconv.Itoa(1), "城市")
				f.SetCellValue("中标公告", "D"+strconv.Itoa(1), "类型")
				f.SetCellValue("中标公告", "E"+strconv.Itoa(1), "采购单位名称")
				f.SetCellValue("中标公告", "F"+strconv.Itoa(1), "联系人")
				f.SetCellValue("中标公告", "G"+strconv.Itoa(1), "联系电话")
				f.SetCellValue("中标公告", "H"+strconv.Itoa(1), "预算金额")
				f.SetCellValue("中标公告", "I"+strconv.Itoa(1), "中标金额")
				f.SetCellValue("中标公告", "J"+strconv.Itoa(1), "中标单位名称")
				f.SetCellValue("中标公告", "K"+strconv.Itoa(1), "发布时间")
				f.SetCellValue("中标公告", "L"+strconv.Itoa(1), "具体内容")
				f.SetCellValue("中标公告", "A"+strconv.Itoa(win+2), win+1)
				f.SetCellValue("中标公告", "B"+strconv.Itoa(win+2), res[i].Title)
				f.SetCellValue("中标公告", "C"+strconv.Itoa(win+2), res[i].City)
				f.SetCellValue("中标公告", "D"+strconv.Itoa(win+2), res[i].IndustryClassification)
				f.SetCellValue("中标公告", "E"+strconv.Itoa(win+2), res[i].TenderName)
				f.SetCellValue("中标公告", "F"+strconv.Itoa(win+2), res[i].ContactPerson)
				f.SetCellValue("中标公告", "G"+strconv.Itoa(win+2), res[i].ContactInformation)
				f.SetCellValue("中标公告", "H"+strconv.Itoa(win+2), res[i].Amount)
				f.SetCellValue("中标公告", "I"+strconv.Itoa(win+2), res[i].WinName)
				f.SetCellValue("中标公告", "J"+strconv.Itoa(win+2), res[i].BidAmount)
				f.SetCellValue("中标公告", "K"+strconv.Itoa(win+2), res[i].ReleaseTime)
				f.SetCellValue("中标公告", "L"+strconv.Itoa(win+2), res[i].Link)
				win += 1
			case "合同公告":
				f.NewSheet("合同公告")
				f.SetCellValue("合同公告", "A"+strconv.Itoa(1), "序号")
				f.SetCellValue("合同公告", "B"+strconv.Itoa(1), "合同名称")
				f.SetCellValue("合同公告", "C"+strconv.Itoa(1), "甲方")
				f.SetCellValue("合同公告", "D"+strconv.Itoa(1), "甲方联系电话")
				f.SetCellValue("合同公告", "E"+strconv.Itoa(1), "乙方")
				f.SetCellValue("合同公告", "F"+strconv.Itoa(1), "乙方联系电话")
				f.SetCellValue("合同公告", "G"+strconv.Itoa(1), "合同金额")
				f.SetCellValue("合同公告", "H"+strconv.Itoa(1), "合同签订时间")
				f.SetCellValue("合同公告", "I"+strconv.Itoa(1), "合同公告时间")
				f.SetCellValue("合同公告", "J"+strconv.Itoa(1), "具体内容")

				f.SetCellValue("合同公告", "A"+strconv.Itoa(ht+2), ht+1)
				f.SetCellValue("合同公告", "B"+strconv.Itoa(ht+2), res[i].Title)
				f.SetCellValue("合同公告", "C"+strconv.Itoa(ht+2), res[i].TenderName)
				f.SetCellValue("合同公告", "D"+strconv.Itoa(ht+2), res[i].ContactInformation)
				f.SetCellValue("合同公告", "E"+strconv.Itoa(ht+2), res[i].WinName)
				f.SetCellValue("合同公告", "F"+strconv.Itoa(ht+2), res[i].ContactInformation)
				f.SetCellValue("合同公告", "G"+strconv.Itoa(ht+2), res[i].Amount)
				f.SetCellValue("合同公告", "H"+strconv.Itoa(ht+2), res[i].ReleaseTime)
				f.SetCellValue("合同公告", "I"+strconv.Itoa(ht+2), res[i].ReleaseTime)
				f.SetCellValue("合同公告", "J"+strconv.Itoa(ht+2), res[i].Link)
				ht += 1
			case "其他公告":
				f.NewSheet("其他公告")
				f.SetCellValue("其他公告", "A"+strconv.Itoa(1), "序号")
				f.SetCellValue("其他公告", "B"+strconv.Itoa(1), "标题")
				f.SetCellValue("其他公告", "C"+strconv.Itoa(1), "城市")
				f.SetCellValue("其他公告", "D"+strconv.Itoa(1), "类型")
				f.SetCellValue("其他公告", "E"+strconv.Itoa(1), "采购单位名称")
				f.SetCellValue("其他公告", "F"+strconv.Itoa(1), "联系人")
				f.SetCellValue("其他公告", "G"+strconv.Itoa(1), "联系电话")
				f.SetCellValue("其他公告", "H"+strconv.Itoa(1), "具体内容")

				f.SetCellValue("其他公告", "A"+strconv.Itoa(qt+2), qt+1)
				f.SetCellValue("其他公告", "B"+strconv.Itoa(qt+2), res[i].Title)
				f.SetCellValue("其他公告", "C"+strconv.Itoa(qt+2), res[i].City)
				f.SetCellValue("其他公告", "D"+strconv.Itoa(qt+2), res[i].IndustryClassification)
				f.SetCellValue("其他公告", "E"+strconv.Itoa(qt+2), res[i].TenderName)
				f.SetCellValue("其他公告", "F"+strconv.Itoa(qt+2), res[i].ContactPerson)
				f.SetCellValue("其他公告", "G"+strconv.Itoa(qt+2), res[i].ContactInformation)
				f.SetCellValue("其他公告", "H"+strconv.Itoa(qt+2), res[i].Link)
				qt += 1
			}
		}
	}
	fmt.Println("f", f.SheetCount)
	execl_bf, _ := f.WriteToBuffer()
	var buff []byte
	if buff, err = ZipFiles(files, buffs, fmt.Sprintf("/tender/xlsx/%s.xlsx", fmt.Sprintf("%s", gtime.Now())), execl_bf.Bytes()); err != nil {
		log.Println("ZipFiles -------->", err)
		r.Response.Write(&DownRes{Code: 50, Message: "压缩文件失败"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fmt.Sprintf("%s.zip", fmt.Sprintf("%s", gtime.Now())+netUrl.QueryEscape(fmt.Sprintf("%s", title))))
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}

// 企业信息
func (c *DownloadController) DownloadEnterpriseFile(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := []*entity.SysEnterprise{}
	m := system_dao.SysEnterprise.Ctx(r.Context())
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		return
	}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "企业名称")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "城市")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "所属行业")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "联系电话")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "入驻时间")
	files := []string{}
	for i := 0; i < len(res); i++ {
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Name)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Location)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Industry)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Contact)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].CreatedAt.Format("Y-m-d H:i:s.u"))
	}
	execl_bf, _ := f.WriteToBuffer()
	var buff []byte
	if buff, err = ZipFiles(files, nil, "/enterprise/xlsx/enterprise.xlsx", execl_bf.Bytes()); err != nil {
		log.Println(err)
		r.Response.Write(&DownRes{Code: 50, Message: "压缩文件失败"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fmt.Sprintf("%s.zip", fmt.Sprintf("%s", gtime.Now())+netUrl.QueryEscape("-企业信息")))
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}

// 政策咨询
func (c *DownloadController) DownloadConsultationFile(r *ghttp.Request) {
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
			return
		}
	})
	if err != nil {
		return
	}
	files := []string{}
	f := excelize.NewFile()
	f.SetCellStr("sheet1", "A"+strconv.Itoa(1), "标题")
	f.SetCellStr("sheet1", "B"+strconv.Itoa(1), "来源")
	f.SetCellValue("sheet1", "C"+strconv.Itoa(1), "定制机关")
	f.SetCellStr("sheet1", "D"+strconv.Itoa(1), "发布时间")
	f.SetCellValue("sheet1", "E"+strconv.Itoa(1), "内容")
	for i := 0; i < len(res); i++ {
		if res[i].Attachment != "" {
			// todo 政策资讯需要上传minio
			//log.Println(res[i].Attachment)
			//path := strings.Split(res[i].Attachment, "tender")
			//if len(path) > 1 {
			//	// minio 位置
			//	files = append(files, "/mnt/tender"+path[1])
			//}
		}
		curl := ""
		if res[i].Url != "" {
			curl = res[i].Url
		} else {
			curl = fmt.Sprintf("https://www.biaoziku.com/policyServices/%d", res[i].Id)
		}
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Title)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].Url)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].Office)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Publish)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), curl)
	}
	execl_bf, _ := f.WriteToBuffer()
	var buff []byte
	if buff, err = ZipFiles(files, nil, "/consultation/xlsx/consultation.xlsx", execl_bf.Bytes()); err != nil {
		log.Println(err)
		r.Response.Write(&DownRes{Code: 50, Message: "压缩文件失败"})
		return
	}

	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+fmt.Sprintf("%s.zip", fmt.Sprintf("%s", gtime.Now())+netUrl.QueryEscape("-政策资讯")))
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}

// 反馈页附件下载

func (c *DownloadController) DownloadFeedbackFile(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	p := strings.Split(url, "=")
	var ids []string
	if len(p) > 1 {
		ids = strings.Split(p[1], ",")
		log.Println("r.URL.RawQuery", p[1], ids, url)
	}
	res := []*desk_entity.Feedback{}
	m := dao.Feedback.Ctx(r.Context())
	m = m.WhereIn("id", ids)
	err := g.Try(r.Context(), func(ctx context.Context) {
		err := m.Scan(&res)
		if err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
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
	//files := []string{}
	var files []string
	var buffs [][]byte
	for i := 0; i < len(res); i++ {
		if res[i].Attachment != "" {
			for _, url_ := range strings.Split(res[i].Attachment, ",") {
				name, buff := libDownloadAtt.GetAtt(url_)
				files = append(files, name)
				buffs = append(buffs, buff)
			}
		}
		f.SetCellStr("sheet1", "A"+strconv.Itoa(i+2), res[i].Company)
		f.SetCellStr("sheet1", "B"+strconv.Itoa(i+2), res[i].ContactPerson)
		f.SetCellValue("sheet1", "C"+strconv.Itoa(i+2), res[i].ContactInformation)
		f.SetCellStr("sheet1", "D"+strconv.Itoa(i+2), res[i].Remarks)
		f.SetCellValue("sheet1", "E"+strconv.Itoa(i+2), res[i].Attachment)
		f.SetCellValue("sheet1", "F"+strconv.Itoa(i+2), res[i].CreatedAt)
		f.SetCellValue("sheet1", "G"+strconv.Itoa(i+2), res[i].Status)
	}
	execl_bf, _ := f.WriteToBuffer()
	var buff []byte
	if buff, err = ZipFiles(files, buffs, "/feedback/xlsx/feedback.xlsx", execl_bf.Bytes()); err != nil {
		log.Println("ZipFiles -------->", err)
		r.Response.Write(&DownRes{Code: 50, Message: "压缩文件失败"})
		return
	}
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+"feedback"+GenUUID()+".zip")
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}

// 招标单个处理接口
func (c *DownloadController) DownloadSingleTenderFile(r *ghttp.Request) {
	url, _ := gurl.Decode(r.RequestURI)
	directory := ""
	directorySp := strings.Split(url, "directory=")
	if len(directorySp) > 1 {
		directory = strings.Split(directorySp[1], "&")[0]
	}
	name := ""
	nameSp := strings.Split(url, "name=")
	if len(nameSp) > 1 {
		name = strings.Split(nameSp[1], "&")[0]
	}
	buff := qbucket.NewQbucketService("announcement-1317075511").GetObject("/" + directory + "/" + name)
	header := r.Response.Header()
	header.Add("Content-Type", "application/octet-stream")
	header.Add("Content-Disposition", "attachment;filename="+netUrl.QueryEscape(fmt.Sprintf("%s", name)))
	header.Add("Content-Length", fmt.Sprintf("%d", len(buff)))
	header.Add("Access-Control-Expose-Headers", "Content-Disposition")
	r.Response.Write(buff)
}
