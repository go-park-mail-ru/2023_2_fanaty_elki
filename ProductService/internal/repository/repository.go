package repository

import "ProductService/entity"

type ProductRepositoryI interface {
	GetProductsByMenuTypeId(id uint) ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	SearchProducts(word string) ([]*entity.Product, error)
	GetRestaurantIdByProduct(id uint) (uint, error)
}
