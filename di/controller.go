package di

import "github.com/gami/simple_arch_example/controller"

func InjectUserController() *controller.User {
	return controller.NewUser(
		InjectUserModel(),
	)
}
