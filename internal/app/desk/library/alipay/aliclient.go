package ali

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"log"
	"sync"
	"tender/internal/app/system/dao"
	"tender/internal/app/system/model/entity"
)

var alipayClient *alipay.Client
var once sync.Once

func GetClient(ctx context.Context) *alipay.Client {
	if alipayClient == nil {
		once.Do(func() {
			InitAlipayClient(ctx)
		})
		return alipayClient
	}
	return alipayClient
}

func InitAlipayClient(ctx context.Context) {
	// NewClientV3 初始化微信客户端 v3
	// mchid：商户ID 或者服务商模式的 sp_mchid
	// serialNo：商户证书的证书序列号
	// apiV3Key：apiV3Key，商户平台获取
	// privateKey：私钥 apiclient_key.pem 读取后的内容
	//	PrivateKey := `MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCFjLa/qYrW27XDk68Z+wCsy+ZANP/JlHY2pKfWefGDj3jTQJVOEU9wnfEdaiSXHt4CGv/Z2Eb4YAmpwgHiaIQ5jq42SziYQCJQbVtoyM92kdWgEiauzchW42/2DonB86EbLCLTNqvZ1x7OJdoAspunJAQgzH/iV519En3O+Tmo+S4fxoey3ivivofGkMZXRLvEcKkXvpHnojngoDRwbVfhtcSbnNOfoPwoCA3DOPmLDWNLIkfB9ifRY0BN0tCUYmcbNrONwD7zsns470IOhf1lFWWcsf0SUqZq/KhcR84RJ/P0jXmwvLOX6hLikH/pv1xaoeg2VMaEUPbYETBCnEglAgMBAAECggEASrHUOMir5dZHCFdelUplK0BzzK+X7PgDUjavxO65XRruukEiAOL/qn48MHaAdqGGDGlrdj1YHG0imnbvNOG+Kq8Z0zYpNkaQqqkveiDuoGOatwfP6O8hwS8+HMIu3HLWReteowGuBo4iJazYDJHZKxei03FtiXv1ktG32f5c+Jtuze77GoP+U8V99AF8JZysTvlQD/0d1LzBHzyT7bqTejMrk0nDn5nufkZRWQbcjkEO3g3ShVqgbI4l3ZdCZHNQLHBLsW+t4Y/AgizZXht/iu0Ojhc/NkNEsHldXQf2C6XKKx4QcksyrBj9anBOz2lCNWXv+Nad+S77TGEqlm50oQKBgQDCg6DuNUhE8zJrOQb8jcA4NZwCJoTpdVF8ZNrDzHp9MS4I40IE7TNF+BsbFi5ZoaP3syjksm8CInN4G6ouHGAKp9cFD9fLBicxS/m1ZFhM+KcfYT1XuGQdptWhINl3TY7AJoKm9rInj1+g/hv178RgJP40Cm9upAicQ9yPtxEWaQKBgQCvw791LVTOETRk11l4uQXINN80aneH5uVY+el/WL96bFHN2PX48qgq3HQiOIvW3IPVRLSjl4qD2tkqLUaovhM+gwdYtzkzOeAPgS+aoEwHv8DO8Z1httup6EQ8lJnZVVtE6M97mSt4UoSjN7pV8oVen8EDYtI8eqa0cKdP8fGEXQKBgDocTui7XFLVAPwNdmvfAU0Jnwj5bv28AdMS4lRac0GRfrpDAocsQgQmQhrOfpxicPTgPlhMgmN13V9FjIroCT1FtTQa7pIFzZGpq/kn+EVOh8cVKtlZCffdzrardKxyrDH1j1TIIKM27w/OupW3wsgiZRsw/udj8/qTP4Jj46P5AoGBAIVJ2ZAfz1bx3xy13ojZLtRTnAygzIKIC9a8tmC3SYWqTSLgbC/cvMC8K2mkHg/TbDo3/xCsJAO328XLTfE7K1bVgKW7VKpPMmYvno0REHcz6CBHRAVM6SnhFJYoTr9spmkMcAOX9UoqsaEg6rKw1okadwF9WFc6396oK4lJvdCVAoGANiYApJRzjeRQGKRowAQLpFDmGGjcJ9//Pqg5uw7vybWVS+XtlMbT84JeqxGQO1VtvRWsLtiv7RE2Or86da16ClCYE8dh+r185kS7IYemshbzIpf7QvEa71Cu45YGWgMO7VY7xR2Q1Y/3AxZLOO517gDWGr4Z5ogg6moQkwBFKfU=
	//`
	//	client, err := alipay.NewClient("2021003184660077", PrivateKey, true)
	//	if err != nil {
	//		xlog.Error(err)
	//		return
	//	}
	var settings *entity.PaySettings
	err := dao.PaySettings.Ctx(ctx).Scan(&settings)
	if err != nil {
		xlog.Error(err)
		fmt.Println("alipay PaySettings---->", err.Error())
		return
	}
	client, err := alipay.NewClient(settings.AlipayAppid, settings.AlipayPrivatekey, true)
	if err != nil {
		xlog.Error(err)
		fmt.Println("alipay NewClient---->", err.Error())
		return
	}
	log.Println("AlipayAppid------------>", settings.AlipayAppid)
	log.Println("AlipayPrivatekey------------>", settings.AlipayPrivatekey)
	log.Println("AlipayAppCertPublicKey------------>", settings.AlipayAppCertPublicKey)
	log.Println("AlipayRootCert------------>", settings.AlipayRootCert)
	log.Println("AlipayPublicCert------------>", settings.AlipayPublicCert)
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8). // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2) // 设置签名类型，不设置默认 RSA2

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign([]byte("alipayCertPublicKey_RSA2 bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err = client.SetCertSnByPath("./manifest/key/appCertPublicKey_2021003184660077.crt", "./manifest/key/alipayRootCert.crt", "./manifest/key/alipayCertPublicKey_RSA2.crt")
	// 证书内容
	err = client.SetCertSnByContent([]byte(settings.AlipayAppCertPublicKey), []byte(settings.AlipayRootCert), []byte(settings.AlipayPublicCert))
	if err != nil {
		fmt.Println("alipay SetCertSnByContent---->", err.Error())
		return
	}

	alipayClient = client
}
