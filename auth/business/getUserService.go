package services

import (
	"my-go-task/auth/repositories"
	"my-go-task/dbModels"
)

// Struct that holds a pointer to the repository.
type GetUserService struct {
	repo *repositories.GetUserRepository
}

// Constructor that receives the repo and builds the service struct.
func NewGetUserService(repo *repositories.GetUserRepository) *GetUserService {
	return &GetUserService{repo: repo}
}

// Calls the repo method to get users.
func (s *GetUserService) GetAllUsersService() ([]dbModels.User, error) {
	return s.repo.GetAllUsers()
}
