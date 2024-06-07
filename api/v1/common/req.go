package common

import "github.com/gogf/gf/v2/frame/g"

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}

// PageReq 公共请求参数
type PageReq struct {
	DateRange []string `p:"dateRange"` //日期范围
	Start     string   `p:"start"`
	End       string   `p:"end"`
	PageNum   int      `p:"pageNum"`  //当前页码
	PageSize  int      `p:"pageSize"` //每页数
	OrderBy   string   //排序方式
}

type UploadReq struct {
	g.Meta `path:"/publish/knowledge" tags:"上传文件" method:"post" summary:"上传"`
	File   string `p:"file"`
}
