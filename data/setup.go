package data

import (
	"context"

	"github.com/environment-toolkit/grid/data/aggregates"
	"github.com/environment-toolkit/grid/data/sagas"

	"github.com/go-apis/eventsourcing/es"
)

func NewClient(
	ctx context.Context,
	pcfg *es.ProviderConfig,
) (es.Client, error) {
	reg, err := es.NewRegistry(
		pcfg.Service,

		aggregates.NewEnvironment(),
		aggregates.NewSpec(),
		aggregates.NewState(),
		aggregates.NewTFState(),

		sagas.NewDeploySaga(),
	)
	if err != nil {
		return nil, err
	}

	return es.NewClient(ctx, pcfg, reg)
}
