package repository

import (
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

//RestaurantRepositoryI interface
type RestaurantRepositoryI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
	GetRestaurantByID(id uint) (*entity.Restaurant, error)
	GetRestaurantByName(name string) (*entity.Restaurant, error)
	GetMenuTypesByRestaurantID(id uint) ([]*entity.MenuType, error)
	GetCategoriesByRestaurantID(id uint) ([]*entity.Category, error)
	GetRestaurantsByCategory(name string) ([]*entity.Restaurant, error)
	GetCategories() ([]*entity.Category, error)
	SearchRestaurants(word string) ([]*entity.Restaurant, error)
	SearchCategories(word string) ([]*entity.Restaurant, error)
	UpdateComments(comment *dto.ReqCreateComment) error
}
