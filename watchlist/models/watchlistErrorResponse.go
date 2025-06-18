package models

type ErrorResponseModel struct {
	Message string            `json:"message" example:"invalid input data"`               // General message
	Errors  map[string]string `json:"errors,omitempty" example:"name:must be valid name"` // Optional field-level validation errors
}
