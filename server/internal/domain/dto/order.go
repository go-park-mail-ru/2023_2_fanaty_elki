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
	Id uint			`json:"Id"`
	Status string	`json:"Status"`
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

type RespGetOrder struct {
	Id uint `json:"Id"`
	Status string `json:"Status"`
	Date time.Time `json:"Date"`
	UpdatedDate time.Time `json:"UpdatedDate"`
}

type RespGetOneOrder struct {
	Status string `json:"Status"`
	Date time.Time `json:"Date"`
	UpdatedDate time.Time `json:"UpdatedDate"`
	Products []*Product
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

func ToEntityUpdateOrder(order *ReqUpdateOrder) *entity.Order {
	return &entity.Order{
		ID: order.Id,
		Status: order.Status,
	}
}