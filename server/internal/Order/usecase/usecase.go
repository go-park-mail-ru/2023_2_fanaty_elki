package usecase

import (
	// "server/internal/domain/entity"
	"server/internal/domain/dto"
	orderRep "server/internal/Order/repository"
)

type UsecaseI interface {
	GetOrders(userId uint) ([]*dto.RespGetOrder, error)
	CreateOrder(reqOrder *dto.ReqCreateOrder) (*dto.RespCreateOrder, error)
	UpdateOrder(reqOrder *dto.ReqUpdateOrder) (error)
	GetOrder(orderid uint) ()
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
		products[product]++
	}

	respOrder, err := or.orderRepo.CreateOrder(dto.ToDBReqCreateOrder(order, &products))
	if err != nil {
		return nil, err
	}

	return respOrder, nil
}

func (or *orderUsecase) UpdateOrder(reqOrder *dto.ReqUpdateOrder) (error) {
	err := or.orderRepo.UpdateOrder(reqOrder)
	if err != nil {
		return err
	}
	return nil
}

func (or *orderUsecase) GetOrders(userId uint) ([]*dto.RespGetOrder, error) {
	return or.orderRepo.GetOrders(userId)
}