package consensus

import (
	"context"
	"encoding/json"
	"github.com/NavExplorer/navexplorer-indexer-go/internal/elastic_cache"
	"github.com/NavExplorer/navexplorer-indexer-go/pkg/explorer"
	"github.com/getsentry/raven-go"
	"github.com/olivere/elastic/v7"
)

type Repository struct {
	Client *elastic.Client
}

func NewRepo(client *elastic.Client) *Repository {
	return &Repository{client}
}

func (r *Repository) GetConsensus() (*explorer.Consensus, error) {
	results, err := r.Client.Search(elastic_cache.ConsensusIndex.Get()).
		Size(1).
		Do(context.Background())
	if err != nil || results == nil {
		raven.CaptureError(err, nil)
		return nil, err
	}

	if len(results.Hits.Hits) == 0 {
		return nil, elastic_cache.ErrRecordNotFound
	}

	var consensus *explorer.Consensus
	hit := results.Hits.Hits[0]
	if err = json.Unmarshal(hit.Source, &consensus); err != nil {
		raven.CaptureError(err, nil)
		return nil, err
	}

	return consensus, nil
}
