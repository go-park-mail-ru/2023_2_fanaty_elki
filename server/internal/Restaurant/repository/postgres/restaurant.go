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
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
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
	return Restaurants, nil
}

func (repo *restaurantRepo) GetRestaurantById(id uint) (*entity.Restaurant, error) {
	restaurant := &entity.Restaurant{}
	row := repo.DB.QueryRow("SELECT id, name, rating, comments_count, category, icon FROM restaurant WHERE id = $1", id)
	err := row.Scan(
		&restaurant.ID,
		&restaurant.Name,
		&restaurant.Rating,
		&restaurant.CommentsCount,
		&restaurant.Category,
		&restaurant.Icon,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}
	return restaurant, nil
}

func (repo *restaurantRepo) GetMenuTypesByRestaurantId(id uint) ([]*entity.MenuType, error) {
	rows, err := repo.DB.Query("SELECT id, name, restaurant_id FROM menu_type WHERE restaurant_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var MenuTypes = []*entity.MenuType{}
	var count = 0
	for rows.Next() {
		count++
		menuType := &entity.MenuType{}
		err = rows.Scan(
			&menuType.ID,
			&menuType.Name,
			&menuType.RestaurantID,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		MenuTypes = append(MenuTypes, menuType)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return MenuTypes, nil
}
