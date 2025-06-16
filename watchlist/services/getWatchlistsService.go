package services

import (
	"my-go-task/dbModels"
)

// Calls the repo method to get watchlists.
func (s *WatchlistService) GetWatchlistsService(userID uint) ([]dbModels.Watchlist, error) {
	return s.repo.GetAllWatchlists(userID)
}
