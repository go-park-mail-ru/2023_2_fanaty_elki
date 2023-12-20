package dto

import "errors"

//ErrInternalServerError 500
var (
	ErrInternalServerError   = errors.New("internal server error")
)
