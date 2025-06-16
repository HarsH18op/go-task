package dbModels

import (
	"gorm.io/gorm"
)

type Watchlist struct {
	gorm.Model        // Embedded struct that provides ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"not null"`          // By default all fields are nullable
	User       User   `gorm:"foreignKey:UserID"` // Association
	UserID     uint   `gorm:"not null"`
	// OwnerID uint // Custom foreign key name
	ScriptCount      uint              `gorm:"default:null"`
	WatchlistScripts []WatchlistScript `gorm:"foreignKey:WatchlistID;constraint:OnDelete:CASCADE"` // One-to-many
}

type WatchlistScript struct {
	gorm.Model            // Embedded struct that provides ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string    `gorm:"not null"` // By default all fields are nullable
	Watchlist   Watchlist // Association
	WatchlistID uint      `gorm:"not null"`
}
