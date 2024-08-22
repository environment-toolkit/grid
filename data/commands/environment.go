package commands

import (
	"github.com/go-apis/eventsourcing/es"
)

type NewEnvironment struct {
	es.BaseNamespaceCommand
	es.BaseCommand

	Name  string `json:"name" required:"true"`
	Title string `json:"title" required:"true"`
}

type DeleteEnvironment struct {
	es.BaseNamespaceCommand
	es.BaseCommand
}
