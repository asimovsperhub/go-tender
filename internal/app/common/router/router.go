package router

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"tender/internal/app/desk/service"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/entity"

	"tender/internal/app/common/controller"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Bind(
		controller.Msg.SendWs,
	)
	group.POST("/upload", controller.Upload.UploadFile)
	// UploadResearchFile
	group.POST("/upload/feedback", controller.Upload.UploadResearchFile)
	group.ALL("/wx_event", func(r *ghttp.Request) {
		if r.Method == http.MethodGet {
			verifyToken(r)
		} else if r.Method == http.MethodPost {
			handleEvent(r)
		}

	})
	// DownloadConsultationFile
	group.GET("/consultation/download", controller.Download.DownloadConsultationFile)
	//DownloadConsultationFilePdf
	group.GET("/consultation/download/pdf", controller.Download.DownloadConsultationFilePdf)
	group.GET("/consultation/download/word", controller.Download.DownloadConsultationFileWord)

	// Download
	//登录验证拦截
	service.GfToken().Middleware(group)
	//context拦截器 + 权限拦截器 用户对应的菜单
	group.Middleware(service.Middleware().Ctx)

	// 招标批量下载
	group.GET("/tender/download", controller.Download.DownloadFile)
	//DownloadFilePdf
	group.GET("/tender/download/pdf", controller.Download.DownloadFilePdf)
	// DownloadFileWord
	group.GET("/tender/download/word", controller.Download.DownloadFileWord)

	// 知识下载
	group.GET("/knowledge/download", controller.Download.DownloadKnowledgeFile)
	//DownloadKnowledgeFilePdf
	group.GET("/knowledge/download/pdf", controller.Download.DownloadKnowledgeFilePdf)
	//DownloadKnowledgeFileWord
	group.GET("/knowledge/download/word", controller.Download.DownloadKnowledgeFileWord)
}

func verifyToken(r *ghttp.Request) {

	token := g.Cfg().MustGet(r.Context(), "mp.token").String()
	signature := r.GetQuery("signature").String()
	timestamp := r.GetQuery("timestamp").String()
	nonce := r.GetQuery("nonce").String()
	echostr := r.GetQuery("echostr").String()

	// 将参数按字典序排序
	strSlice := []string{token, timestamp, nonce}
	sort.Strings(strSlice)
	str := strings.Join(strSlice, "")

	// 对排序后的参数进行sha1签名
	h := sha1.New()
	h.Write([]byte(str))
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	// 将签名结果与微信传递的签名进行比较
	if sha1Hash == signature {
		r.Response.Write(echostr) // 验证成功，返回echostr给微信服务器
	} else {
		r.Response.Write("Invalid signature") // 验证失败
	}
}

// 微信消息结构体
type WechatMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Event        string
	EventKey     string
}

