package repository

import (
	"database/sql"
	"server/internal/domain/entity"
)

type restaurantRepo struct {
	DB *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *restaurantRepo {
	return &restaurantRepo{
		DB: db,
	}
}

func (repo *restaurantRepo) GetRestaurants() ([]*entity.Restaurant, error) {

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

func (repo *restaurantRepo) GetRestaurantById() (*entity.Restaurant, error) {

}
