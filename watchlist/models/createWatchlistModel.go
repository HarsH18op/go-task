package models

type CreateWatchlistRequestModel struct {
	Name   string `json:"name" validate:"required"`
	UserId uint   `json:"userId" validate:"required"`
}

type CreateWatchlistResponseModel struct {
	ID      uint   `json:"watchlistId"`
	Message string `json:"message"`
}
