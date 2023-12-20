package dto

import (
	"server/internal/domain/entity"
	"time"
)

type ReqCreateOrder struct {
	UserId  uint                   `json:"UserID"`
	Address *ReqCreateOrderAddress `json:"Address"`
}

type ReqUpdateOrder struct {
	Id     uint  `json:"Id"`
	Status uint8 `json:"Status"`
}

type DBReqCreateOrder struct {
	Products     []*entity.CartProduct
	UserId       uint
	Status       uint8
	Price        uint
	Date         time.Time
	Address      *DBCreateOrderAddress
	DeliveryTime uint8
}

type RespCreateOrder struct {
	Id           uint            `json:"Id"`
	Status       uint8           `json:"Status"`
	Date         time.Time       `json:"Date"`
	Address      *entity.Address `json:"Address"`
	Price        uint            `json:"Sum"`
	DeliveryTime uint8           `json:"DeliveryTime"`
}

// Для слайса заказов
type RespGetOrder struct {
	Id           uint              `json:"Id"`
	Status       uint8             `json:"Status"`
	Date         time.Time         `json:"Date"`
	Address      *RespOrderAddress `json:"Address"`
	Price        uint              `json:"Sum"`
	DeliveryTime uint8             `json:"DeliveryTime"`
	//UpdatedDate time.Time `json:"UpdatedDate"`
}

// Для конкретного заказа
type RespGetOneOrder struct {
	Id           uint              `json:"Id"`
	Status       uint8             `json:"Status"`
	Date         time.Time         `json:"Date"`
	Address      *RespOrderAddress `json:"Address"`
	OrderItems   []*OrderItems     `json:"OrderItems"`
	Price        uint              `json:"Sum"`
	DeliveryTime uint8             `json:"DeliveryTime"`
}

type ReqGetOneOrder struct {
	OrderId uint `json:"OrderId"`
	UserId  uint
}

type OrderItems struct {
	RestaurantName string                 `json:"RestaurantName"`
	Products       []*RespGetOrderProduct `json:"Products"`
}

//easyjson:json
type RespOrders []*RespGetOrder

func ToEntityCreateOrder(order *ReqCreateOrder) *entity.Order {
	return &entity.Order{
		Status:       0,
		UserID:       order.UserId,
		Date:         time.Now(),
		Address:      ToEntityCreateOrderAddress(order.Address),
		Price:        0,
		DeliveryTime: 0,
	}
}

func ToDBReqCreateOrder(order *entity.Order, products []*entity.CartProduct) *DBReqCreateOrder {
	return &DBReqCreateOrder{
		UserId:       order.UserID,
		Products:     products,
		Status:       order.Status,
		Price:        order.Price,
		Date:         order.Date,
		Address:      ToDBCreateOrderAddress(order.Address),
		DeliveryTime: order.DeliveryTime,
	}
}

func ToEntityUpdateOrder(order *ReqUpdateOrder) *entity.Order {
	return &entity.Order{
		ID:     order.Id,
		Status: order.Status,
	}
}
