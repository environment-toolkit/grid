package aggregates

import (
	"context"
	"time"

	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/events"
	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/eventsourcing/es"
)

type Environment struct {
	es.BaseAggregateSourced

	State models.EnvironmentState `json:"state" required:"true"`
	Name  string                  `json:"name" required:"true"`
	Title string                  `json:"title" required:"true"`

	CreatedAt time.Time  `json:"created_at" required:"true"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	SyncedAt  *time.Time `json:"synced_at"`
}

func (a *Environment) HandleNewEnvironment(ctx context.Context, cmd *commands.NewEnvironment) error {
	return a.Apply(ctx, &events.EnvironmentCreated{
		State:     models.EnvironmentStateCreating,
		Name:      cmd.Name,
		Title:     cmd.Title,
		CreatedAt: time.Now(),
	})
}

func (a *Environment) HandleDeleteEnvironment(ctx context.Context, cmd *commands.DeleteEnvironment) error {
	return a.Apply(ctx, &events.EnvironmentDeleted{
		State:     models.EnvironmentStateDeleting,
		DeletedAt: time.Now(),
	})
}

func NewEnvironment() *Environment {
	return &Environment{
		State: models.EnvironmentStateNew,
	}
}
