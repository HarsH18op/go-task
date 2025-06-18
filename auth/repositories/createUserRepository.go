package repositories

import (
	"my-go-task/dbModels"
)

func (r *UserRepository) CreateUser(user *dbModels.User) error {
	return r.db.Create(user).Error
}
