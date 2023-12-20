package repository

import (
	"AuthService/dto"
	"AuthService/entity"
)

//SessionRepositoryI interface
type SessionRepositoryI interface {
	Create(cookie *entity.Cookie) error
	Check(sessionToken string) (*entity.Cookie, error)
	Delete(cookie *dto.DBDeleteCookie) error
	Expire(cookie *entity.Cookie) error
	CreateCsrf(sessionToken string, csrfToken string) error
	GetCsrf(sessionToken string) (string, error)
}
