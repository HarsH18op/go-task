package dbModels

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model             // Embedded struct that provides ID, CreatedAt, UpdatedAt, DeletedAt
	ID         uint        `gorm:"primaryKey;autoIncrement"`
	Name       string      `gorm:"not null"` // By default all fields are nullable
	Email      string      `gorm:"not null"`
	Pancard    string      `gorm:"unique;not null;size:10"`
	Mobile     string      `gorm:"unique;not null;size:10;column:mobile_number"` // This json tag is required for accepting this field as given key in json
	Age        uint        `gorm:"default:18"`                                   // Nullable
	Bio        string      `gorm:"default:null"`                                 // Nullable
	Birthday   *time.Time  `gorm:"type:date;default:null"`                       // Use pointer to allow null values in this field
	Watchlists []Watchlist `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	// Watchlists []Watchlist `gorm:"foreignKey:OwnerID"` // Uses custom fk -> OwnerID instead of UserID
	// Watchlists []Watchlist `gorm:"foreignKey:UserUUID;references:Mobile"` // references a field other than ID
	// Watchlists []Watchlist `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// Options: CASCADE, SET NULL, SET DEFAULT, RESTRICT, NO ACTION

}

// Override default table name "users"
func (User) TableName() string {
	return "user_details"
}

func UserExists(db *gorm.DB, userID uint) (bool, error) {
	var count int64
	err := db.Model(&User{}).Where("id = ?", userID).Count(&count).Error
	log.Println(count, "************")
	if err != nil {
		return false, err
	}
	return count > 0, nil
}