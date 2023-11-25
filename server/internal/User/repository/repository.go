package repository

import (
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	// "server/internal/domain/entity"
)

type UserRepositoryI interface {
	FindUserById(id uint) (*dto.DBGetUser, error)
	CreateUser(user *dto.DBCreateUser) (uint, error)
	UpdateUser(user *dto.DBUpdateUser) error
	FindUserByUsername(value string) (*dto.DBGetUser, error)
	FindUserByEmail(value string) (*dto.DBGetUser, error)
	FindUserByPhone(value string) (*dto.DBGetUser, error)
	GetAdminById(id uint) (*entity.Admin, error)
	GetAdminByUsername(username string) (*entity.Admin, error)
}
