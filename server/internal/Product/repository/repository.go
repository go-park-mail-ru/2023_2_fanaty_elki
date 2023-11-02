package repository

import "server/internal/domain/entity"

type ProductRepositoryI interface {
	GetProductsByMenuTypeId(id uint) ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
}
