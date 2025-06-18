package models

type CreateWatchlistRequestModel struct {
	Name   string `json:"name" validate:"required" example:"Range Rover"`
	UserId uint   `json:"userId" validate:"required" example:"1"`
}

type CreateWatchlistResponseModel struct {
	ID      uint   `json:"watchlistId" example:"1"`
	Message string `json:"message" example:"Watchlist created"`
}
