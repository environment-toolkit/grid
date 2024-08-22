package resolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/environment-toolkit/grid/data/models"
)

type Session interface {
	NewValue(value string) (*models.Value, error)
	Resolve() error
}

type session struct {
	ctx       context.Context
	variables map[string]string

	vars    []*models.Value
	secrets []*models.Value
	grids   []*models.Value
}

func (m *session) NewValue(token string) (*models.Value, error) {
	// we only support "var" and "secret" values
	// var:abc -> var, abc
	// secret:abc -> secret, abc

	before, after, ok := strings.Cut(token, ":")
	if !ok {
		return nil, fmt.Errorf("invalid token: %s", token)
	}

	key := models.ValueType(before)
	v := &models.Value{Key: after, ValueType: key}

	switch key {
	case models.Var:
		m.vars = append(m.vars, v)
		return v, nil
	case models.Secret:
		m.secrets = append(m.secrets, v)
		return v, nil
	case models.Grid:
		m.grids = append(m.grids, v)
		return v, nil
	default:
		return nil, fmt.Errorf("invalid value type: %s", key)
	}
}

func (m *session) Resolve() error {
	for _, v := range m.vars {
		if value, ok := m.variables[v.Key]; ok {
			v.Value = &value
		}
	}

	for _, v := range m.secrets {
		// resolve the secret value.
		value := "ssssh"
		v.Value = &value
	}

	for _, v := range m.grids {
		// resolve the secret value.
		value := "not done yet"
		v.Value = &value
	}

	return nil
}

func NewSession(ctx context.Context, variables map[string]string) Session {
	return &session{
		ctx:       ctx,
		variables: variables,
	}
}
