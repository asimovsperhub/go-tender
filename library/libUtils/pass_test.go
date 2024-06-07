package libUtils

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gurl"
	"log"
	"testing"
)

func Test_Pass(t *testing.T) {
	Md5()
}

func Md5() {
	pass := gmd5.MustEncryptString("123456")
	log.Println(pass)
	log.Println(EncryptPassword(pass, "YDIOyBk2Qm"))
}

func Test_Gurl(t *testing.T) {
	// 解析URL并返回其组件
	data, err := gurl.ParseURL("https://biaoziku.com/api/v1/desk/index/announcement?keyWords=35%&city=&type=%E6%84%8F%E5%90%91%E5%85%AC%E5%91%8A&industry=&dateRange=&pageNum=1&pageSize=16", -1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}
