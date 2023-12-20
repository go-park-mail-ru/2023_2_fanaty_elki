package entity

import "time"

//Cookie struct
type Cookie struct {
	UserID       uint
	SessionToken string
	MaxAge       time.Duration
}

//DBDeleteCookie struct
type DBDeleteCookie struct {
	SessionToken string
}

//ToDBDeleteCookie transforms cookie to DBDeleteCookie 
func ToDBDeleteCookie(cookie *Cookie) *DBDeleteCookie {
	return &DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}
}
