package build

import (
	api "github.com/gami/simple_arch_example/gen/openapi"
	"github.com/gami/simple_arch_example/gen/schema"
)

func FromSchemaUser(m *schema.User) *api.User {
	return &api.User{
		Id: m.ID,
	}
}
