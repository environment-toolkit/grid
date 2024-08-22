package events

import (
	"time"

	"github.com/environment-toolkit/grid/data/models"
)

type SpecCreated struct {
	State     models.SpecState  `json:"state"`
	Name      string            `json:"name"`
	Content   []byte            `json:"content"`
	Variables map[string]string `json:"variables"`
	CreatedAt time.Time         `json:"created_at"`
}

type SpecDeleted struct {
	State     models.SpecState `json:"state"`
	DeletedAt time.Time        `json:"deleted_at"`
}
