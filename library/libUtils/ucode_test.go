package libUtils

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// EscapeUnicode 字符转码成unicode编码
func EscapeUnicode(text string) string {
	unicodeText := strconv.QuoteToASCII(text)
	// 去掉返回内容两端多余的双引号
	return unicodeText[1 : len(unicodeText)-1]
}

// UnescapeUnicode 将unicode编码转换成中文
func UnescapeUnicode(uContent string) (string, error) {
	// 转码前需要先增加上双引号，Quote增加双引号会将\u转义成\\u，同时会处理一些非打印字符
	content := strings.Replace(strconv.Quote(uContent), `\\u`, `\u`, -1)
	text, err := strconv.Unquote(content)
	if err != nil {
		return "", err
	}
	return text, nil
}

func Test_N(t *testing.T) {
	text := "1685370532604【爬虫工程师_深圳】王锦 6年.pdf"
	textL := strings.Split(text, "")
	fmt.Println(textL)
	for _, v := range textL {
		fmt.Println(v)

	}
	//text = strings.Replace(text, "】", "", -1)
	//fmt.Println(text)
}
