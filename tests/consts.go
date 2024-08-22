package tests

import (
	"github.com/environment-toolkit/grid/data/models"
	"github.com/google/uuid"
)

var OrganisationId = uuid.MustParse("8de2b552-7516-482d-8ab0-eedda3d02c98")

var Environment1Id = uuid.MustParse("4e517f6c-4365-4475-8982-99a18f1d5c75")
var Environment1Name = "prod"
var Environment1Title = "production"

var Spec1Id = uuid.MustParse("928d2eb8-c95b-47b2-9fb1-d7756998a480")
var Spec1File = "specs/spec1.yml"
var Spec1Variables = map[string]string{
	"image": "nginx:latest",
}
var Spec2Id = uuid.MustParse("885e3c30-aab1-471b-829f-abed9ea7a544")

var State1Id = uuid.MustParse("64d4f7b6-eb64-444d-9af1-a4da18195d82")
var State1Target = models.Target{
	EnvironmentId: Environment1Id,
	Region:        "us-east-1",
}
