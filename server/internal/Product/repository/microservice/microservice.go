package microservice

import (
	"context"

	productRep "server/internal/Product/repository"
	"server/internal/domain/entity"
	product "server/proto/product"
)

type ProductMicroService struct {
	client product.ProductRPCClient
}

func NewProductMicroService(client product.ProductRPCClient) productRep.ProductRepositoryI {
	return &ProductMicroService{
		client: client,
	}
}

func (pm *ProductMicroService) GetProductsByMenuTypeId(id uint) ([]*entity.Product, error) {
	ctx := context.Background()

	grpcid := product.ID{ID: uint64(id)}

	grpcproducts, err := pm.client.GetProductsByMenuTypeId(ctx, &grpcid)

	if err != nil {
		return nil, err
	}

	products := []*entity.Product{}

	for _, grpcproduct := range grpcproducts.Products {
		entproduct := &entity.Product{
			ID:          uint(grpcproduct.ID),
			Name:        grpcproduct.Name,
			Price:       grpcproduct.Price,
			CookingTime: int(grpcproduct.CookingTime),
			Portion:     grpcproduct.Portion,
			Description: grpcproduct.Description,
			Icon:        grpcproduct.Icon,
		}

		products = append(products, entproduct)
	}

	return products, nil

}

func (pm *ProductMicroService) GetProductByID(id uint) (*entity.Product, error) {
	ctx := context.Background()

	grpcid := product.ID{ID: uint64(id)}

	grpcproduct, err := pm.client.GetProductByID(ctx, &grpcid)

	if err != nil {
		return nil, err
	}

	entproduct := &entity.Product{
		ID:          uint(grpcproduct.ID),
		Name:        grpcproduct.Name,
		Price:       grpcproduct.Price,
		CookingTime: int(grpcproduct.CookingTime),
		Portion:     grpcproduct.Portion,
		Description: grpcproduct.Description,
		Icon:        grpcproduct.Icon,
	}

	return entproduct, nil
}

func (pm *ProductMicroService) SearchProducts(word string) ([]*entity.Product, error) {
	ctx := context.Background()

	grpcword := product.Word{Word: word}

	grpcproducts, err := pm.client.SearchProducts(ctx, &grpcword)

	if err != nil {
		return nil, err
	}

	products := []*entity.Product{}

	for _, grpcproduct := range grpcproducts.Products {
		entproduct := &entity.Product{
			ID:          uint(grpcproduct.ID),
			Name:        grpcproduct.Name,
			Price:       grpcproduct.Price,
			CookingTime: int(grpcproduct.CookingTime),
			Portion:     grpcproduct.Portion,
			Description: grpcproduct.Description,
			Icon:        grpcproduct.Icon,
		}

		products = append(products, entproduct)
	}

	return products, nil

}

func (pm *ProductMicroService) GetRestaurantIdByProduct(id uint) (uint, error) {
	ctx := context.Background()

	grpcid := product.ID{ID: uint64(id)}

	grpcrestid, err := pm.client.GetRestaurantIdByProduct(ctx, &grpcid)

	if err != nil {
		return 0, err
	}

	restid := grpcrestid.ID

	return uint(restid), nil
}
