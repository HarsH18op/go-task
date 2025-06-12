package services

import (
	"my-go-task/models"
	"my-go-task/repositories"
	"time"
)

// Struct that holds a pointer to the repository.
type UserService struct {
	repo *repositories.UserRepository
}

// Constructor that receives the repo and builds the service struct.
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Calls the repo method to get users.
func (s *UserService) GetAllUsersService() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) CreateUserService(req models.CreateUserRequest) (*models.User, error) {
	// Manage nullable date field
	var birthday *time.Time
	if req.Birthday != "" {
		parsedBirthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			birthday = &parsedBirthday
		}
	}
	user := &models.User{
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