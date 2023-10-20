package repository

import (
	"database/sql"
	"sync"
	"server/internal/domain/entity"
)

type RestaurantRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		mu: sync.RWMutex{},
		DB: db,
	}
}

func (repo *RestaurantRepo) GetRestaurants() ([]*entity.Restaurant, error){
	return []*entity.Restaurant{}, nil
}