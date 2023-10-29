package repository

import (
	"server/internal/domain/entity"
	"server/internal/domain/dto"
)

type SessionRepositoryI interface {
	Create(cookie *entity.Cookie) error 
	Check(sessionToken string) (*entity.Cookie, error)
	Delete(cookie *dto.DBDeleteCookie) error
}