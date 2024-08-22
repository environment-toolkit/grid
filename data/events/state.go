package events

import (
	"time"

	"github.com/environment-toolkit/grid/data/models"
	"github.com/google/uuid"
)

type StateUpdating struct {
	State      models.StateState `json:"state"`
	SpecId     uuid.UUID         `json:"spec_id"`
	Target     models.Target     `json:"target"`
	UpdatingAt time.Time         `json:"updating_at"`
}

type StateDeleted struct {
	State     models.StateState `json:"state"`
	DeletedAt time.Time         `json:"deleted_at"`
}
