package search

import (
	"github.com/gogf/gf/v2/frame/g"
	"sync"
	"tender/utility/zinc"
)

var onceTs sync.Once
var ts KnowledgeSearchService

func SearchService(ctx g.Ctx) KnowledgeSearchService {
	onceTs.Do(func() {
		ts = NewZincSearchService(ctx)
	})
	return ts
}

type KnowledgeSearchService interface {
	IndexName() string
	AddDocuments(documents DocItems, primaryKey ...string) (bool, error)
	DeleteDocuments(identifiers []string) error
}

func NewZincSearchService(ctx g.Ctx) KnowledgeSearchService {
	s := g.Cfg()
	zts := &zincSearchServant{
		indexName: s.MustGet(ctx, "zinc.index").String(),
		client:    zinc.NewClient(ctx, s),
	}
	zts.createIndex()

	return zts
}
