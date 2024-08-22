package tests

import (
	"context"

	"github.com/environment-toolkit/grid/data"
	"github.com/environment-toolkit/grid/handler"

	"github.com/go-apis/eventsourcing/es"
	_ "github.com/go-apis/eventsourcing/es/providers/data/sqlite"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/noop"

	"github.com/go-apis/utils/xservice"
)

type ServiceTester interface {
	Client() es.Client
}

type serviceTester struct {
	cli es.Client
}

func (h *serviceTester) Client() es.Client {
	return h.cli
}

func NewServiceTester() (ServiceTester, error) {
	ctx := context.Background()

	cfg, err := xservice.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	appcfg := &handler.Config{}
	if err := cfg.Parse(appcfg); err != nil {
		return nil, err
	}
	var pcfg es.ProviderConfig
	if err := cfg.Parse(&pcfg); err != nil {
		return nil, err
	}

	cli, err := data.NewClient(ctx, &pcfg)
	if err != nil {
		return nil, err
	}

	return &serviceTester{
		cli: cli,
	}, nil
}
