package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	uuid "github.com/satori/go.uuid"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"tender/internal/app/common/search"
	commonService "tender/internal/app/common/service"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/model/do"
	system_dao "tender/internal/app/system/dao"
	"tender/internal/app/system/library/libPdfUtils"
	system_entity "tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/library/liberr"
)

var (
	Upload = UploadController{}
)

type UploadController struct {
	BaseController
}

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

func UploadFile(srcPath string) (string, error) {
	fd, err := os.ReadFile(srcPath)
	if err != nil {
		return "", err
	}

	filetype := gfile.ExtName(srcPath)
	fileName := gfile.Basename(srcPath)

	fileInfo, err := os.Stat(srcPath)
	commonService.S3.S3Upload(fd, filetype, fileName, fileInfo.Size())
	// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
	url := "https://biaoziku.com/tender/" + filetype + "/" + fileName
	return url, nil
}

func GetIntegral(ctx context.Context, intype int, integral int) (err error) {
	res := (*system_entity.MemberIntegral)(nil)
	err = g.Try(ctx, func(ctx context.Context) {
		err = system_dao.MemberIntegral.Ctx(ctx).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取积分信息失败")
	})
	if err != nil {
		return
	}
	if res != nil {
		if intype == 0 {
			if integral > res.Ordinary {
				err = errors.New("积分超过上限")
				return
			}
		} else {
			if integral > res.Select {
				err = errors.New("积分超过上限")
				return
			}
		}
	}
	return
}

// GenUUID 生成一个随机的唯一ID
func GenUUID() string {
	return uuid.NewV4().String()
}

func (c *UploadController) SystemImportFileVideo(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	if v := form.File["file"]; len(v) > 0 {
		// 但文件上传
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&DownRes{
				Code:    -1,
				Message: "read file failed",
			})
			return
		}
		defer file.Close()
		fileName := v[0].Filename
		savePath := "/tmp/" + GenUUID() + gfile.Basename(fileName)
		distFile, err := gfile.Create(savePath)
		if err != nil {
			r.Response.Write(err)
			return
		}
		if _, err := io.Copy(distFile, file); err != nil {
			r.Response.Write(err)
			return
		}
		// todo 异步切帧
		//vp, cp, err := libVideoUtils.ProcessVideo(savePath, "/tmp")
		//shortUrl, err := UploadFile(vp)
		//fmt.Println("short url: ", shortUrl)
		//coverUrl, err := UploadFile(cp)
		//if err != nil {
		//	r.Response.Write(err)
		//	return
		//}
		filetype := form.Value["type"]
		file_type := "default"
		if filetype != nil {
			file_type = filetype[0]
		}
		fd, err := os.ReadFile(savePath)
		if err != nil {
			fmt.Println(err)
		}
		commonService.S3.S3Upload(fd, file_type, v[0].Filename, v[0].Size)
		url := "https://biaoziku.com/tender/" + file_type + "/" + v[0].Filename

		integral := 0
		authority := 1
		id, err := dao.MemberKnowledge.Ctx(r.Context()).InsertAndGetId(do.MemberKnowledge{
			Title:                   fileName,
			KnowledgeType:           "普通",
			Authority:               authority,
			PrimaryClassification:   "视频课程",
			SecondaryClassification: "",
			IntegralSetting:         integral,
			Content:                 "",
			ReviewMessage:           "",
			// ReviewMessage:           req.ReviewMessage,
			ReviewStatus: 1,
			// 0:其他,1:视频
			Type:     1,
			IsAdmin:  1,
			VideoUrl: url,
			//CoverUrl:      coverUrl,
			//ShortvideoUrl: shortUrl,
			Display:   1,
			Crawler:   0,
			UserId:    service.Context().GetUserId(r.Context()),
			UserName:  service.Context().GetLoginUser(r.Context()).UserName,
			CreatedAt: gtime.Now(),
		})
		res, err := g.Redis().Do(r.Context(), "RPUSH", "cutting", fmt.Sprintf("%d_%s", id, savePath))
		if err != nil {
			fmt.Println("redis-err---------->", err, fmt.Sprintf("%d_%s", id, savePath))
		}
		fmt.Println("redis---------->", res.String())
		r.Response.Write(&DownRes{Code: 0, Message: "success"})
		return
	}
}

