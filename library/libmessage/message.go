package libmessage

import (
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/patrickmn/go-cache"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Message struct {
	verificationCodeCache    *cache.Cache
	verificationCodeReqCache *cache.Cache
	client                   *dysmsapi.Client
}

var (
	Once   sync.Once
	Entity *Message
)

func GetEntity() *Message {
	Once.Do(func() {
		Entity = new(Message)
		// 两分钟内只能发送一次验证码
		Entity.verificationCodeReqCache = cache.New(time.Minute*2, time.Minute*2)
		// 验证码 15 分钟过期
		Entity.verificationCodeCache = cache.New(time.Minute*15, time.Minute*15)
		client, _err := CreateClient()
		if _err != nil {
			log.Println(_err)
		}
		Entity.client = client
	})
	return Entity
}

// 使用AK&SK初始化账号Client
func CreateClient() (_result *dysmsapi.Client, _err error) {
	accessKeyId := "LTAI5t99st1FjpHYn88EDuNU"
	accessKeySecret := "l2N8p1nOi8Takz1mwcERdTdU0o29qf"
	config := &openapi.Config{}
	config.AccessKeyId = &accessKeyId
	config.AccessKeySecret = &accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

func (m *Message) SendCode(Phone string) error {
	PhoneNumbers := Phone
	SignName := "云鹤网络科技"
	TemplateCode := "SMS_259610319"
	Code := m.RandCode()
	TemplateParam := fmt.Sprintf("{'code':'%s'}", Code)
	_, found := m.verificationCodeReqCache.Get(Phone)
	if found {
		err := errors.New("请勿重复发送验证码")
		return err
	}
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers: &PhoneNumbers,
		SignName:     &SignName,
		TemplateCode: &TemplateCode,
		// 模版内容有参数的话必须带上这个 一般就是$中的内容
		TemplateParam: &TemplateParam,
	}

	m.verificationCodeReqCache.SetDefault(Phone, 1)
	m.verificationCodeCache.SetDefault(Phone, Code)
	sendResp, _err := m.client.SendSms(sendReq)
	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		//if !g.IsNil(_err) {
		//	panic(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		//	return _err
		//}
		log.Println("sendResp ------------------------->", _err)
		if tea.BoolValue(util.EqualString(code, tea.String("MOBILE_SEND_LIMIT"))) ||
			tea.BoolValue(util.EqualString(code, tea.String("isv.BUSINESS_LIMIT_CONTROL"))) {
			return errors.New("请求过于频繁, 请稍后再试")
		}
		return errors.New(tea.StringValue(sendResp.Body.Message))
	}

	return nil
}
func (m *Message) CheckVerificationCode(Phone, Code string) (err error) {
	cacheCode, found := m.verificationCodeCache.Get(Phone)
	if !found {
		err = errors.New("验证码已失效")
		return
	}
	cc, sure := cacheCode.(string)
	if !sure {
		err = errors.New("内部服务出错")
		return
	}
	if cc != Code {
		err = errors.New("验证码输入错误")
		return
	} else {
		m.verificationCodeCache.Delete(Phone) //验证码验证成功后删除
	}
	return
}
func (m *Message) RandCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
