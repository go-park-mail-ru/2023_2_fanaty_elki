package service

import (
	"server/internal/domain/entity"
)

type UserStorage interface{
	GetOne(id string) *entity.User
	GetAll() *[]entity.User
	Create(user entity.User) 
	Delete(id string) 
}

type UserService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *UserService {
	return &UserService{storage: storage}
}

	