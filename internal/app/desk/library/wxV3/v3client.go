package wxV3

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
	"log"
	"sync"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/entity"
)

var wxclient *wechat.ClientV3
var once sync.Once

func GetClient(ctx context.Context) *wechat.ClientV3 {
	if wxclient == nil {
		once.Do(func() {
			InitWxV3Client(ctx)
		})
		return wxclient
	}
	return wxclient
}

func InitWxV3Client(ctx context.Context) {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	//PrivateKey := `
	//-----BEGIN PRIVATE KEY-----
	//MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDXWyFrXm5/SWs1
	//TZOiUVnwoiiQKHnvn555yqCMv/vcNu4oRRgr5ywy/VYXf/yLHRDCLwelj1+Y7EG7
	//5F/+NKWd141jXYEWdEkgV0QgRyMnDXYII6xverWcm76Q0WghBzoxKJdvGBlNjJdh
	//VNC4djm9zPc8oMV7yiX+zVeWCXnx7FxrQ+biG5R2s4T6L0rJugAG4Q+4+vv3X+WD
	//CBWRfr3LG229JZXDYKmamne5vRoMczb2jzwgVeLNblv2gJHvfIJ1lEuMvb4fE8CX
	//X0X7d5VF94J/04fUCzZI9eyS+BgN+aKpIJPgMddbSn2JRgLBwBd5GZ3Wv92h+fV2
	//+W+5sJRfAgMBAAECggEAdtbGIV7R8RHNxHNTxd3ImP6TDcIYT8Awjprfb+w9iu8R
	//C03dodSG0nh0OrGY5hea+N1FBfqRUW0GdS34PupEfk7FqhCePIrdE6i0Ym9/GXzX
	//JnSxIi9/6rUTOw0jvh4j4Z65ukd9JwsvOjm4mCI9iNyDjoRVlhMeVnZwH86I6ZN2
	//EvmnuftpUMJ/v+IsGsFzQ36Gs47zeJoZM2IGpcI3UOUGD4/0+VWgYSbmXPD1giMh
	//uH9d5GQqj/r2zu7QK7tRsWk8ELwuXI17Ob2mfvLB1S3Pg3TiMrWgdd0FZ+qp5Bor
	//2uc8ABTZLIhax7iqptH4AQ8kNFUZE//aUBnwgWrykQKBgQD+NVJczqTvTjo7SQth
	//5NYop0vDSB8dZiId5pXk9WRn5K1gZScA8FG3Z/2jjHXjoDNvBg4EGioUIyxW/wiM
	//tyyUs361b8IrkSQxebWRgH9AYOgBK0uTCaQdGkBPjv37vJLFEIoYe5JaOeqCaaY8
	//kBSyv4YMJxARd6W1EQ5MbCncFwKBgQDY37S+KAaolT0jBouBkTr6xGC5VNY93RBg
	//6RBem8VpiRoy1AUIhfBq9hoo+S1JjkkyV29e6oANA0+VaFkSXHdyKNFJvA9SPyMd
	//UyEDpqhkzt1Iwz7mVPDIyVxZdQl3MhBqBamRGBF7WdZ0Uu1Yv08qr8y92pL/pAaq
	//ynKtCXnO+QKBgQD7zK1sqHPPZtlfaBcSPK21TyFIqePIJyacH49SX4+5KVZjKU7d
	//Ky6GUUd3/OW5NzI7QvAXOCN+FukQs0YwvDA4iyNzbOQUa4xCRaCII5PonRSyM8SE
	//PQSGnz4ckDMca8ml0aA4aA6ruLqFu0iirMUT2YpZ90d/Rdip5d8X6/v6uQKBgGKt
	//LrTSy2zLMp0MLk+Ov/I3hfbthelx+zDM+pjBcYo6SoRT+dJN9v0D6xU3gwaTyfQw
	//2kiqlAbXuc82Qkjbb7GupsNQ4PvAJH8EQuJBYx8zDHY9+BGfFkuVawJ9AQiN6/AB
	//kYymEdY5Ix2cWcfmi+PVR2ge2oGcpTfyk0juJdEpAoGBAOsi8BUWAHo2ubh9fshb
	//wj8liqjKQmobf40QO1h9ukp5f2u9bY6TN6uS6uyeIZAH0R4abvpcdYb3SbVmpynY
	//sff1yIh0NUQTd0QcTAFNX4Q1YnyyOlLrOZPT+VykKY0O4w1Z039fK70ESAApkuXD
	//b9Z/o5veuqDXWMG2CLfKZiJl
	//-----END PRIVATE KEY-----
	//`
	//client, err := wechat.NewClientV3("1645242806", "39E5FC202CDA58542566BFA81FBB29EA1A38B7D8", "BbjLJfOIw03YMtTsuuBSgB5OsyNP9lIi", PrivateKey)
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	// 39E5FC202CDA58542566BFA81FBB29EA1A38B7D8
	// 4D1FB333E14494FB751D1938E7DCEE5783C93DAF
	var settings *entity.PaySettings
	err := dao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		return
	}
	log.Println("WeixinMchid------------>", settings.WeixinMchid)
	log.Println("WeixinSerialno------------>", settings.WeixinSerialno)
	log.Println("WeixinApikey------------>", settings.WeixinApikey)
	log.Println("WeixinPrivatekey------------>", settings.WeixinPrivatekey)
	client, err := wechat.NewClientV3(settings.WeixinMchid, settings.WeixinSerialno, settings.WeixinApikey, settings.WeixinPrivatekey)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 设置微信平台API证书和序列号（推荐开启自动验签，无需手动设置证书公钥等信息）
	//client.SetPlatformCert([]byte(""), "")

	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return
	}

	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn

	wxclient = client
}
