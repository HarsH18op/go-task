package repositories

import "gorm.io/gorm"

// Defines a struct to hold the DB connection.
type WatchlistRepository struct {
	Db *gorm.DB
}

// Constructor function — builds the struct and stores the db in it, Used for injecting the DB from main.go.
func NewWatchlistRepository(db *gorm.DB) *WatchlistRepository {
	return &WatchlistRepository{Db: db}
}
