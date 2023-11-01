package repository

import (
	"server/internal/domain/entity"
	"server/internal/domain/dto"
)

type UserRepositoryI interface{
	GetUserById(id uint) (*entity.User, error) 
	CreateUser(user *dto.DBCreateUser) (uint, error)
	FindUserBy(field string, value string) (*entity.User, error)
	UpdateUser(user *dto.DBUpdateUser) (error)
}