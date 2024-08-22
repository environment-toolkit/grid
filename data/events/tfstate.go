package events

import (
	"time"

	"github.com/environment-toolkit/grid/data/models"

	"github.com/google/uuid"
)

type TFStateUpdated struct {
	State        models.TFStateState `json:"state"`
	DeploymentId uuid.UUID           `json:"deployment_id"`
	Key          string              `json:"key"`

	StateFile interface{} `json:"state_file"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type TFStateLocked struct {
	State        models.TFStateState `json:"state"`
	DeploymentId uuid.UUID           `json:"deployment_id"`
	Key          string              `json:"key"`

	LockedAt time.Time `json:"locked_at"`
}

type TFStateUnlocked struct {
	State        models.TFStateState `json:"state"`
	DeploymentId uuid.UUID           `json:"deployment_id"`
	Key          string              `json:"key"`

	LockedAt *time.Time `json:"locked_at"`
}

type TFStateDeleted struct {
	State     models.TFStateState `json:"state"`
	DeletedAt time.Time           `json:"deleted_at"`
}
