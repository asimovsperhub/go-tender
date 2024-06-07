package desk

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 发布知识
type VipOrderWxNotifyReq struct {
	g.Meta `path:"/pay/wxnotify" tags:"会员订阅" method:"post" summary:"订单回调通知"`
}
type VipOrderWxNotifyRes struct {
}
