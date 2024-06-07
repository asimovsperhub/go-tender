package qbucket

import (
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)

//var onceCos sync.Once
//var qbucket QbucketService

//func QbucketToService(bucket string) QbucketService {
//	onceCos.Do(func() {
//		qbucket = NewQbucketService(bucket)
//	})
//	return qbucket
//}

type QbucketService interface {
	Upload(fd io.Reader, filetype string, filename string, filesize int64) error
	GetObject(ObjectName string) []byte
}

func NewQbucketService(bucket string) QbucketService {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.ap-guangzhou.myqcloud.com", bucket))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDPeeOs5B0ZBYfYCakdGfohKKWTBdrwlKg", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: "CBwVptgFYaNyOglkX6WihSyJNFuMcgV4",     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	ccos := &CosServant{
		client: c,
	}
	return ccos
}
