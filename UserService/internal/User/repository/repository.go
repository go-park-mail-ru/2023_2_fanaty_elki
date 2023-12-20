package repository

import (
	"UserService/internal/dto"
)

//UserRepositoryI interface
type UserRepositoryI interface {
	FindUserByID(id uint) (*dto.DBGetUser, error)
	CreateUser(user *dto.DBCreateUser) (uint, error)
	UpdateUser(user *dto.DBUpdateUser) error
	FindUserByUsername(value string) (*dto.DBGetUser, error)
	FindUserByEmail(value string) (*dto.DBGetUser, error)
	FindUserByPhone(value string) (*dto.DBGetUser, error)
}
