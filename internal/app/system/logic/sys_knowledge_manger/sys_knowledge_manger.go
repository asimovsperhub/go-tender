package sys_knowledge_manger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gen2brain/go-fitz"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	g_udid "github.com/google/uuid"
	"github.com/ledongthuc/pdf"
	"github.com/satori/go.uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"tender/api/v1/system"
	"tender/internal/app/desk/dao"
	"tender/internal/app/desk/model/do"
	desk_entity "tender/internal/app/desk/model/entity"
	"tender/internal/app/system/consts"
	system_dao "tender/internal/app/system/dao"
	system_do "tender/internal/app/system/model/do"
	"tender/internal/app/system/model/entity"
	"tender/internal/app/system/service"
	"tender/internal/packed/websocket"
	"tender/library/liberr"
)

func init() {
	service.RegisterKnowledgeManger(New())
}

func New() *sSysKnowledgeManger {
	return &sSysKnowledgeManger{}
}

type sSysKnowledgeManger struct {
}

func (s *sSysKnowledgeManger) List(ctx context.Context, req *system.KnowledgeSearchReq) (res *system.KnowledgeSearchRes, err error) {
	res = new(system.KnowledgeSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberKnowledge.Ctx(ctx)
	order := "created_at desc"
	if req.Title != "" {
		m = m.Where("title like ?", "%"+req.Title+"%")
	}
	if req.PrimaryClassification != "" {
		m = m.Where("primary_classification = ?", req.PrimaryClassification)
	}
	if req.SecondaryClassification != "" {
		m = m.Where("secondary_classification = ?", req.SecondaryClassification)
	}
	if req.OriginalIndustry != "" {
		m = m.Where("original_type = ?", req.OriginalIndustry)
	}
	if req.Attachment != nil {
		if *req.Attachment == 1 {
			m = m.Where("attachment_url <> ''")
		} else if *req.Attachment == 0 {
			m = m.Where("attachment_url = ''")
		}
	}
	if req.Display != nil {
		m = m.Where("display = ? ", *req.Display)
	}
	if req.Authority != nil {
		m = m.Where("authority = ? ", *req.Authority)
	}
	if req.KnowledgeType != "" {
		m = m.Where("knowledge_type = ? ", req.KnowledgeType)
	}
	if len(req.DateRange) > 0 {
		m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0]+" 00:00:00", req.DateRange[1]+" 23:59:59")
	} else {
		if req.Start != "" {
			m = m.Where("created_at >=?", req.Start+" 00:00:00")
		}
		if req.End != "" {
			m = m.Where("created_at <=?", req.End+" 23:59:59")
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取知库列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.KnowledgeList)
		log.Println("获取知库列表-------->", res.KnowledgeList)
		liberr.ErrIsNil(ctx, err, "获取知库列表数据失败")
	})
	return
}

// ImportList
func (s *sSysKnowledgeManger) ImportList(ctx context.Context, req *system.KnowledgeReq) (res *system.KnowledgeRes, err error) {
	res = new(system.KnowledgeRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberKnowledge.Ctx(ctx)
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	m = m.Where("is_admin = ?", 1)
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取知库导入列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order("created_at desc").Scan(&res.KnowledgeList)
		log.Println("获取知库导入列表-------->", res.KnowledgeList)
		liberr.ErrIsNil(ctx, err, "获取知库导入列表数据失败")
	})
	return
}
func (s *sSysKnowledgeManger) KnowledgeList(ctx context.Context, req *system.KnowledgeReviewSearchReq) (res *system.KnowledgeReviewSearchRes, err error) {
	res = new(system.KnowledgeReviewSearchRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.MemberKnowledge.Ctx(ctx)
	order := "created_at desc"
	if req.KeyWords != "" {
		m = m.Where("title like ?", "%"+req.KeyWords+"%")
	}
	if req.PrimaryClassification != "" {
		m = m.Where("primary_classification = ?", req.PrimaryClassification)
	}
	if req.SecondaryClassification != "" {
		m = m.Where("secondary_classification = ?", req.SecondaryClassification)
	}
	m = m.Where("review_status <> 1")
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取知库审查列表失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.KnowledgeList)
		log.Println("获取知库审查列表-------->", res.KnowledgeList)
		liberr.ErrIsNil(ctx, err, "获取知库审查列表数据失败")
	})
	return
}

