package dto

import "AuthService/entity"

//DBDeleteCookie struct
type DBDeleteCookie struct {
	SessionToken string
}

//ToDBDeleteCookie transforms cookie to DBDeleteCookie
func ToDBDeleteCookie(cookie *entity.Cookie) *DBDeleteCookie {
	return &DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}
}
