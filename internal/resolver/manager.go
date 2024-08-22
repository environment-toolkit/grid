package resolver

import "context"

type Manager interface {
	NewSession(ctx context.Context, variables map[string]string) Session
}

type manager struct {
}

func (m *manager) NewSession(ctx context.Context, variables map[string]string) Session {
	return NewSession(ctx, variables)
}

func NewManager() (Manager, error) {
	return &manager{}, nil
}