func handleEvent(r *ghttp.Request) {
	// 读取HTTP请求的Body数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		r.Response.Write("Read body error: " + err.Error())
		return
	}
	// 解析XML消息体
	var msg WechatMessage
	err = xml.Unmarshal(body, &msg)
	if err != nil {
		r.Response.Write("Parse body error: " + err.Error())
		return
	}

	switch msg.MsgType {
	case "event":
		switch msg.Event {
		case "subscribe":
			// 处理用户订阅事件
			fmt.Println("User subscribed: ", msg.FromUserName)
			finduser := &entity.SysWxUser{}
			wsUserDao := dao.SysWxUser.Ctx(r.Context())
			if err = wsUserDao.Where(dao.SysWxUser.Columns().OpenId, msg.FromUserName).Scan(finduser); err != nil {
				g.Log().Errorf(r.Context(), "func:handleEvent,查询用户失败：%s", err)
			}
			// 如果用户不存在，则创建用户
			if errors.Is(err, sql.ErrNoRows) {
				user := &entity.SysWxUser{
					OpenId:    msg.FromUserName,
					CreatedAt: gtime.Now(),
					UpdatedAt: gtime.Now(),
				}
				if _, err = wsUserDao.Data(user).Insert(); err != nil {
					fmt.Println(err)
				}
			}
			// 处理扫码携带的sence_id
			if msg.EventKey != "" {
				events := strings.Split(msg.EventKey, "_")
				if len(events) == 2 {
					if err := scanHandle(r.Context(), events[1], msg.FromUserName); err != nil {
						g.Log().Errorf(r.Context(), "func:handleEvent,处理扫码携带的sence_id失败：%s", err)
						return
					}
				}
			}

		case "unsubscribe":
			// 处理用户取消订阅事件
		case "CLICK":
			// 处理菜单点击事件
		case "SCAN":

			finduser := &entity.SysWxUser{}
			wsUserDao := dao.SysWxUser.Ctx(r.Context())
			if err = wsUserDao.Where(dao.SysWxUser.Columns().OpenId, msg.FromUserName).Scan(finduser); err != nil {
				g.Log().Errorf(r.Context(), "func:handleEvent,查询用户失败：%s", err)
			}
			// 如果用户不存在，则创建用户
			if errors.Is(err, sql.ErrNoRows) {
				user := &entity.SysWxUser{
					OpenId:    msg.FromUserName,
					CreatedAt: gtime.Now(),
					UpdatedAt: gtime.Now(),
				}
				if _, err = wsUserDao.Data(user).Insert(); err != nil {
					fmt.Println(err)
				}
			}

			// 处理用户扫描带参数二维码事件
			if err := scanHandle(r.Context(), msg.EventKey, msg.FromUserName); err != nil {
				g.Log().Errorf(r.Context(), "func:handleEvent,处理用户扫描带参数二维码事件失败：%s", err)
				return
			}
			reply := &Message{
				ToUserName:   msg.FromUserName,
				FromUserName: msg.ToUserName,
				CreateTime:   int64(msg.CreateTime),
				MsgType:      "text",
				Content:      "扫码成功",
			}
			xmlReply, err := xml.Marshal(reply)
			if err != nil {
				r.Response.Write("Failed to marshal XML")
				return
			}
			r.Response.Header().Set("Content-Type", "application/xml")
			r.Response.Write(xmlReply)
		default:
			// 其他事件消息
			fmt.Println("Unknown event: ", msg.Event)
		}
	case "text":

		reply := &Message{
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   int64(msg.CreateTime),
			MsgType:      "text",
			Content:      "感谢您联系我们的公众号",
		}
		xmlReply, err := xml.Marshal(reply)
		if err != nil {
			r.Response.Write("Failed to marshal XML")
			return
		}

		r.Response.Header().Set("Content-Type", "application/xml")
		r.Response.Write(xmlReply)
	}
}

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

func scanHandle(ctx context.Context, senceId string, openId string) error {
	var token string
	// 查询是否有未扫描的二维码记录是否存在
	if count, _ := dao.SysWxQrcode.Ctx(ctx).Where("sence_id = ? AND is_scan = ?", senceId, 0).Count(); count == 1 {

		// 根据openid查询wxuser
		wsUserDao := dao.SysWxUser.Ctx(ctx)
		finduser := &entity.SysWxUser{}
		if err := wsUserDao.Where(dao.SysWxUser.Columns().OpenId, openId).Scan(finduser); err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("func:scanHandle,查询用户失败：%s", err)
		} else {
			// 根据wxuser mainid 查询用户信息
			user := &entity.MemberUser{}
			if err := dao.MemberUser.Ctx(ctx).Where(dao.MemberUser.Columns().Id, finduser.MainId).Scan(user); err != nil {
				g.Log().Errorf(ctx, "func:scanHandle,查询用户失败：%s", err)
				token, _ = service.GfToken().GenerateToken(ctx, gconv.String(openId)+"-"+gmd5.MustEncryptString(openId), nil)
			} else {
				// 如果用户表中存在该用户，则生成对应的token
				key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
				// 不能返回用户密码
				user.UserPassword = ""
				token, err = service.GfToken().GenerateToken(ctx, key, user)
				if err != nil {
					g.Log().Errorf(ctx, "func:scanHandle,生成token失败：%s", err)
				}
			}
		}

		wsQrcodeDao := dao.SysWxQrcode.Ctx(ctx)
		if rs, err := wsQrcodeDao.Where("sence_id = ? AND is_scan = ?", senceId, 0).Update(map[string]interface{}{
			"open_id": openId,
			"is_scan": "1",
			"token":   token,
		}); err != nil {
			return fmt.Errorf("func:scanHandle,更新二维码记录失败：%s", err)
		} else {
			if rows, _ := rs.RowsAffected(); rows == 0 {
				g.Log().Warningf(ctx, "func:scanHandle,更新二维码记录失败，未找到记录,senceId:%s,oepnId:%s", senceId, openId)
			}
		}
	} else {
		g.Log().Warningf(ctx, "func:scanHandle,未找到未扫描的二维码记录,senceId:%s,oepnId:%s", senceId, openId)
	}
	return nil
}
