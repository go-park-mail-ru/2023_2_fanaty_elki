package usecase

import (
	cartRep "server/internal/Cart/repository"
	orderRep "server/internal/Order/repository"
	productRep "server/internal/Product/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetOrders(userId uint) (*dto.RespOrders, error)
	CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error)
	UpdateOrder(reqOrder *dto.ReqUpdateOrder) error
	GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error)
}

type orderUsecase struct {
	orderRepo orderRep.OrderRepositoryI
	cartRepo  cartRep.CartRepositoryI
	prodRepo  productRep.ProductRepositoryI
}

func NewOrderUsecase(orderRepI orderRep.OrderRepositoryI, cartRepI cartRep.CartRepositoryI,
	prodRepI productRep.ProductRepositoryI) *orderUsecase {
	return &orderUsecase{
		orderRepo: orderRepI,
		cartRepo:  cartRepI,
		prodRepo:  prodRepI,
	}
}

func (or *orderUsecase) CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error) {
	if len(reqOrder.Address.City) == 0 || len(reqOrder.Address.Street) == 0 || len(reqOrder.Address.House) == 0 {
		return nil, entity.ErrBadRequest
	}

	order := dto.ToEntityCreateOrder(reqOrder)

	cart, err := or.cartRepo.GetCartByUserID(reqOrder.UserId)
	if err != nil {
		return nil, err
	}

	products, err := or.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	if len(products.Products) == 0 {
		return nil, entity.ErrNotFound
	}

	for _, product := range products.Products {
		pr, err := or.prodRepo.GetProductByID(product.ProductID)
		if err != nil {
			return nil, err
		}
		if pr != nil {
			order.Price += uint(pr.Price) * uint(product.ItemCount)
		}
	}

	order.DeliveryTime = 30
	respOrder, err := or.orderRepo.CreateOrder(dto.ToDBReqCreateOrder(order, products.Products))
	respOrder.Address = order.Address

	if err != nil {
		return nil, err
	}

	return respOrder, nil
}

func (or *orderUsecase) UpdateOrder(reqOrder *dto.ReqUpdateOrder) error {
	err := or.orderRepo.UpdateOrder(reqOrder)
	if err != nil {
		return err
	}
	return nil
}

func (or *orderUsecase) GetOrders(userId uint) (*dto.RespOrders, error) {
	orders, err := or.orderRepo.GetOrders(userId)

	var respOrders dto.RespOrders

	for _, order := range orders {
		respOrders = append(respOrders, order)
	}
	return &respOrders, err
}

func (or *orderUsecase) GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error) {
	return or.orderRepo.GetOrder(reqOrder)
}
