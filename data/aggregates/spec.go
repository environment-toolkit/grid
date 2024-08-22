package aggregates

import (
	"context"
	"errors"
	"time"

	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/events"
	"github.com/environment-toolkit/grid/data/models"
	"gopkg.in/yaml.v3"

	"github.com/go-apis/eventsourcing/es"
)

var ErrSpecNameExists = errors.New("spec name exists")
var ErrSpecNameRequired = errors.New("spec name is required")

type SpecName struct {
	Name string `yaml:"name" required:"true"`
}

type Spec struct {
	es.BaseAggregateSourced

	specQuery es.Query[*Spec]

	State     models.SpecState  `json:"state" required:"true"`
	Name      string            `json:"name" required:"true"`
	Content   []byte            `json:"content" required:"true" gorm:"type:jsonb;serializer:json"`
	Variables map[string]string `json:"variables" required:"true" gorm:"type:jsonb;serializer:json"`

	CreatedAt time.Time  `json:"created_at" required:"true"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	SyncedAt  *time.Time `json:"synced_at"`
}

func (a *Spec) HandleNewSpec(ctx context.Context, cmd *commands.NewSpec) error {
	var spec SpecName
	if err := yaml.Unmarshal(cmd.Content, &spec); err != nil {
		return err
	}
	if spec.Name == "" {
		return ErrSpecNameRequired
	}

	filter := es.Filter{
		Where: []es.WhereClause{
			{
				Column: "id",
				Op:     es.OpNotEqual,
				Args:   cmd.AggregateId,
			}, {
				Column: "name",
				Op:     es.OpEqual,
				Args:   spec.Name,
			},
		},
	}
	found, err := a.specQuery.Find(ctx, filter)
	if err != nil {
		return err
	}
	if len(found) > 0 {
		return ErrSpecNameExists
	}

	return a.Apply(ctx, &events.SpecCreated{
		State:     models.SpecStateCreated,
		Name:      spec.Name,
		Content:   cmd.Content,
		Variables: cmd.Variables,
		CreatedAt: time.Now(),
	})
}

func (a *Spec) HandleDeleteSpec(ctx context.Context, cmd *commands.DeleteSpec) error {
	return a.Apply(ctx, &events.SpecDeleted{
		State:     models.SpecStateDeleted,
		DeletedAt: time.Now(),
	})
}

func NewSpec() *Spec {
	specQuery := es.NewQuery[*Spec]()

	return &Spec{
		specQuery: specQuery,

		State: models.SpecStateNew,
	}
}
