package models

type ErrorResponseModel struct {
	Message string            `json:"message" example:"invalid input data"`          // General message
	Errors  map[string]string `json:"errors,omitempty" example:"email:must be a valid email address"` // Optional field-level validation errors
}
