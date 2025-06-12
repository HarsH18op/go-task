package dbModels

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model            // Embedded struct that provides ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Name       string     `gorm:"not null"` // By default all fields are nullable
	Email      string     `gorm:"not null"`
	Age        uint       `gorm:"default:18"` // Nullable
	Pancard    string     `gorm:"unique;not null;size:10"`
	Mobile     string     `gorm:"unique;not null;;size:10;column:mobile_number"` // This json tag is required for accepting this field as given key in json
	Bio        string     `gorm:"default:null"`                                  // Nullable
	Birthday   *time.Time `gorm:"type:date;default:null"`                        // Use pointer to allow null values in this field                                           // Nullable
}

// Override default table name "users"
func (User) TableName() string {
	return "user_details"
}