func (s *sSysKnowledgeManger) KnowledgeReview(ctx context.Context, req *system.KnowledgeReviewReq) (res *system.KnowledgeReviewRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberKnowledge.Ctx(ctx).TX(tx).WherePri(req.Id).Update(do.MemberKnowledge{
				ReviewStatus:    req.Status,
				UpdatedAt:       gtime.Now(),
				OpreviewMessage: req.OpReviewMessage,
				// 操作管理员id
				// OperationId: service.Context().GetUserId(ctx),
			})
			liberr.ErrIsNil(ctx, err, "审核知识失败")
		})
		return err
	})
	Knowledge := (*desk_entity.MemberKnowledge)(nil)
	m := dao.MemberKnowledge.Ctx(ctx)
	m = m.Where("id = ?", req.Id)
	err = g.Try(ctx, func(ctx context.Context) {
		err = m.Scan(&Knowledge)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
	})
	switch req.Status {
	case 1:
		req.OpReviewMessage = "知识审核通过: " + req.OpReviewMessage
		userinfo := (*entity.MemberUser)(nil)
		user_m := system_dao.MemberUser.Ctx(ctx)
		err = g.Try(ctx, func(ctx context.Context) {
			user_m = user_m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, Knowledge.UserId))
			err = user_m.Limit(1).Scan(&userinfo)
		})
		setting := (*entity.MemberIntegral)(nil)
		err = g.Try(ctx, func(ctx context.Context) {
			err = system_dao.MemberIntegral.Ctx(ctx).Scan(&setting)
			liberr.ErrIsNil(ctx, err, "获取积分信息失败")
		})

		if Knowledge.VideoUrl != "" {
			err = g.Try(ctx, func(ctx context.Context) {
				user_m = user_m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, Knowledge.UserId))
				_, err = user_m.Update(system_do.MemberUser{Integral: userinfo.Integral + setting.VideoIntegral})
				liberr.ErrIsNil(ctx, err, "更新用户积分失败")
			})
		} else {
			err = g.Try(ctx, func(ctx context.Context) {
				user_m = user_m.Where(fmt.Sprintf("%s='%d'", system_dao.MemberUser.Columns().Id, Knowledge.UserId))
				_, err = user_m.Update(system_do.MemberUser{Integral: userinfo.Integral + setting.KnowledgeIntegral})
				liberr.ErrIsNil(ctx, err, "更新用户积分失败")
			})
		}
	case 2:
		req.OpReviewMessage = "知识审核未通过: " + req.OpReviewMessage
	}
	if req.Status != 0 {
		msgId := g_udid.New().String()
		_, err = system_dao.SysWsMsg.Ctx(ctx).Insert(system_do.SysWsMsg{
			MessageId: msgId,
			UserId:    Knowledge.UserId,
			Content:   req.OpReviewMessage,
			IsRead:    0,
			IsDel:     0,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		})
		if err != nil {
			return nil, fmt.Errorf("插入消息失败: %w", err)
		}

		websocket.SendToUser(uint64(Knowledge.UserId), &websocket.WResponse{
			Event: "context",
			Data: map[string]string{
				"messageId": msgId,
				"content":   req.OpReviewMessage,
			},
		})
	}
	return
}

func (s *sSysKnowledgeManger) Del(ctx context.Context, req *system.KnowledgeDelReq) (res *system.KnowledgeDelRes, err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.MemberKnowledge.Ctx(ctx).TX(tx).WherePri(req.Id).Delete()
			liberr.ErrIsNil(ctx, err, "删除知识失败")
		})
		return err
	})
	return
}

func (s *sSysKnowledgeManger) ProcessVideo(ctx context.Context, req *system.KnowledgeProcessVideoReq) (res *system.KnowledgeProcessVideoRes, err error) {
	err = ffmpeg.Input("./sample_data/test.mp4", ffmpeg.KwArgs{"ss": 1}).
		Output("./sample_data/out1.mp4", ffmpeg.KwArgs{"t": 15}).OverWriteOutput().Run()

	reader := ExampleReadFrameAsJpeg("./sample_data/test.mp4", 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}
	err = imaging.Save(img, "./sample_data/out1.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	return
}

func in(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func CutVideo(inputVideoPath, outputDir, startTime string, duration int) string {
	_formatArr := []string{"mp4", "flv"}
	_, _file := filepath.Split(inputVideoPath)
	_tmps := strings.Split(_file, ".")
	_ext := _tmps[len(_tmps)-1]
	if !in(_ext, _formatArr) {
		fmt.Println("格式不支持")
	}
	_name := uuid.NewV4()

	_resultVideoPath := filepath.Join(outputDir, fmt.Sprintf("%s.%s", _name.String(), _ext))
	err := ffmpeg.Input(inputVideoPath).
		Output(_resultVideoPath, ffmpeg.KwArgs{"ss": startTime, "t": duration, "c:v": "copy", "c:a": "copy"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Println(err)
	}
	return _resultVideoPath
}

func (s *sSysKnowledgeManger) ProcessPdf(ctx context.Context, req *system.KnowledgeProcessPdfReq) (res *system.KnowledgeProcessPdfRes, err error) {
	pdf.DebugOn = true
	content, err := readPdf("./sample_data/consultation.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	ParsePdf()

	//result, err := ConvertWordToPdf("./sample_data/doctest.docx", "./sample_data")
	//fmt.Printf(result)

	//docs := search.DocItems{{
	//	"id":      1,
	//	"user_id": 1,
	//	"content": content,
	//}}
	//_, err = commonService.TS.AddDocuments(docs, fmt.Sprintf("%d", 1))
	//if err != nil {
	//	fmt.Println(err)
	//}
	return
}

func ParsePdf() {
	doc, err := fitz.New("./sample_data/consultation.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir("./sample_data", "fitz")
	if err != nil {
		panic(err)
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", n)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(text)
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as html

	f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.html", -1)))
	if err != nil {
		panic(err)
	}

	f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<style>\nbody{background-color:slategray}\ndiv{position:relative;background-color:white;margin:1em auto;box-shadow:1px 1px 8px -2px black}\np{position:absolute;white-space:pre;margin:0}\n</style>\n</head>\n<body>")
	for n := 0; n < doc.NumPage(); n++ {
		html, err := doc.HTML(n, false)
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(html)
		if err != nil {
			panic(err)
		}

	}
	f.WriteString("</body>\n</html>")
	f.Close()
}
func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ConvertWordToPdf(serverFile, workDir string) (string, error) {
	cmd := exec.Command("soffice", "--headless", "--invisible", "--convert-to", "pdf", serverFile, "--outdir", workDir)
	data, err := cmd.Output()
	if err != nil {
		fmt.Println("convert failed: ", err)
		return "", err
	}

	return string(data), nil
}
