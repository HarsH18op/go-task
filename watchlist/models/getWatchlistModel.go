package models

type GetWatchlistResponseModel struct {
	ID   uint   `json:"watchlistId" example:"1"`
	Name string `json:"watchlistName" example:"Range Rover"`
}
