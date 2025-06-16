package services

import (
	"fmt"
	"my-go-task/dbModels"
	"my-go-task/watchlist/models"
)

func (s *WatchlistService) CreateWatchlistService(req models.CreateWatchlistRequestModel) (*dbModels.Watchlist, error) {
	// Validate UserID before creating watchlist
	exists, err := dbModels.UserExists(s.repo.Db, req.UserId) // assuming repo has db *gorm.DB
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("user with ID %d does not exist", req.UserId)
	}

	watchlist := &dbModels.Watchlist{
		Name:        req.Name,
		UserID:      req.UserId,
		ScriptCount: 1,
	}
	err2 := s.repo.CreateWatchlist(watchlist)
	if err2 != nil {
		return nil, err2
	}
	return watchlist, nil
}
