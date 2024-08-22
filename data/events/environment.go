package events

import (
	"time"

	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/eventsourcing/es"
)

type EnvironmentCreated struct {
	es.BaseEvent `es:"publish"`

	State     models.EnvironmentState `json:"state"`
	Name      string                  `json:"name"`
	Title     string                  `json:"title"`
	CreatedAt time.Time               `json:"created_at"`
}

type EnvironmentDeleted struct {
	es.BaseEvent `es:"publish"`

	State     models.EnvironmentState `json:"state"`
	DeletedAt time.Time               `json:"deleted_at"`
}
