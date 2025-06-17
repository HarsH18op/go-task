package models

type ErrorResponseModel struct {
	Message string            `json:"message"`          // General message
	Errors  map[string]string `json:"errors,omitempty"` // Optional field-level validation errors, omit if not exists in payload
}
