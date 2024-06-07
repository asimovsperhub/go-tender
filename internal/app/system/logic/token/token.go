package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"log"
	"tender/gftoken"
	"tender/internal/app/common/consts"
	commonModel "tender/internal/app/common/model"
	"tender/internal/app/system/service"
)

type sToken struct {
	*gftoken.GfToken
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *commonModel.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	log.Println("cc", opt)
	liberr.ErrIsNil(ctx, err)
	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	// 当您尝试将具体类型分配或传递（或转换）为接口类型时，会出现此编译时错误；并且类型本身不实现接口，只是指向类型的指针
	service.RegisterGToken(New())
}
