package repository

import (
	"server/internal/domain/dto"
	// "server/internal/domain/entity"
)

type OrderRepositoryI interface {
	CreateOrder(order *dto.DBReqCreateOrder) (*dto.RespCreateOrder, error)
	UpdateOrder(order *dto.ReqUpdateOrder) (error)
	GetOrders(userId uint) ([]*dto.RespGetOrder, error)
	GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error)
}