// Package User provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package User

// Created defines model for Created.
type Created struct {
	Id      uint64 `json:"id"`
	Message string `json:"message"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// User defines model for User.
type User struct {

	// identifier
	Id uint64 `json:"id"`

	// name
	Name    string `json:"name"`
	Profile Error  `json:"profile"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {

	// name
	Name string `json:"name"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

