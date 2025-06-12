package repositories

import (
	"my-go-task/dbModels"

	"gorm.io/gorm"
)

// Defines a struct to hold the DB connection.
type CreateUserRepository struct {
	db *gorm.DB
}

// Constructor function — builds the struct and stores the db in it, Used for injecting the DB from main.go.
func NewCreateUserRepository(db *gorm.DB) *CreateUserRepository {
	return &CreateUserRepository{db: db}
}

func (r *CreateUserRepository) CreateUser(user *dbModels.User) error {
	return r.db.Create(user).Error
}
