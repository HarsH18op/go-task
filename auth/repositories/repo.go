package repositories

import "gorm.io/gorm"

// Defines a struct to hold the DB connection.
type UserRepository struct {
	db *gorm.DB
}

// Constructor function — builds the struct and stores the db in it, Used for injecting the DB from main.go.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}