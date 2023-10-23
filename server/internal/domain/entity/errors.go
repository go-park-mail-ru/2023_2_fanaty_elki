package entity

import "errors"

var (
	ErrNotFound            = errors.New("item is not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidUsername     = errors.New("invalid username")
	ErrInvalidEmail		   = errors.New("invalid email")
	ErrInvalidPhoneNumber  = errors.New("invalid phone number")
	ErrInvalidBirthday	   = errors.New("invalid birthday")
	ErrConflictUsername    = errors.New("username already exists")
	ErrConflictEmail       = errors.New("email already exists")
	ErrConflictPhoneNumber = errors.New("phone number already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
	ErrPermissionDenied    = errors.New("permission denied")
	ErrProblemsReadingData = errors.New("problems with reading data")
	ErrUnmarshalingJson	   = errors.New("problems with unmarshaling json")
)