package commands

import (
	"github.com/go-apis/eventsourcing/es"
)

type NewSpec struct {
	es.BaseNamespaceCommand
	es.BaseCommand

	Content   []byte            `json:"data" required:"true"`
	Variables map[string]string `json:"variables"`
}

type DeleteSpec struct {
	es.BaseNamespaceCommand
	es.BaseCommand
}
