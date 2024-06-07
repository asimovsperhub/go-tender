package service

import (
	"context"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/gogf/gf/v2/frame/g"
)

func Mp() *core.Client {
	return wechatClient
}

var (
	accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(g.Cfg().MustGet(context.TODO(), "mp.wxappid").String(), g.Cfg().MustGet(context.TODO(), "mp.wxappsecret").String(), nil)
	wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
)
