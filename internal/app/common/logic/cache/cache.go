package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-cache/cache"
	"log"
	"tender/internal/app/common/consts"
	"tender/internal/app/common/service"
)

func init() {
	service.RegisterCache(New())
}

func New() *sCache {
	var (
		ctx            = gctx.New()
		cacheContainer *cache.GfCache
	)
	// 缓存前缀 Tender
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	// 缓存数据库  redis
	model := g.Cfg().MustGet(ctx, "system.cache.model").String()
	log.Println("prefix----------------->", prefix)
	log.Println("model----------------->", model)
	if model == consts.CacheModelRedis {
		// redis
		cacheContainer = cache.NewRedis(prefix)
	} else {
		// memory
		cacheContainer = cache.New(prefix)
	}
	return &sCache{
		GfCache: cacheContainer,
		prefix:  prefix,
	}
}

type sCache struct {
	*cache.GfCache
	prefix string
}
