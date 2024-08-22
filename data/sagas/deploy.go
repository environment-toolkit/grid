package sagas

import (
	"context"

	"github.com/environment-toolkit/grid/data/events"
	"github.com/go-apis/eventsourcing/es"
)

type deploySaga struct {
	es.BaseSaga `es:"group=internal"`
}

func (s *deploySaga) HandleStateUpdating(ctx context.Context, evt *es.Event, data *events.StateUpdating) ([]es.Command, error) {
	// do nothing on everything else.
	return nil, nil
}

func NewDeploySaga() es.IsSaga {
	return &deploySaga{}
}
