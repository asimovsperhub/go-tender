package libmessage

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	time "github.com/alibabacloud-go/darabonba-time/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"log"
	"math/rand"
	"testing"
	stime "time"
)

func Test_message(t *testing.T) {
	PhoneNumbers := "15399225820"
	SignName := "云鹤网络科技"
	TemplateCode := "SMS_259610319"
	TemplateParam := "{'code':'123456'}"
	err := _main([]*string{&PhoneNumbers, &SignName, &TemplateCode, &TemplateParam})
	if err != nil {
		panic(err)
	}
}

// 使用AK&SK初始化账号Client
func CreateClient_() (_result *dysmsapi.Client, _err error) {
	accessKeyId := "LTAI5t99st1FjpHYn88EDuNU"
	accessKeySecret := "l2N8p1nOi8Takz1mwcERdTdU0o29qf"
	config := &openapi.Config{}
	config.AccessKeyId = &accessKeyId
	config.AccessKeySecret = &accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient_()
	if _err != nil {
		return _err
	}
	// 1.发送短信
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  args[0],
		SignName:      args[1],
		TemplateCode:  args[2],
		TemplateParam: args[3],
	}
	sendResp, _err := client.SendSms(sendReq)
	if _err != nil {
		log.Println("err----------------->", _err)
		return _err
	}
	log.Println("sendResp------------------>", sendResp)
	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		return _err
	}

	bizId := sendResp.Body.BizId
	// 2. 等待 10 秒后查询结果
	_err = util.Sleep(tea.Int(10000))
	if _err != nil {
		return _err
	}
	// 3.查询结果
	phoneNums := string_.Split(args[0], tea.String(","), tea.Int(-1))
	for _, phoneNum := range phoneNums {
		queryReq := &dysmsapi.QuerySendDetailsRequest{
			PhoneNumber: util.AssertAsString(phoneNum),
			BizId:       bizId,
			SendDate:    time.Format(tea.String("yyyyMMdd")),
			PageSize:    tea.Int64(10),
			CurrentPage: tea.Int64(1),
		}
		queryResp, _err := client.QuerySendDetails(queryReq)
		if _err != nil {
			return _err
		}

		dtos := queryResp.Body.SmsSendDetailDTOs.SmsSendDetailDTO
		// 打印结果
		for _, dto := range dtos {
			if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("3"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送成功，接收时间: " + tea.StringValue(dto.ReceiveDate)))
			} else if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("2"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送失败"))
			} else {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 正在发送中..."))
			}

		}
	}
	return _err
}

func Test_crate_code(t *testing.T) {
	log.Println(fmt.Sprintf("%06v", rand.New(rand.NewSource(stime.Now().UnixNano())).Int31n(1000000)))

}

func Test(t *testing.T) {
	// c := GetEntity()

}
