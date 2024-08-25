package handler

import (
	"context"
	"net/http"

	"github.com/environment-toolkit/grid/controllers"
	"github.com/environment-toolkit/grid/data"
	"github.com/environment-toolkit/grid/data/aggregates"
	"github.com/environment-toolkit/grid/data/commands"

	"go.uber.org/zap"

	"github.com/go-apis/eventsourcing/es"
	_ "github.com/go-apis/eventsourcing/es/providers/data/pg"
	_ "github.com/go-apis/eventsourcing/es/providers/data/sqlite"
	_ "github.com/go-apis/eventsourcing/es/providers/stream/noop"
	"github.com/go-chi/chi/v5"

	"github.com/go-apis/utils/xes"
	"github.com/go-apis/utils/xlog"
	"github.com/go-apis/utils/xopenapi"
	"github.com/go-apis/utils/xservice"

	"github.com/swaggest/rest/nethttp"
)

func TrySet(cfg *es.ProviderConfig, pubsub es.MemoryBusPubSub) {
	if cfg.Stream.Memory == nil || pubsub == nil {
		return
	}

	cfg.Stream.Memory.PubSub = pubsub
}

type Config struct {
	Security struct {
		SignKey string
	}
}

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, svc *xservice.Service, pubsub es.MemoryBusPubSub) (http.Handler, error) {
	log := xlog.Logger(ctx)

	appcfg := &Config{}
	if err := svc.Parse(appcfg); err != nil {
		log.Error("failed to parse config", zap.Error(err))
		return nil, err
	}

	security, err := xes.NewSecurity(appcfg.Security.SignKey)
	if err != nil {
		log.Error("failed to create security", zap.Error(err))
		return nil, err
	}

	pcfg := &es.ProviderConfig{}
	if err := svc.Parse(pcfg); err != nil {
		return nil, err
	}

	TrySet(pcfg, pubsub)

	cli, err := data.NewClient(ctx, pcfg)
	if err != nil {
		return nil, err
	}

	chi.RegisterMethod("LOCK")
	chi.RegisterMethod("UNLOCK")

	s := xopenapi.New(svc.ServiceConfig)

	// Setup middlewares.
	s.Wrap(
		es.CreateUnit(cli),
	)

	s.Post("/replay/deployment", xes.NewReplayAllInteractor("Deployment"))
	s.Post("/replay/deploymentrevision", xes.NewReplayAllInteractor("DeploymentRevision"))
	s.Post("/replay/config", xes.NewReplayAllInteractor("Config"))

	s.Method(http.MethodGet, "/state/{namespace}/{id}/{key}", controllers.GetState())
	s.Method(http.MethodPost, "/state/{namespace}/{id}/{key}", controllers.PostState())
	s.Method(http.MethodDelete, "/state/{namespace}/{id}/{key}", controllers.DeleteState())
	s.Method("LOCK", "/state/{namespace}/{id}/{key}", controllers.LockState())
	s.Method("UNLOCK", "/state/{namespace}/{id}/{key}", controllers.UnlockState())

	// Security
	s.Wrap(
		nethttp.HTTPBearerSecurityMiddleware(s.OpenAPICollector, "auth", "Authentication", "JWT"),
		es.CreateUnit(cli),
		security.Middleware(true),
	)

	s.Get("/specs", xes.NewPagingEntityInteractor[*aggregates.Spec, *PagedSpecsInput]())
	s.Get("/specs/{id}", xes.NewGetEntityInteractor[*aggregates.Spec]())
	s.Post("/specs", xes.NewCommandInteractor[*commands.NewSpec]())
	s.Post("/specs/delete", xes.NewCommandInteractor[*commands.DeleteSpec]())

	s.Get("/environments", xes.NewPagingEntityInteractor[*aggregates.Environment, *PagedEnvironmentsInput]())
	s.Get("/environments/{id}", xes.NewGetEntityInteractor[*aggregates.Environment]())
	s.Post("/environments", xes.NewCommandInteractor[*commands.NewEnvironment]())
	s.Post("/environments/delete", xes.NewCommandInteractor[*commands.DeleteEnvironment]())

	return s, nil
}
