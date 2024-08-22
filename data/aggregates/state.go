package aggregates

import (
	"context"
	"errors"
	"time"

	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/events"
	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/eventsourcing/es"
)

var ErrStateUpdating = errors.New("state is updating")
var ErrStateDeleting = errors.New("state is deleting")

type State struct {
	es.BaseAggregateSourced

	State models.StateState `json:"state" required:"true"`
	Name  string            `json:"name" required:"true"`
	Title string            `json:"title" required:"true"`

	CreatedAt time.Time  `json:"created_at" required:"true"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	SyncedAt  *time.Time `json:"synced_at"`
}

func (a *State) HandleUpdateState(ctx context.Context, cmd *commands.UpdateState) error {
	if a.State == models.StateStateUpdating {
		return ErrStateUpdating
	}
	if a.State == models.StateStateDeleting {
		return ErrStateDeleting
	}

	return a.Apply(ctx, &events.StateUpdating{
		State:      models.StateStateUpdating,
		SpecId:     cmd.SpecId,
		Target:     cmd.Target,
		UpdatingAt: time.Now(),
	})
}

func (a *State) HandleDeleteState(ctx context.Context, cmd *commands.DeleteState) error {
	if a.State == models.StateStateUpdating {
		return ErrStateUpdating
	}
	if a.State == models.StateStateDeleting {
		return ErrStateDeleting
	}

	return a.Apply(ctx, &events.StateDeleted{
		State:     models.StateStateDeleting,
		DeletedAt: time.Now(),
	})
}

func NewState() *State {
	return &State{
		State: models.StateStateNew,
	}
}
