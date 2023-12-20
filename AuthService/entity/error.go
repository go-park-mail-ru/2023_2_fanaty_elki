package entity

import "errors"

//ErrNotFound errors
var (
	ErrNotFound              = errors.New("item is not found")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrInvalidUsername       = errors.New("invalid username")
	ErrInvalidEmail          = errors.New("invalid email")
	ErrInvalidPhoneNumber    = errors.New("invalid phone number")
	ErrInvalidBirthday       = errors.New("invalid birthday")
	ErrInvalidIcon           = errors.New("invalid icon")
	ErrConflictUsername      = errors.New("username already exists")
	ErrConflictEmail         = errors.New("email already exists")
	ErrConflictPhoneNumber   = errors.New("phone number already exists")
	ErrBadRequest            = errors.New("bad request")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrInternalServerError   = errors.New("internal server error")
	ErrPermissionDenied      = errors.New("permission denied")
	ErrProblemsReadingData   = errors.New("problems with reading data")
	ErrUnmarshalingJSON      = errors.New("problems with unmarshaling json")
	ErrBadContentType        = errors.New("bad content-type")
	ErrCreatingCookie        = errors.New("problems with creating cookie")
	ErrDeletingCookie        = errors.New("problems with deleting cookie")
	ErrFailCSRF              = errors.New("csrf failed")
	StatusFailCSRF           = 419
	StatusConflicUsername    = 491
	StatusConflicEmail       = 492
	StatusConflicPhoneNumber = 493
)
