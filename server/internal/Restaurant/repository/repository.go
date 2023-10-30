package repository

import "server/internal/domain/entity"

type RestaurantRepositoryI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
}
