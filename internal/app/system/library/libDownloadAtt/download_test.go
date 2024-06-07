package libDownloadAtt

import (
	"bufio"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gurl"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
	"testing"
)

func Test_Down(*testing.T) {
	name, buff := Get("https://www.gdzwfw.gov.cn/ggzy-portal/base/sys-file/trading-download?attachGuid=1d74c21e61c644e285a48373c74c316a1677469505859")
	log.Println(name)
	log.Println(buff)
}

func Get(url string) (string, []byte) {
	//发起网络请求
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	name := "file"
	// 获取filename
	getDispos := res.Header.Get("Content-Disposition")
	if getDispos != "" {
		_, params, err := mime.ParseMediaType(getDispos)
		if err != nil {
			log.Println(err)
		}
		var ok bool
		name, ok = params["filename"]
		if !ok {
			log.Println(ok)
		} else {
			name, _ = gurl.Decode(name)
		}
	}
	// defer 释放资源
	defer res.Body.Close()
	//path := strings.Split(url, "/")
	//name := path[len(path)-1]
	var buff []byte
	buff, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return name, buff
}

func Test_file(t *testing.T) {
	url := "https://www.gdzwfw.gov.cn/ggzy-portal/base/sys-file/trading-download?attachGuid=1d74c21e61c644e285a48373c74c316a1677469505859"
	//发起网络请求
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	// defer 释放资源
	defer res.Body.Close()
	//定义文件名字
	path := strings.Split(url, "/")
	name := path[len(path)-1]
	//创建文件
	out, err := os.Create("./" + name + ".docx")
	if err != nil {
		log.Println(err)
		return
	}
	// defer延迟调用 关闭文件，释放资源
	defer out.Close()
	//添加缓冲 bufio 是通过缓冲来提高效率。
	wt := bufio.NewWriter(out)
	_, _ = io.Copy(wt, res.Body)
	//将缓存的数据写入到文件中
	_ = wt.Flush()
}
