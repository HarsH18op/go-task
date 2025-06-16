package models

type CreateWatchlistRequestModel struct {
	Name   string `json:"name" validate:"required"`
	UserId uint   `json:"user_id" validate:"required"`
}

type CreateWatchlistResponseModel struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	UserId      uint   `json:"user_id"`
	ScriptCount uint   `json:"script_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
