package controller

import (
	"context"
	"net/http"

	"github.com/gami/simple_arch_example/controller/build"
	"github.com/gami/simple_arch_example/model"
)

type User struct {
	user model.User
}

func NewUser(user model.User) *User {
	return &User{
		user: user,
	}
}

// GetUser processes (GET /user/{user_id})
func (c *User) GetUser(w http.ResponseWriter, r *http.Request, userId uint64) {
	ctx := context.Background()
	user, err := c.user.FindByID(ctx, userId)
	if err != nil {
		respond500(w, err)
		return
	}

	res := build.FromSchemaUser(user)
	respondOK(w, res)
}

// CreateUser processes (POST /users)
func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := c.user.Create(ctx)
	if err != nil {
		respond500(w, err)
		return
	}

	res := build.Created(id)
	respondOK(w, res)
}
