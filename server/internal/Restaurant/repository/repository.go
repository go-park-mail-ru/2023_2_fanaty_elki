package repository

import (
	"server/internal/domain/entity"
	"server/internal/domain/dto"
)

type RestaurantRepositoryI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
	GetRestaurantById(id uint) (*entity.Restaurant, error)
	GetMenuTypesByRestaurantId(id uint) ([]*entity.MenuType, error)
	GetCategoriesByRestaurantId(id uint) ([]*entity.Category, error)
	GetRestaurantsByCategory(name string) ([]*entity.Restaurant, error)
	GetCategories() ([]*entity.Category, error)
	SearchRestaurants(word string) ([]*entity.Restaurant, error)
	SearchCategories(word string) ([]*entity.Restaurant, error)
	UpdateComments(comment *dto.ReqCreateComment) error
}
