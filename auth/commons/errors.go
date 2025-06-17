package constants

import "errors"

type userCreationConstants struct {
	InvalidInputError     error
	ValidationFailedError error
	InternalServerError   error
}

var USER_CREATION_ERRORS = userCreationConstants{
	InvalidInputError:     errors.New("invalid input"),
	ValidationFailedError: errors.New("validation failed"),
	InternalServerError:   errors.New("failed to create user due to internal server error"),
}
