package constants

import "errors"

type watchlistCreationConstants struct {
	InvalidInputError     error
	ValidationFailedError error
	InternalServerError   error
	UserExistenceError    error
	UserNotFoundError     error
}

var WATCHLIST_CREATION_ERRORS = watchlistCreationConstants{
	InvalidInputError:     errors.New("invalid input"),
	ValidationFailedError: errors.New("validation failed"),
	InternalServerError:   errors.New("failed to create watchlist, Internal server error"),
	UserExistenceError:    errors.New("error checking user existence"),
	UserNotFoundError:     errors.New("user with given ID does not exist"),
}

func IsClientError(err error) bool {
	return errors.Is(err, WATCHLIST_CREATION_ERRORS.UserNotFoundError) ||
		errors.Is(err, WATCHLIST_CREATION_ERRORS.UserExistenceError)
}
