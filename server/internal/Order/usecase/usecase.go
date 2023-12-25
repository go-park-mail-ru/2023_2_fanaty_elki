package usecase

import (
	"fmt"
	cartRep "server/internal/Cart/repository"
	orderRep "server/internal/Order/repository"
	productRep "server/internal/Product/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"
)

//OrderUsecaseI interface
type OrderUsecaseI interface {
	GetOrders(userId uint) (*dto.RespOrders, error)
	CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error)
	UpdateOrder(reqOrder *dto.ReqUpdateOrder) error
	GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error)
}

//OrderUsecase struct
type OrderUsecase struct {
	orderRepo orderRep.OrderRepositoryI
	cartRepo  cartRep.CartRepositoryI
	prodRepo  productRep.ProductRepositoryI
}

//NewOrderUsecase crates order usecase
func NewOrderUsecase(orderRepI orderRep.OrderRepositoryI, cartRepI cartRep.CartRepositoryI,
	prodRepI productRep.ProductRepositoryI) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: orderRepI,
		cartRepo:  cartRepI,
		prodRepo:  prodRepI,
	}
}

//CreateOrder creates order
func (or *OrderUsecase) CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error) {
	if len(reqOrder.Address.City) == 0 || len(reqOrder.Address.Street) == 0 || len(reqOrder.Address.House) == 0 {
		return nil, entity.ErrBadRequest
	}

	order := dto.ToEntityCreateOrder(reqOrder)

	cart, err := or.cartRepo.GetCartByUserID(reqOrder.UserID)
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

	go func() {
		upOrd := &dto.ReqUpdateOrder{
			Status: 1,
			ID: respOrder.ID,
		}
		DelTime := time.Duration(order.DeliveryTime) * time.Minute
		curTime :=  DelTime / 100 * 10
		time.Sleep(curTime)
		err := or.UpdateOrder(upOrd)
		if err != nil {
			fmt.Println("ПАЛУНДРА") 
		}

		upOrd.Status = 2
		time.Sleep(DelTime - curTime)
		err = or.UpdateOrder(upOrd)
		if err != nil {
			fmt.Println("СВИСТАТЬ ВСЕХ НАВЕРХ")
		}
	}()

	return respOrder, nil
}

//UpdateOrder updates order
func (or *OrderUsecase) UpdateOrder(reqOrder *dto.ReqUpdateOrder) error {
	err := or.orderRepo.UpdateOrder(reqOrder)
	if err != nil {
		return err
	}
	return nil
}

//GetOrders gets orders
func (or *OrderUsecase) GetOrders(userId uint) (*dto.RespOrders, error) {
	orders, err := or.orderRepo.GetOrders(userId)

	var respOrders dto.RespOrders

	for _, order := range orders {
		respOrders = append(respOrders, order)
	}
	return &respOrders, err
}

//GetOrder gets order
func (or *OrderUsecase) GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error) {
	return or.orderRepo.GetOrder(reqOrder)
}
