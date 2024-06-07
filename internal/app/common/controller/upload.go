package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"tender/internal/app/common/qbucket"
	"tender/internal/app/common/service"
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

func (c *UploadController) UploadFile(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	filetype := form.Value["type"]
	file_type := "default"
	if filetype != nil {
		file_type = filetype[0]
	}
	if v := form.File["file"]; len(v) > 0 {
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&Res{
				Code:    -1,
				Message: "read file failed",
				Url:     "",
			})
			return
		}
		defer file.Close()
		fd, err := ioutil.ReadAll(file)
		//log.Println("cont_type----->", v[0].Header["Content-Type"][0])
		tt := strconv.FormatInt(gtime.Now().UnixMilli(), 10)
		service.S3.S3Upload(fd, file_type, tt+v[0].Filename, v[0].Size)
		// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
		url := "https://biaoziku.com/tender/" + file_type + "/" + tt + v[0].Filename
		r.Response.Write(&Res{Code: 0, Message: "success", Url: url})
		return
	}
}

// 上传调研反馈的附件
func (c *UploadController) UploadResearchFile(r *ghttp.Request) {

	form := r.GetMultipartForm()
	if form == nil {
		r.Response.Write(&Res{Code: -1, Message: "no file", Url: ""})
		return
	}
	// filetype := form.Value["type"]
	file_type := "feedback"
	//if filetype != nil {
	//	file_type = filetype[0]
	//}
	if v := form.File["file"]; len(v) > 0 {
		file, err := v[0].Open()
		if err != nil {
			log.Println(err)
			r.Response.Write(&Res{
				Code:    -1,
				Message: "read file failed",
				Url:     "",
			})
			return
		}
		defer file.Close()
		// fd, err := ioutil.ReadAll(file)
		//log.Println("cont_type----->", v[0].Header["Content-Type"][0])
		filename := v[0].Filename
		filename = strings.Replace(filename, " ", "", -1)
		tt := strconv.FormatInt(gtime.Now().UnixMilli(), 10)
		err = qbucket.NewQbucketService("announcement-1317075511").Upload(file, "/"+file_type, tt+filename, v[0].Size)
		if err != nil {
			log.Println("upload ResearchFile err ", err)
		}
		// service.S3.S3Upload(fd, file_type, strconv.FormatInt(gtime.Now().UnixMilli(), 10)+v[0].Filename, v[0].Size)
		// http://42.193.247.183:9000/tender/icon/0903_udid-nft-037-6159103.zip
		url := "https://biaoziku.com/announcement-1317075511/" + file_type + "/" + tt + filename
		r.Response.Write(&Res{Code: 0, Message: "success", Url: url})
		return
	}
}
