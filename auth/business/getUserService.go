package services

import (
	"my-go-task/dbModels"
)

// Calls the repo method to get users.
func (s *UserService) GetAllUsersService() ([]dbModels.User, error) {
	return s.repo.GetAllUsers()
}
