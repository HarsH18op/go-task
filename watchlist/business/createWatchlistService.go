package services

import (
	"my-go-task/dbModels"
	constants "my-go-task/watchlist/commons"
	"my-go-task/watchlist/models"
)

func (s *WatchlistService) CreateWatchlistService(req models.CreateWatchlistRequestModel) (*dbModels.Watchlist, error) {
	// Validate UserID before creating watchlist
	exists, err := dbModels.UserExists(s.repo.Db, req.UserId) // assuming repo has db *gorm.DB
	if err != nil {
		return nil, constants.WATCHLIST_CREATION_ERRORS.UserExistenceError
	}
	if !exists {
		return nil, constants.WATCHLIST_CREATION_ERRORS.UserNotFoundError
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
