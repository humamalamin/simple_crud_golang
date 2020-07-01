package helpers

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested data is not exists
	ErrNotFound = errors.New("Your requested data is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your data already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given param is not valid")
	// ErrFieldRequired
	ErrFieldRequired = errors.New("Field is required.")
)
