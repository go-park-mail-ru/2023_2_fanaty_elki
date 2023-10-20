package repository

import (
	"database/sql"
	"sync"
)

type Restaurant struct {
	ID            uint    `json:"ID"`
	Name          string  `json:"Name"`
	Rating        float32 `json:"Rating"`
	CommentsCount int     `json:"CommentsCount"`
	Icon          string  `json:"Icon"`
	Category      string  `json:"Category"`
}

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

func (repo *RestaurantRepo) GetRestaurants() ([]*Restaurant, error) {

	repo.mu.RLock()
	defer repo.mu.RUnlock()

	rows, err := repo.DB.Query("SELECT id, name, rating, comments_count, category, icon FROM restaurant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*Restaurant{}
	for rows.Next() {
		restaurant := &Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Category,
			&restaurant.Icon,
		)
		if err != nil {
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Restaurants, nil
}
