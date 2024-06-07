package common

import "github.com/gogf/gf/v2/frame/g"

type SendWsReq struct {
	g.Meta  `path:"/msg/sendWs" tags:"" method:"post" summary:""`
	UserId  uint64 `p:"userId" v:"required#用户id不能为空"`
	Content string `p:"content" v:"required#消息内容不能为空"`
}

type SendWsRes struct {
	g.Meta `mime:"application/json"`
}
