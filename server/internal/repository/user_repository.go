package repository

import (
	"database/sql"
	"sync"
	"server/internal/domain/entity"
	)


type UserRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		mu: sync.RWMutex{},
		DB: db,
	}
}


func (repo *UserRepo) GetUserById(id uint) *entity.User {
	return &entity.User{}
}

func (repo *UserRepo) Create(user *entity.User) error {
	return nil
}

func (repo *UserRepo) FindUserBy(field string, value string) (*entity.User, error) {
	return &entity.User{}, nil
}