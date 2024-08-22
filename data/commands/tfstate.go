package commands

import (
	"github.com/go-apis/eventsourcing/es"
	"github.com/google/uuid"
)

type UpdateTFState struct {
	es.BaseCommand

	DeploymentId uuid.UUID `json:"deployment_id" required:"true"`
	Key          string    `json:"key" required:"true"`
	StateFile    string    `json:"state_file" required:"true"`
}

type DeleteTFState struct {
	es.BaseCommand

	DeploymentId uuid.UUID `json:"deployment_id" required:"true"`
	Key          string    `json:"key" required:"true"`
}

type LockTFState struct {
	es.BaseCommand

	DeploymentId uuid.UUID `json:"deployment_id" required:"true"`
	Key          string    `json:"key" required:"true"`
}

type UnlockTFState struct {
	es.BaseCommand

	DeploymentId uuid.UUID `json:"deployment_id" required:"true"`
	Key          string    `json:"key" required:"true"`
}
