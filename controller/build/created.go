package build

import (
	api "github.com/gami/simple_arch_example/gen/openapi"
)

func Created(id uint64) *api.Created {
	return &api.Created{
		Id:      id,
		Message: "Created",
	}
}
