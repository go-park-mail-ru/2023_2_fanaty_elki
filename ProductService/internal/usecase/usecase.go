package usecase

import (
	//"ProductService/entity"
	productRep "ProductService/internal/repository"
	product "ProductService/proto"
)

type ProductUsecaseI interface {
	GetProductsByMenuTypeId(grpcid *product.ID) (*product.ProductSlice, error)
	GetProductByID(grpcid *product.ID) (*product.Product, error)
	SearchProducts(grpcword *product.Word) (*product.ProductSlice, error)
	GetRestaurantIdByProduct(grpcid *product.ID) (*product.ID, error)
}

type productUsecase struct {
	productRepo productRep.ProductRepositoryI
}

func NewProductUsecase(productRep productRep.ProductRepositoryI) *productUsecase {
	return &productUsecase{
		productRepo: productRep,
	}
}

func (pu productUsecase) GetProductsByMenuTypeId(grpcid *product.ID) (*product.ProductSlice, error) {
	id := uint(grpcid.ID)

	products, err := pu.productRepo.GetProductsByMenuTypeId(id)

	if err != nil {
		return nil, err
	}

	productSlice := &product.ProductSlice{}

	for _, entproduct := range products {
		grpcproduct := &product.Product{
			ID:          uint64(entproduct.ID),
			Name:        entproduct.Name,
			Price:       entproduct.Price,
			CookingTime: int64(entproduct.CookingTime),
			Portion:     entproduct.Portion,
			Description: entproduct.Description,
			Icon:        entproduct.Icon,
		}
		productSlice.Products = append(productSlice.Products, grpcproduct)
	}

	return productSlice, err
}

func (pu productUsecase) GetProductByID(grpcid *product.ID) (*product.Product, error) {
	id := uint(grpcid.ID)

	entproduct, err := pu.productRepo.GetProductByID(id)

	if err != nil {
		return nil, err
	}

	grpcproduct := &product.Product{
		ID:          uint64(entproduct.ID),
		Name:        entproduct.Name,
		Price:       entproduct.Price,
		CookingTime: int64(entproduct.CookingTime),
		Portion:     entproduct.Portion,
		Description: entproduct.Description,
		Icon:        entproduct.Icon,
	}

	return grpcproduct, nil
}

func (pu productUsecase) SearchProducts(grpcword *product.Word) (*product.ProductSlice, error) {
	word := grpcword.Word

	products, err := pu.productRepo.SearchProducts(word)

	if err != nil {
		return nil, err
	}

	productSlice := &product.ProductSlice{}

	for _, entproduct := range products {
		grpcproduct := &product.Product{
			ID:          uint64(entproduct.ID),
			Name:        entproduct.Name,
			Price:       entproduct.Price,
			CookingTime: int64(entproduct.CookingTime),
			Portion:     entproduct.Portion,
			Description: entproduct.Description,
			Icon:        entproduct.Icon,
		}
		productSlice.Products = append(productSlice.Products, grpcproduct)
	}

	return productSlice, nil

}

func (pu productUsecase) GetRestaurantIdByProduct(grpcid *product.ID) (*product.ID, error) {
	id := uint(grpcid.ID)

	restid, err := pu.productRepo.GetRestaurantIdByProduct(id)

	if err != nil {
		return nil, err
	}

	grpcrestid := &product.ID{ID: uint64(restid)}

	return grpcrestid, nil
}
