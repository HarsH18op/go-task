package repositories

import (
	"my-go-task/dbModels"
)

// Actual method to fetch watchlists from the DB using GORM.
func (r *WatchlistRepository) GetAllWatchlists(userID uint) ([]dbModels.Watchlist, error) {
	var watchlists []dbModels.Watchlist
	err := r.Db.Where("user_id = ?", userID).Find(&watchlists).Error
	return watchlists, err
}
