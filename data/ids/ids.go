package ids

import (
	"github.com/google/uuid"
)

var (
	spec    = uuid.NewSHA1(uuid.NameSpaceURL, []byte("spec.envt.io"))
	tfstate = uuid.NewSHA1(uuid.NameSpaceURL, []byte("tfstate.envt.io"))
)

func generate(id uuid.UUID, parts ...string) uuid.UUID {
	b := id
	for _, p := range parts {
		b = uuid.NewSHA1(b, []byte(p))
	}
	return b
}

func SpecId(externalType string, externalId uuid.UUID) uuid.UUID {
	return generate(spec, externalType, externalId.String())
}

func TFStateId(deploymentId uuid.UUID, key string) uuid.UUID {
	return generate(tfstate, deploymentId.String(), key)
}
