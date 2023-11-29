package entity

import "time"

type Cookie struct {
	UserID       uint
	SessionToken string
	MaxAge       time.Duration
}

type DBDeleteCookie struct {
	SessionToken string
}

func ToDBDeleteCookie(cookie *Cookie) *DBDeleteCookie {
	return &DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}
}
