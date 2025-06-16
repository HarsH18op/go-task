package repositories

import (
	"my-go-task/dbModels"
)

func (r *WatchlistRepository) CreateWatchlist(watchlist *dbModels.Watchlist) error {
	return r.Db.Create(watchlist).Error
}
