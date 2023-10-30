package repository

import (
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UserRepositoryI interface {
	FindUserById(id uint) (*entity.User, error)
	CreateUser(user *dto.DBCreateUser) (uint, error)
	FindUserByUsername(value string) (*entity.User, error)
	FindUserByEmail(value string) (*entity.User, error)
	FindUserByPhone(value string) (*entity.User, error)
}
