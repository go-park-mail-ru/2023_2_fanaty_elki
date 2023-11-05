package dto

import (
	"server/internal/domain/entity"
	"time"
)

type ReqCreateOrder struct {
	Products []uint `json:"Products"`
	UserId uint `json:"UserID"`
}

type ReqUpdateOrder struct {
	Id uint
	Status string
}


type DBReqCreateOrder struct {
	Products *map[uint]int
	UserId uint
	Status string
	Date time.Time
}

type RespCreateOrder struct {
	Id uint			`json:"Id"`
	Status string	`json:"Status"`
	Date time.Time	`json:"Date"`
}

func ToEntityCreateOrder(order *ReqCreateOrder) *entity.Order{
	return &entity.Order{
		Status: "Wait",
		UserId: order.UserId,
		Date: time.Now(),
	}
}

func ToDBReqCreateOrder(order *entity.Order, products *map[uint]int) *DBReqCreateOrder{
	return &DBReqCreateOrder{
		UserId: order.UserId,
		Products: products,
		Status: order.Status,
		Date: order.Date,
	}
} 