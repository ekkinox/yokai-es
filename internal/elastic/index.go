package elastic

import (
	"context"

	"github.com/ankorstore/yokai/log"
	"github.com/ankorstore/yokai/trace"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Indexer struct {
	client *elasticsearch.Client
}

func NewIndexer(client *elasticsearch.Client) *Indexer {
	return &Indexer{
		client: client,
	}
}

func (i *Indexer) ListIndex(ctx context.Context) (*esapi.Response, error) {
	ctx, span := trace.CtxTracer(ctx).Start(ctx, "Indexer.ListIndex")
	defer span.End()

	log.CtxLogger(ctx).Info().Msg("Indexer.ListIndex")

	return i.client.Cat.Indices(
		i.client.Cat.Indices.WithContext(ctx),
		i.client.Cat.Indices.WithFormat("json"),
	)
}
