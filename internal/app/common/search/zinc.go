package search

import (
	"tender/utility/types"
	"tender/utility/zinc"
)

type zincSearchServant struct {
	indexName     string
	client        *zinc.ZincClient
	publicFilter  string
	privateFilter string
	friendFilter  string
}

func (s *zincSearchServant) Name() string {
	return "Zinc"
}

type DocItems []map[string]interface{}

func (s *zincSearchServant) IndexName() string {
	return s.indexName
}

func (s *zincSearchServant) AddDocuments(data DocItems, primaryKey ...string) (bool, error) {
	buf := make(DocItems, 0, len(data)+1)
	if len(primaryKey) > 0 {
		buf = append(buf, map[string]types.Any{
			"index": map[string]types.Any{
				"_index": s.indexName,
				"_id":    primaryKey[0],
			},
		})
	} else {
		buf = append(buf, map[string]types.Any{
			"index": map[string]types.Any{
				"_index": s.indexName,
			},
		})
	}
	buf = append(buf, data...)
	return s.client.BulkPushDoc(buf)
}

func (s *zincSearchServant) DeleteDocuments(identifiers []string) error {
	for _, id := range identifiers {
		if err := s.client.DelDoc(s.indexName, id); err != nil {
			return err
		}
	}
	return nil
}

func (s *zincSearchServant) createIndex() {
	// 不存在则创建索引
	s.client.CreateIndex(s.indexName, &zinc.ZincIndexProperty{
		"id": &zinc.ZincIndexPropertyT{
			Type:     "numeric",
			Index:    true,
			Store:    true,
			Sortable: true,
		},
		"title": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"user_id": &zinc.ZincIndexPropertyT{
			Type:  "numeric",
			Index: true,
			Store: true,
		},
		"primary_classification": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"cover_url": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"details_url": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"video_url": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"video_introduction": &zinc.ZincIndexPropertyT{
			Type:  "text",
			Index: true,
			Store: true,
		},
		"content": &zinc.ZincIndexPropertyT{
			Type:           "text",
			Index:          true,
			Store:          true,
			Aggregatable:   true,
			Highlightable:  true,
			Analyzer:       "gse_search",
			SearchAnalyzer: "gse_standard",
		},
	})
}
