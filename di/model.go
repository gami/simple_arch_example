package di

import (
	"github.com/gami/simple_arch_example/model"
)

func InjectUserModel() model.User {
	return model.NewUser(
		InjectUserDB(),
	)
}
