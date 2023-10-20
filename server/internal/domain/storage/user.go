package storage

import (
	"server/internal/domain/entity"
	
)


type Storage interface{
	GetOne(id string) *entity.User
	GetAll() *[]entity.User
	Create(user entity.User) 
	Delete(id string) 
}
