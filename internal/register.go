package internal

import (
	"github.com/ekkinox/yokai-es/internal/elastic"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Provide(elastic.NewIndexer),
	)
}
