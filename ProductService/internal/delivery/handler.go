package delivery

import (
	productUsecase "ProductService/internal/usecase"
	product "ProductService/proto"
	"context"
)

type ProductManager struct {
	product.UnimplementedProductRPCServer
	ProductUC productUsecase.ProductUsecaseI
}

func NewProductManager(uc productUsecase.ProductUsecaseI) product.ProductRPCServer {
	return ProductManager{ProductUC: uc}
}

func (pm ProductManager) GetProductsByMenuTypeId(ctx context.Context, grpcid *product.ID) (*product.ProductSlice, error) {
	resp, err := pm.ProductUC.GetProductsByMenuTypeId(grpcid)
	return resp, err
}

func (pm ProductManager) GetProductByID(ctx context.Context, grpcid *product.ID) (*product.Product, error) {
	resp, err := pm.ProductUC.GetProductByID(grpcid)
	return resp, err
}

func (pm ProductManager) SearchProducts(ctx context.Context, grpcword *product.Word) (*product.ProductSlice, error) {
	resp, err := pm.ProductUC.SearchProducts(grpcword)
	return resp, err
}

func (pm ProductManager) GetRestaurantIdByProduct(ctx context.Context, grpcid *product.ID) (*product.ID, error) {
	resp, err := pm.ProductUC.GetRestaurantIdByProduct(grpcid)
	return resp, err
}
