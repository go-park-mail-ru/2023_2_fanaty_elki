package dto

import "AuthService/entity"

type DBDeleteCookie struct {
	SessionToken string
}

func ToDBDeleteCookie(cookie *entity.Cookie) *DBDeleteCookie {
	return &DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}
}
