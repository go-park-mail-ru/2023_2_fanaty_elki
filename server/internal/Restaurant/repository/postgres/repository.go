package repository

import (
	"context"
	"server/internal/domain/entity"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	getRestaurantsList = "SELECT id, name, rating, comments_count, category, icon FROM restaurant"
)

type restaurantRepo struct {
	DB *pgxpool.Pool
}

func NewRestaurantRepo(db *pgxpool.Pool) *restaurantRepo {
	return &restaurantRepo{
		DB: db,
	}
}

func (repo *restaurantRepo) GetRestaurants() ([]*entity.Restaurant, error) {

	rows, err := repo.DB.Query(context.Background(), getRestaurantsList)
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
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Restaurants, nil
}
