package qbucket

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func Test_CosPut(t *testing.T) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://announcement-1317075511.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDPeeOs5B0ZBYfYCakdGfohKKWTBdrwlKg", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: "CBwVptgFYaNyOglkX6WihSyJNFuMcgV4",     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := "ygp/1.png"
	// 1.通过字符串上传对象
	//f := strings.NewReader("test")
	//
	//_, err := c.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	panic(err)
	//}
	// 2.通过本地文件上传对象
	//_, err = c.Object.PutFromFile(context.Background(), name, "../test", nil)
	//if err != nil {
	//	panic(err)
	//}
	// 3.通过文件流上传对象
	fd, err := os.Open("/Users/apple/Desktop/1.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	// bytes.NewReader(fd)
	res, err := c.Object.Put(context.Background(), name, fd, nil)
	log.Println(res.Body)
	if err != nil {
		panic(err)
	}
}

func Test_CosGet(t *testing.T) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://announcement-1317075511.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKIDPeeOs5B0ZBYfYCakdGfohKKWTBdrwlKg", // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: "CBwVptgFYaNyOglkX6WihSyJNFuMcgV4",     // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	name := "ygp/1.png"
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
}

func Test_GetObject(t *testing.T) {
	// https://announcement-1317075511.cos.ap-guangzhou.myqcloud.com/szygcgpt/《南澳01-03地块拆迁安置房项目通风空气调节及油烟净化工程专业分包》招标文件评审表.pdf
	// /szygcgpt/《南澳01-03地块拆迁安置房项目通风空气调节及油烟净化工程专业分包》招标文件评审表.pdf
	// https://announcement-1317075511.cos.ap-guangzhou.myqcloud.com/szygcgpt/%22%E5%AE%9A%E6%A0%87%E8%AF%84%E5%AE%A1%E6%B5%81%E7%A8%8B%E5%8D%95.pdf%22?q-sign-algorithm=sha1&q-ak=AKIDNbKxa__fmYqJPRJYBbzS18lLsidgPEslMMYn0GwDp7smii-Af-iUhe_u2NjOZW15&q-sign-time=1685029350;1685032950&q-key-time=1685029350;1685032950&q-header-list=host&q-url-param-list=&q-signature=6eb00baec1f430dfdfd345216fc19edf9368ba5e&x-cos-security-token=HCm8bLXWtoAzw8WR0gy9uHT8248qFPGa29f543920ea6c09623a1afdacd990b3caqt2N3msxhTv5B3IqofaVROi8_wkGEwY_DFXNNV_kQK2UwxodRQxlaLMSh4zTBBtZ4PQMseYM5ZymaxW3y_UP-Br_qyfyc_iHcRkkpKETj6TEMULpj7oucDjIXcWSSI_xlquQuNvF5zCvpoTzAhjtJGQflkzhTHAgsjnjG-Gh05WTVyIYa1oIe0Krj0sXsyR
	res := NewQbucketService("announcement-1317075511").GetObject("/zfcg_szggzy/[SZDL2023000641A]官龙山校区学生公寓加装隐形防护网.docx")
	log.Println(res)
}

func Test_PutObject(t *testing.T) {
	fd, err := os.Open("/Users/apple/Desktop/1.png")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	res := NewQbucketService("announcement-1317075511").Upload(fd, "/ygp", "2.png", 1)
	log.Println(res)
}
