package elastic

import (
	"net/http"

	"github.com/ankorstore/yokai/config"
	"go.uber.org/fx"

	"github.com/elastic/go-elasticsearch/v8"
)

const ModuleName = "elastic"

var ElasticModule = fx.Module(
	ModuleName,
	fx.Provide(
		ProvideElasticSearchClient,
	),
)

type ProvideElasticSearchClientParams struct {
	fx.In
	Config       *config.Config
	RoundTripper http.RoundTripper
}

func ProvideElasticSearchClient(p ProvideElasticSearchClientParams) (*elasticsearch.Client, error) {
	var cfg = elasticsearch.Config{
		CloudID:   p.Config.GetString("modules.elastic.cloud_id"),
		APIKey:    p.Config.GetString("modules.elastic.api_key"),
		Transport: p.RoundTripper,
	}

	return elasticsearch.NewClient(cfg)
}
