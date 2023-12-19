package repository

import "ProductService/entity"

//ProductRepositoryI interface
type ProductRepositoryI interface {
	GetProductsByMenuTypeID(id uint) ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	SearchProducts(word string) ([]*entity.Product, error)
	GetRestaurantIDByProduct(id uint) (uint, error)
}
