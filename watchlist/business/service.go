package services

import "my-go-task/watchlist/repositories"

// Struct that holds a pointer to the repository.
type WatchlistService struct {
	repo *repositories.WatchlistRepository
}

// Constructor that receives the repo and builds the service struct.
func NewWatchlistService(repo *repositories.WatchlistRepository) *WatchlistService {
	return &WatchlistService{repo: repo}
}
