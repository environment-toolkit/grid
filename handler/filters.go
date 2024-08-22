package handler

import (
	"github.com/environment-toolkit/grid/data/models"

	"github.com/go-apis/utils/xes"
	"github.com/google/uuid"
)

type PagedEnvironmentsInput struct {
	xes.BasePagingInput

	Ids    *[]uuid.UUID               `query:"ids" where:"id,in"`
	States *[]models.EnvironmentState `query:"states" where:"state,in"`
	Names  *[]string                  `query:"names" where:"name,in"`
}

type PagedSpecsInput struct {
	xes.BasePagingInput

	Ids    *[]uuid.UUID        `query:"ids" where:"id,in"`
	States *[]models.SpecState `query:"states" where:"state,in"`
	Names  *[]string           `query:"names" where:"name,in"`
}

type FindDeploymentsInput struct {
	xes.BaseFindInput

	Ids *[]uuid.UUID `query:"ids" where:"id,in"`

	SpecIds        *[]uuid.UUID `query:"spec_ids"  where:"spec_id,in"`
	EnvironmentIds *[]uuid.UUID `query:"environment_ids" where:"environment_id,in"`
}
