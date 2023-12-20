package delivery

import (
	productUsecase "ProductService/internal/usecase"
	product "ProductService/proto"
	"context"
)

//ProductManager struct
type ProductManager struct {
	product.UnimplementedProductRPCServer
	ProductUC productUsecase.ProductUsecaseI
}

//NewProductManager creates product rpc server
func NewProductManager(uc productUsecase.ProductUsecaseI) product.ProductRPCServer {
	return ProductManager{ProductUC: uc}
}

//GetProductsByMenuTypeID handles get product by menu type id request 
func (pm ProductManager) GetProductsByMenuTypeID(ctx context.Context, grpcid *product.ID) (*product.ProductSlice, error) {
	resp, err := pm.ProductUC.GetProductsByMenuTypeID(grpcid)
	return resp, err
}

//GetProductByID handles get product by id request
func (pm ProductManager) GetProductByID(ctx context.Context, grpcid *product.ID) (*product.Product, error) {
	resp, err := pm.ProductUC.GetProductByID(grpcid)
	return resp, err
}

//SearchProducts handles search products request
func (pm ProductManager) SearchProducts(ctx context.Context, grpcword *product.Word) (*product.ProductSlice, error) {
	resp, err := pm.ProductUC.SearchProducts(grpcword)
	return resp, err
}

//GetRestaurantIDByProduct handles get restaurant id by product request
func (pm ProductManager) GetRestaurantIDByProduct(ctx context.Context, grpcid *product.ID) (*product.ID, error) {
	resp, err := pm.ProductUC.GetRestaurantIDByProduct(grpcid)
	return resp, err
}
