package services

import (
	"my-go-task/auth/repositories"
)

// Struct that holds a pointer to the repository.
type UserService struct {
	repo *repositories.UserRepository
}

// Constructor that receives the repo and builds the service struct.
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}
