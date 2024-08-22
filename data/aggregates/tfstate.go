package aggregates

import (
	"context"
	"fmt"
	"time"

	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/events"
	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/eventsourcing/es"
	"github.com/google/uuid"
)

var (
	ErrInvalid  = fmt.Errorf("invalid deployment id or key")
	ErrLocked   = fmt.Errorf("TF state is locked")
	ErrNotFound = fmt.Errorf("TF state not found")
)

type TFState struct {
	es.BaseAggregateSourced

	State        models.TFStateState `json:"state" required:"true"`
	DeploymentId uuid.UUID           `json:"deployment_id" required:"true"`
	Key          string              `json:"key" required:"true"`
	StateFile    string              `json:"state_file" required:"true"`

	CreatedAt time.Time  `json:"created_at" required:"true"`
	UpdatedAt *time.Time `json:"updated_at"`
	LockedAt  *time.Time `json:"locked_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (a *TFState) validate(deploymentId uuid.UUID, key string) bool {
	return a.State == models.TFStateNew || (a.DeploymentId == deploymentId && a.Key == key)
}

func (a *TFState) HandleUpdateTFState(ctx context.Context, cmd *commands.UpdateTFState) error {
	if !a.validate(cmd.DeploymentId, cmd.Key) {
		return ErrInvalid
	}

	return a.Apply(ctx, &events.TFStateUpdated{
		State:        models.TFStateUnlocked,
		DeploymentId: cmd.DeploymentId,
		Key:          cmd.Key,
		StateFile:    cmd.StateFile,
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    time.Now(),
	})
}

func (a *TFState) HandleLockTFState(ctx context.Context, cmd *commands.LockTFState) error {
	if !a.validate(cmd.DeploymentId, cmd.Key) {
		return ErrInvalid
	}
	if a.State == models.TFStateLocked {
		return ErrLocked
	}

	return a.Apply(ctx, &events.TFStateLocked{
		State:        models.TFStateLocked,
		DeploymentId: cmd.DeploymentId,
		Key:          cmd.Key,
		LockedAt:     time.Now(),
	})
}

func (a *TFState) HandleUnlockTFState(ctx context.Context, cmd *commands.UnlockTFState) error {
	if !a.validate(cmd.DeploymentId, cmd.Key) {
		return ErrInvalid
	}

	return a.Apply(ctx, &events.TFStateUnlocked{
		State:        models.TFStateUnlocked,
		DeploymentId: cmd.DeploymentId,
		Key:          cmd.Key,
	})
}

func (a *TFState) HandleDeleteTFState(ctx context.Context, cmd *commands.DeleteTFState) error {
	if !a.validate(cmd.DeploymentId, cmd.Key) {
		return ErrInvalid
	}
	if a.State == models.TFStateNew {
		return ErrNotFound
	}

	return a.Apply(ctx, &events.TFStateDeleted{
		State:     models.TFStateDeleted,
		DeletedAt: time.Now(),
	})
}

func NewTFState() *TFState {
	return &TFState{
		State:     models.TFStateNew,
		CreatedAt: time.Now(),
	}
}
