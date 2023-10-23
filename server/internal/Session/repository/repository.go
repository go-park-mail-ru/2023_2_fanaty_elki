package repository

import (
	"server/internal/domain/entity"
)

type SessionRepositoryI interface {
	Create(cookie *entity.Cookie) error 
	Check(sessionToken string) (*entity.Cookie, error)
	Delete(cookie *entity.Cookie) error
}