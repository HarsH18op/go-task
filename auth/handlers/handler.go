package handlers

import services "my-go-task/auth/business"

// Struct that holds a reference to the service.
type UserHandler struct {
	service *services.UserService
}

// Constructor to create a handler using the given service.
func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}
