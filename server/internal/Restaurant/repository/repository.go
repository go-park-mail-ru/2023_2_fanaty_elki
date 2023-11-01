package repository

import "server/internal/domain/entity"

type RestaurantRepositoryI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
	GetRestaurantById(id uint) (*entity.Restaurant, error)
	GetMenuTypesByRestaurantId(id uint) ([]*entity.MenuType, error)
}
