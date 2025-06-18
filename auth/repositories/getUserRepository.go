package repositories

import (
	"my-go-task/dbModels"
)

// Actual method to fetch users from the DB using GORM.
func (r *UserRepository) GetAllUsers() ([]dbModels.User, error) {
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