func (c *UploadController) SystemImportFilePdf(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	if v := form.File["file"]; len(v) > 0 {
		// 但文件上传
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&DownRes{
				Code:    -1,
				Message: "read file failed",
			})
			return
		}
		fileName := v[0].Filename
		savePath := "/tmp/" + GenUUID() + gfile.Basename(fileName)
		distFile, err := gfile.Create(savePath)
		if err != nil {
			r.Response.Write(err)
			return
		}
		if _, err := io.Copy(distFile, file); err != nil {
			r.Response.Write(err)
			return
		}

		htmlPath, content, err := libPdfUtils.ParsePdf(savePath)
		if err != nil {
			r.Response.Write(err)
			return
		}
		htmlContent, err := libPdfUtils.ReadContent(htmlPath)

		defer file.Close()
		filetype := form.Value["type"]
		file_type := "default"
		if filetype != nil {
			file_type = filetype[0]
		}
		fd, err := os.ReadFile(savePath)
		if err != nil {
			fmt.Println(err)
		}
		commonService.S3.S3Upload(fd, file_type, v[0].Filename, v[0].Size)
		// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
		url := "https://biaoziku.com/tender/" + file_type + "/" + v[0].Filename

		integral := 0
		authority := 1
		knowledgeId, err := dao.MemberKnowledge.Ctx(r.Context()).InsertAndGetId(do.MemberKnowledge{
			Title:                   fileName,
			KnowledgeType:           "普通",
			Authority:               authority,
			PrimaryClassification:   "",
			SecondaryClassification: "",
			IntegralSetting:         integral,
			Content:                 htmlContent,
			// ReviewMessage:           req.ReviewMessage,
			ReviewStatus: 1,
			// 0:其他,1:视频
			Type:          0,
			IsAdmin:       1,
			AttachmentUrl: url,
			ReviewMessage: "",
			Display:       1,
			Crawler:       0,
			UserId:        service.Context().GetUserId(r.Context()),
			UserName:      service.Context().GetLoginUser(r.Context()).UserName,
			CreatedAt:     gtime.Now(),
		})

		docs := search.DocItems{{
			"id":      knowledgeId,
			"user_id": service.Context().GetUserId(r.Context()),
			"content": content,
		}}
		_, err = commonService.TS.AddDocuments(docs, fmt.Sprintf("%d", knowledgeId))
		if err != nil {
			fmt.Println(err)
		}

		r.Response.Write(&DownRes{Code: 0, Message: "success"})
		return
	}
}

func (c *UploadController) SystemImportFileWord(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	if v := form.File["file"]; len(v) > 0 {
		// 但文件上传
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&DownRes{
				Code:    -1,
				Message: "read file failed",
			})
			return
		}
		fileName := v[0].Filename
		newFileName := GenUUID() + gfile.Basename(fileName)
		savePath := "/tmp/" + newFileName
		distFile, err := gfile.Create(savePath)
		if err != nil {
			r.Response.Write(err)
			return
		}
		if _, err := io.Copy(distFile, file); err != nil {
			r.Response.Write(err)
			return
		}

		result, err := libPdfUtils.ConvertWordToPdf(savePath, "/tmp")
		fmt.Printf(result)

		pdfPath := "/tmp/" + gfile.Name(newFileName) + ".pdf"
		htmlPath, content, err := libPdfUtils.ParsePdf(pdfPath)
		htmlContent, err := libPdfUtils.ReadContent(htmlPath)
		defer file.Close()
		filetype := form.Value["type"]
		file_type := "default"
		if filetype != nil {
			file_type = filetype[0]
		}
		fd, err := os.ReadFile(savePath)
		if err != nil {
			fmt.Println(err)
		}
		commonService.S3.S3Upload(fd, file_type, v[0].Filename, v[0].Size)
		// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
		url := "https://biaoziku.com/tender/" + file_type + "/" + v[0].Filename

		integral := 0
		authority := 1
		knowledgeId, err := dao.MemberKnowledge.Ctx(r.Context()).InsertAndGetId(do.MemberKnowledge{
			Title:                   fileName,
			KnowledgeType:           "普通",
			Authority:               authority,
			PrimaryClassification:   "",
			SecondaryClassification: "",
			IntegralSetting:         integral,
			Content:                 htmlContent,
			ReviewStatus:            1,
			Type:                    0,
			IsAdmin:                 1,
			AttachmentUrl:           url,
			ReviewMessage:           "",
			Display:                 1,
			Crawler:                 0,
			UserId:                  service.Context().GetUserId(r.Context()),
			UserName:                service.Context().GetLoginUser(r.Context()).UserName,
			CreatedAt:               gtime.Now(),
		})

		docs := search.DocItems{{
			"id":      knowledgeId,
			"user_id": service.Context().GetUserId(r.Context()),
			"content": content,
		}}
		_, err = commonService.TS.AddDocuments(docs, fmt.Sprintf("%d", knowledgeId))
		if err != nil {
			fmt.Println(err)
		}

		r.Response.Write(&DownRes{Code: 0, Message: "success"})
		return
	}
}

