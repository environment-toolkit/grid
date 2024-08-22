package commands

import (
	"github.com/environment-toolkit/grid/data/models"
	"github.com/go-apis/eventsourcing/es"
	"github.com/google/uuid"
)

type UpdateState struct {
	es.BaseNamespaceCommand
	es.BaseCommand

	SpecId uuid.UUID     `json:"spec_id" required:"true"`
	Target models.Target `json:"target" required:"true"`
}

type DeleteState struct {
	es.BaseNamespaceCommand
	es.BaseCommand
}
