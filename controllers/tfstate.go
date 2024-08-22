package controllers

import (
	"context"
	"database/sql"
	"io"
	"net/http"

	"github.com/environment-toolkit/grid/data/aggregates"
	"github.com/environment-toolkit/grid/data/commands"
	"github.com/environment-toolkit/grid/data/ids"

	"github.com/go-apis/eventsourcing/es"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func getInputs(r *http.Request) (ctx context.Context, deploymentId uuid.UUID, key string, err error) {
	namespace := chi.URLParam(r, "namespace")
	deploymentIdStr := chi.URLParam(r, "id")

	ctx = es.SetNamespace(r.Context(), namespace)
	deploymentId, err = uuid.Parse(deploymentIdStr)
	key = chi.URLParam(r, "key")
	return
}

func handleError(w http.ResponseWriter, err error) {
	switch err {
	case aggregates.ErrInvalid:
		w.WriteHeader(http.StatusBadRequest)
	case aggregates.ErrLocked:
		w.WriteHeader(http.StatusLocked)
	case aggregates.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case es.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetState() http.HandlerFunc {
	query := es.NewQuery[*aggregates.TFState]()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, deploymentId, key, err := getInputs(r)
		if err != nil {
			handleError(w, err)
			return
		}

		id := ids.TFStateId(deploymentId, key)
		out, err := query.Get(ctx, id)
		if err != nil {
			handleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(out.StateFile))
		w.WriteHeader(http.StatusOK)
	}
}

func PostState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, deploymentId, key, err := getInputs(r)
		if err != nil {
			handleError(w, err)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			handleError(w, err)
			return
		}

		cmd := &commands.UpdateTFState{
			BaseCommand: es.BaseCommand{
				AggregateId: ids.TFStateId(deploymentId, key),
			},
			DeploymentId: deploymentId,
			Key:          key,
			StateFile:    string(body),
		}

		unit, err := es.GetUnit(ctx)
		if err != nil {
			handleError(w, err)
			return
		}
		if err := unit.Dispatch(ctx, cmd); err != nil {
			handleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, deploymentId, key, err := getInputs(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		cmd := &commands.DeleteTFState{
			BaseCommand: es.BaseCommand{
				AggregateId: ids.TFStateId(deploymentId, key),
			},
			DeploymentId: deploymentId,
			Key:          key,
		}

		unit, err := es.GetUnit(ctx)
		if err != nil {
			handleError(w, err)
			return
		}
		if err := unit.Dispatch(ctx, cmd); err != nil {
			handleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func LockState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, deploymentId, key, err := getInputs(r)
		if err != nil {
			handleError(w, err)
			return
		}

		cmd := &commands.LockTFState{
			BaseCommand: es.BaseCommand{
				AggregateId: ids.TFStateId(deploymentId, key),
			},
			DeploymentId: deploymentId,
			Key:          key,
		}

		unit, err := es.GetUnit(ctx)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := unit.Dispatch(ctx, cmd); err != nil {
			handleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func UnlockState() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, deploymentId, key, err := getInputs(r)
		if err != nil {
			handleError(w, err)
			return
		}

		cmd := &commands.UnlockTFState{
			BaseCommand: es.BaseCommand{
				AggregateId: ids.TFStateId(deploymentId, key),
			},
			DeploymentId: deploymentId,
			Key:          key,
		}

		unit, err := es.GetUnit(ctx)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := unit.Dispatch(ctx, cmd); err != nil {
			handleError(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
