package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"tender/internal/app/common/bidsearch"
	"tender/internal/app/common/search"
)

var TS search.KnowledgeSearchService
var BD bidsearch.BidSearchService

func Initialize(ctx g.Ctx) {
	TS = search.SearchService(ctx)
	BD = bidsearch.SearchService(ctx)
}
