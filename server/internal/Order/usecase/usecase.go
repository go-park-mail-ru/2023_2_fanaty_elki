package usecase

import (
	// "server/internal/domain/entity"
	"server/internal/domain/dto"
	orderRep "server/internal/Order/repository"
)

type UsecaseI interface {
//	GetOrder(id uint) *entity.Order
	CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error)
//	UpdateOrder()
}

type orderUsecase struct {
	orderRepo orderRep.OrderRepositoryI
}

func NewOrderUsecase(repI orderRep.OrderRepositoryI) *orderUsecase{
	return &orderUsecase{
		orderRepo: repI,
	}
}

func (or *orderUsecase) CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error) {
	order := dto.ToEntityCreateOrder(reqOrder)
	
	products := make(map[uint]int)
	for _, product := range reqOrder.Products {
		// if _, ok := products[product]; !ok {
		// 	products[product] = 0
		// }
		products[product]++
	}

	respOrder, err := or.orderRepo.CreateOrder(dto.ToDBReqCreateOrder(order, &products))
	if err != nil {
		return nil, err
	}

	return respOrder, nil
}