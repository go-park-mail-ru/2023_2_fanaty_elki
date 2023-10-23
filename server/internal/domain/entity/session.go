package entity

import "time"

type Cookie struct {
	UserID uint
	SessionToken string
	MaxAge       time.Duration
}