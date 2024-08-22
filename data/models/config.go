package models

import "github.com/google/uuid"

type Target struct {
	EnvironmentId uuid.UUID `json:"environment_id"`
	Region        string    `json:"region"`
}

type Resource struct {
	Type  string
	Name  string
	Props map[string]interface{}
}

type Access struct {
	Inbound  []string
	Outbound []string
}

type Manifest struct {
	Resource

	Resources []*Resource
	Access    *Access
}
