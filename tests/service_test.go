package tests

import (
	"context"
	"os"
	"testing"

	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/eventsourcing/es"
	"github.com/stretchr/testify/require"
)

func TestCommands(t *testing.T) {
	tester, err := NewServiceTester()
	require.NoError(t, err)

	t.Run("create environments", func(t *testing.T) {
		ctx := context.Background()
		unit, err := tester.Client().Unit(ctx)
		require.NoError(t, err)

		cmds := []es.Command{
			&commands.NewEnvironment{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: Environment1Id,
				},
				Name:  Environment1Name,
				Title: Environment1Title,
			},
		}

		errd := unit.Dispatch(ctx, cmds...)
		require.NoError(t, errd)
	})

	t.Run("create spec", func(t *testing.T) {
		ctx := context.Background()
		unit, err := tester.Client().Unit(ctx)
		require.NoError(t, err)

		content, err := os.ReadFile(Spec1File)
		require.NoError(t, err)

		cmds := []es.Command{
			&commands.NewSpec{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: Spec1Id,
				},
				Content:   content,
				Variables: Spec1Variables,
			},
			&commands.UpdateState{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: State1Id,
				},
				SpecId: Spec1Id,
				Target: models.Target{
					EnvironmentId: Environment1Id,
					Region:        "us-west-1",
				},
			},
		}

		errd := unit.Dispatch(ctx, cmds...)
		require.NoError(t, errd)
	})

	t.Run("create spec fails", func(t *testing.T) {
		ctx := context.Background()
		unit, err := tester.Client().Unit(ctx)
		require.NoError(t, err)

		content, err := os.ReadFile(Spec1File)
		require.NoError(t, err)

		cmds := []es.Command{
			&commands.NewSpec{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: Spec1Id,
				},
				Content:   content,
				Variables: Spec1Variables,
			},
			&commands.NewSpec{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: Spec2Id,
				},
				Content:   content,
				Variables: Spec1Variables,
			},
		}

		errd := unit.Dispatch(ctx, cmds...)
		require.Error(t, errd)
	})

	t.Run("update state", func(t *testing.T) {
		ctx := context.Background()
		unit, err := tester.Client().Unit(ctx)
		require.NoError(t, err)

		cmds := []es.Command{
			&commands.UpdateState{
				BaseNamespaceCommand: es.BaseNamespaceCommand{
					Namespace: OrganisationId.String(),
				},
				BaseCommand: es.BaseCommand{
					AggregateId: State1Id,
				},
				SpecId: Spec1Id,
				Target: State1Target,
			},
		}

		errd := unit.Dispatch(ctx, cmds...)
		require.Error(t, errd)
	})

}
