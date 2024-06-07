package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"tender/internal/app/common/controller"
)

type BaseController struct {
	controller.BaseController
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
