package libDownloadAtt

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetAtt(url string) (string, []byte) {
	//发起网络请求
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	// name := ""
	// 获取filename
	// getDispos := res.Header.Get("Content-Disposition")
	//if getDispos != "" {
	//	_, params, err := mime.ParseMediaType(getDispos)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	//var ok bool
	//	//name, ok = params["filename"]
	//	//if !ok {
	//	//	//log.Println(ok)
	//	//	path := strings.Split(url, "/")
	//	//	name = path[len(path)-1]
	//	//}
	//	//name, _ = gurl.Decode(name)
	//
	//}
	// defer 释放资源
	path := strings.Split(url, "/")
	name := path[len(path)-1]
	defer res.Body.Close()
	var buff []byte
	buff, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return name, buff
}
