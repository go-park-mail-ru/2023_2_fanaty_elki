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
	var res = entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.jpg",
		Category:      "Fastfood",
	}
	return []*entity.Restaurant{&res}, nil
}