// excel 导入知识
func (c *UploadController) SystemImportFileExcel(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	if v := form.File["file"]; len(v) > 0 {
		// 但文件上传
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&DownRes{
				Code:    -1,
				Message: "read file failed",
			})
			return
		}
		defer file.Close()
		f, err := excelize.OpenReader(file)
		if err != nil {
			file.Close()
		}
		f.Path = v[0].Filename
		defer func() {
			if err = f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		// 获取 Sheet1 上所有单元格
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			r.Response.Write(&DownRes{Code: 50, Message: "读取 Sheet1 失败"})
			return
		}
		if len(rows) > 1 {
			for _, row := range rows[1:] {
				//// 标题
				//fmt.Print(row[0], "\t")
				//// 知识类型
				//fmt.Print(row[1], "\t")
				//// 一级分类
				//fmt.Print(row[2], "\t")
				//// 二级分类
				//fmt.Print(row[3], "\t")
				//// 阅读下载权限
				//fmt.Print(row[4], "\t")
				//// 积分设置
				//fmt.Print(row[5], "\t")
				//// 知识内容
				//fmt.Print(row[6], "\t")
				integral_type := 0
				authority := 0
				if row[1] == "普通" {
					integral_type = 0
				} else {
					integral_type = 1
				}
				if row[4] == "所有人" {
					authority = 0
				} else {
					authority = 1
				}
				integral, _ := strconv.Atoi(row[5])
				err = GetIntegral(r.Context(), integral_type, integral)
				if err != nil {
					r.Response.Write(&DownRes{Code: 50, Message: "积分超过上限"})
					return
				}
				err = g.DB().Transaction(r.Context(), func(ctx context.Context, tx gdb.TX) error {
					err = g.Try(ctx, func(ctx context.Context) {
						knowledgeId, e := dao.MemberKnowledge.Ctx(ctx).TX(tx).InsertAndGetId(do.MemberKnowledge{
							Title:                   row[0],
							KnowledgeType:           row[1],
							Authority:               authority,
							PrimaryClassification:   row[2],
							SecondaryClassification: row[3],
							IntegralSetting:         integral,
							Content:                 row[6],
							ReviewMessage:           "",
							// ReviewMessage:           req.ReviewMessage,
							ReviewStatus: 1,
							// 0:其他,1:视频
							Type:      0,
							IsAdmin:   1,
							Display:   1,
							Crawler:   0,
							UserId:    service.Context().GetUserId(ctx),
							UserName:  service.Context().GetLoginUser(ctx).UserName,
							CreatedAt: gtime.Now(),
						})
						if e != nil {
							r.Response.Write(&DownRes{Code: 50, Message: "插入数据失败"})
							return
						}
						docs := search.DocItems{{
							"id":      knowledgeId,
							"user_id": service.Context().GetUserId(r.Context()),
							"content": row[6],
						}}
						_, err = commonService.TS.AddDocuments(docs, fmt.Sprintf("%d", knowledgeId))
						if err != nil {
							fmt.Println(err)
						}
					})
					return err
				})
			}
		}
		r.Response.Write(&DownRes{Code: 0, Message: "success"})
		return
	}
}

// 其他文件导入
func (c *UploadController) SystemImportFileOther(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	if v := form.File["file"]; len(v) > 0 {
		// 但文件上传
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&DownRes{
				Code:    -1,
				Message: "read file failed",
			})
			return
		}
		defer file.Close()
		filetype := form.Value["type"]
		file_type := "default"
		if filetype != nil {
			file_type = filetype[0]
		}
		fd, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		commonService.S3.S3Upload(fd, file_type, v[0].Filename, v[0].Size)
		// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
		url := "https://biaoziku.com/tender/" + file_type + "/" + v[0].Filename

		integral := 0
		authority := 1
		_, err = dao.MemberKnowledge.Ctx(r.Context()).Insert(do.MemberKnowledge{
			Title:                   v[0].Filename,
			KnowledgeType:           "普通",
			Authority:               authority,
			PrimaryClassification:   "",
			SecondaryClassification: "",
			IntegralSetting:         integral,
			Content:                 "其他文件上传",
			ReviewStatus:            1,
			Type:                    0,
			IsAdmin:                 1,
			AttachmentUrl:           url,
			ReviewMessage:           "",
			Display:                 1,
			Crawler:                 0,
			UserId:                  service.Context().GetUserId(r.Context()),
			UserName:                service.Context().GetLoginUser(r.Context()).UserName,
			CreatedAt:               gtime.Now(),
		})
		r.Response.Write(&DownRes{Code: 0, Message: "success"})
		return
	}
}
