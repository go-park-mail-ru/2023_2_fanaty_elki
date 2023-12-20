package entity

import "time"

//Cookie entity
type Cookie struct {
	UserID       uint
	SessionToken string
	MaxAge       time.Duration
}
