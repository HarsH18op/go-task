package services

import (
	"my-go-task/auth/models"
	"my-go-task/dbModels"
	"time"
)

func (s *UserService) CreateUserService(req models.CreateUserRequestModel) (*dbModels.User, error) {
	// Manage nullable date field
	var birthday *time.Time
	if req.Birthday != "" {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			birthday = &parsedBirthday
		}
	}
	user := &dbModels.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Pancard:  req.Pancard,
		Mobile:   req.MobileNumber,
		Birthday: birthday,
		Bio:      req.Bio,
	}
	// err := repositories.CreateUser(user)
	err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// If implemented in normal way
// func GetAllUsersService() ([]models.User, error) {
// 	return repositories.GetAllUsers()
// }
