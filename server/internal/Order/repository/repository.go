package repository

import (
	"server/internal/domain/dto"
	// "server/internal/domain/entity"
)

//OrderRepositoryI interface
type OrderRepositoryI interface {
	CreateOrder(order *dto.DBReqCreateOrder) (*dto.RespCreateOrder, error)
	UpdateOrder(order *dto.ReqUpdateOrder) error
	GetOrders(UserID uint) ([]*dto.RespGetOrder, error)
	GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error)
}
