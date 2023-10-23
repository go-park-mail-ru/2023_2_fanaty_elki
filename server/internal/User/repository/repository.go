package repository

import "server/internal/domain/entity"


type UserRepositoryI interface{
	GetUserById(id uint) (*entity.User, error)
	CreateUser(user *entity.User) (uint, error)
	FindUserBy(field string, value string) (*entity.User, error)
}