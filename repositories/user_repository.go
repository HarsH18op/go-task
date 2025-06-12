package repositories

import (
	"my-go-task/models"
	"my-go-task/utils/postgres"

	"gorm.io/gorm"
)

// Defines a struct to hold the DB connection.
type UserRepository struct {
	db *gorm.DB
}

// Constructor function — builds the struct and stores the db in it, Used for injecting the DB from main.go.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Actual method to fetch users from the DB using GORM.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *UserRepository) CreateUser(user *models.User) error {
	return postgres.DB.Create(user).Error
}

//	func GetAllUsers() ([]models.User, error) {
//		var users []models.User
//		err := postgres.DB.Find(&users).Error
//		return users, err
//	}
// func CreateUser(user *models.User) error {
// 	return postgres.DB.Create(user).Error
// }
