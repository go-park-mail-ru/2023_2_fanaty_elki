package usecase

import (
	productRep "server/internal/Product/repository"
	"server/internal/domain/entity"
)

//ProductUsecaseI interface
type ProductUsecaseI interface {
	GetProductByID(id uint) (*entity.Product, error)
}

//ProductUsecase struct
type ProductUsecase struct {
	productRepo productRep.ProductRepositoryI
}

//NewProductUsecase creates new product usecase 
func NewProductUsecase(productRep productRep.ProductRepositoryI) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRep,
	}
}

//GetProductByID gets product by id from repository
func (pu ProductUsecase) GetProductByID(id uint) (*entity.Product, error) {
	product, err := pu.productRepo.GetProductByID(id)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	if product == nil {
		return nil, entity.ErrNotFound
	}
	return product, nil
}
