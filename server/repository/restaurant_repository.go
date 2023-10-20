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

func (repo *RestaurantRepo) GetRestaurants() ([]*entity.Restaurant, error) {

	repo.mu.RLock()
	defer repo.mu.RUnlock()

	rows, err := repo.DB.Query("SELECT id, name, rating, comments_count, category, icon FROM restaurant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*entity.Restaurant{}
	for rows.Next() {
		restaurant := &entity.Restaurant{}
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
