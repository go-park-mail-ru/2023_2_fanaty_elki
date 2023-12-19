package usecase

import (
	productRep "ProductService/internal/repository"
	product "ProductService/proto"
)

//ProductUsecaseI interface
type ProductUsecaseI interface {
	GetProductsByMenuTypeID(grpcid *product.ID) (*product.ProductSlice, error)
	GetProductByID(grpcid *product.ID) (*product.Product, error)
	SearchProducts(grpcword *product.Word) (*product.ProductSlice, error)
	GetRestaurantIDByProduct(grpcid *product.ID) (*product.ID, error)
}

//ProductUsecase struct
type ProductUsecase struct {
	productRepo productRep.ProductRepositoryI
}

//NewProductUsecase create product usecase 
func NewProductUsecase(productRep productRep.ProductRepositoryI) *ProductUsecase {
	return &ProductUsecase{
		productRepo: productRep,
	}
}

//GetProductsByMenuTypeID gets products by menu type
func (pu ProductUsecase) GetProductsByMenuTypeID(grpcid *product.ID) (*product.ProductSlice, error) {
	id := uint(grpcid.ID)

	products, err := pu.productRepo.GetProductsByMenuTypeID(id)

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

//GetProductByID gets product by id
func (pu ProductUsecase) GetProductByID(grpcid *product.ID) (*product.Product, error) {
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

//SearchProducts searches products
func (pu ProductUsecase) SearchProducts(grpcword *product.Word) (*product.ProductSlice, error) {
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

//GetRestaurantIDByProduct gets restaurant id by product
func (pu ProductUsecase) GetRestaurantIDByProduct(grpcid *product.ID) (*product.ID, error) {
	id := uint(grpcid.ID)

	restid, err := pu.productRepo.GetRestaurantIDByProduct(id)

	if err != nil {
		return nil, err
	}

	grpcrestid := &product.ID{ID: uint64(restid)}

	return grpcrestid, nil
}
