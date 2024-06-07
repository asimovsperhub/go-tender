package bidsearch

import (
	"github.com/gogf/gf/v2/frame/g"
	"sync"
	"tender/utility/zinc"
)

var onceTs sync.Once
var bd BidSearchService

func SearchService(ctx g.Ctx) BidSearchService {
	onceTs.Do(func() {
		bd = NewBidZincSearchService(ctx)
	})

	return bd
}

type BidSearchService interface {
	Query(indexName string, q interface{}) (interface{}, error)
}

type zincSearchServant struct {
	indexName     string
	client        *zinc.ZincClient
	publicFilter  string
	privateFilter string
	friendFilter  string
}

func NewBidZincSearchService(ctx g.Ctx) BidSearchService {
	s := g.Cfg()
	zts := &zincSearchServant{
		indexName: s.MustGet(ctx, "bdzinc.index").String(),
		client:    zinc.NewClient(ctx, s),
	}
	zts.client.ZincPassword = s.MustGet(ctx, "bdzinc.password").String()
	zts.client.ZincUser = s.MustGet(ctx, "bdzinc.user").String()
	zts.client.ZincHost = s.MustGet(ctx, "bdzinc.host").String()
	return zts
}

func (s *zincSearchServant) Query(indexName string, q interface{}) (interface{}, error) {
	result, err := s.client.EsQuery(indexName, q)
	if err != nil {
		return nil, err
	}
	//res := new(*deskentity.Bid)
	// var source map[string]interface{}
	hits := result.Hits.Hits
	return hits, nil
	// print(hits)
	//for i := 0; i < len(hits); i++ {
	//	log.Println(hits[i].Source)
	//}
}
