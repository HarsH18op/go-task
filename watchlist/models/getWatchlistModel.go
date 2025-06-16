package models

type GetWatchlistResponseModel struct {
	ID          uint   `json:"watchlistId"`
	Name        string `json:"watchlistName"`
}