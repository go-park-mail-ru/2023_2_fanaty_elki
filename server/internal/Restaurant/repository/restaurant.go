package repository

import "server/internal/domain/entity"

type RestaurantI interface {
	GetRestaurants() ([]*entity.Restaurant, error) 
}