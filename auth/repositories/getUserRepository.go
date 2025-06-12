package repositories

import (
	"my-go-task/dbModels"

	"gorm.io/gorm"
)

// Defines a struct to hold the DB connection.
type GetUserRepository struct {
	db *gorm.DB
}

// Constructor function — builds the struct and stores the db in it, Used for injecting the DB from main.go.
func NewGetUserRepository(db *gorm.DB) *GetUserRepository {
	return &GetUserRepository{db: db}
}

// Actual method to fetch users from the DB using GORM.
func (r *GetUserRepository) GetAllUsers() ([]dbModels.User, error) {
	var users []dbModels.User
	err := r.db.Find(&users).Error
	return users, err
}

// If implemented in normal way
//	func GetAllUsers() ([]models.User, error) {
//		var users []models.User
//		err := postgres.DB.Find(&users).Error
//		return users, err
//	}
// func CreateUser(user *models.User) error {
// 	return postgres.DB.Create(user).Error
// }
