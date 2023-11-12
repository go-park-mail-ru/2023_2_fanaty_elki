package usecase

import (
	"fmt"
	productRep "server/internal/Product/repository"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetProductByID(id uint) (*entity.Product, error)
}

type productUsecase struct {
	productRepo productRep.ProductRepositoryI
}

func NewProductUsecase(productRep productRep.ProductRepositoryI) *productUsecase {
	return &productUsecase{
		productRepo: productRep,
	}
}

func (pu productUsecase) GetProductByID(id uint) (*entity.Product, error) {
	product, err := pu.productRepo.GetProductByID(id)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	if product == nil {
		fmt.Println("no product")
		return nil, entity.ErrNotFound
	}
	return product, nil
}