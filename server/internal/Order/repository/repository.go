package repository

import (
	"server/internal/domain/dto"
	// "server/internal/domain/entity"
)

type OrderRepositoryI interface {
	CreateOrder(order *dto.DBReqCreateOrder) (*dto.RespCreateOrder, error)
}