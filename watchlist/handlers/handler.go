package handlers

import "my-go-task/watchlist/services"

// Struct that holds a reference to the service.
type WatchlistHandler struct {
	service *services.WatchlistService
}

// Constructor to create a handler using the given service.
func NewWatchlistHandler(service *services.WatchlistService) *WatchlistHandler {
	return &WatchlistHandler{service: service}
}
