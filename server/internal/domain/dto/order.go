package dto

import (
	"server/internal/domain/entity"
	"time"
)

//ReqCreateOrder dto
type ReqCreateOrder struct {
	UserID  uint                   `json:"UserID"`
	Address *ReqCreateOrderAddress `json:"Address"`
}

//ReqUpdateOrder dto
type ReqUpdateOrder struct {
	ID     uint  `json:"Id"`
	Status uint8 `json:"Status"`
}

//DBReqCreateOrder dto
type DBReqCreateOrder struct {
	Products     []*entity.CartProduct
	UserID       uint
	Status       uint8
	Price        uint
	Date         time.Time
	Address      *DBCreateOrderAddress
	DeliveryTime uint8
}

//RespCreateOrder dto
type RespCreateOrder struct {
	ID           uint            `json:"Id"`
	Status       uint8           `json:"Status"`
	Date         time.Time       `json:"Date"`
	Address      *entity.Address `json:"Address"`
	Price        uint            `json:"Sum"`
	DeliveryTime uint8           `json:"DeliveryTime"`
}

//RespGetOrder dto Для слайса заказов
type RespGetOrder struct {
	ID           uint              `json:"Id"`
	Status       uint8             `json:"Status"`
	Date         time.Time         `json:"Date"`
	Address      *RespOrderAddress `json:"Address"`
	Price        uint              `json:"Sum"`
	DeliveryTime uint8             `json:"DeliveryTime"`
	//UpdatedDate time.Time `json:"UpdatedDate"`
}

//RespGetOneOrder dto Для конкретного заказа
type RespGetOneOrder struct {
	ID           uint              `json:"Id"`
	Status       uint8             `json:"Status"`
	Date         time.Time         `json:"Date"`
	Address      *RespOrderAddress `json:"Address"`
	OrderItems   []*OrderItems     `json:"OrderItems"`
	Price        uint              `json:"Sum"`
	DeliveryTime uint8             `json:"DeliveryTime"`
}

//ReqGetOneOrder dto
type ReqGetOneOrder struct {
	OrderID uint `json:"OrderId"`
	UserID  uint
}

//OrderItems dto
type OrderItems struct {
	RestaurantName string                 `json:"RestaurantName"`
	Products       []*RespGetOrderProduct `json:"Products"`
}

//easyjson:json
type RespOrders []*RespGetOrder

//ToEntityCreateOrder transforms ReqCreateOrder to Order
func ToEntityCreateOrder(order *ReqCreateOrder) *entity.Order {
	return &entity.Order{
		Status:       0,
		UserID:       order.UserID,
		Date:         time.Now(),
		Address:      ToEntityCreateOrderAddress(order.Address),
		Price:        0,
		DeliveryTime: 0,
	}
}

//ToDBReqCreateOrder transforms order and cart products to DBReqCreateOrder
func ToDBReqCreateOrder(order *entity.Order, products []*entity.CartProduct) *DBReqCreateOrder {
	return &DBReqCreateOrder{
		UserID:       order.UserID,
		Products:     products,
		Status:       order.Status,
		Price:        order.Price,
		Date:         order.Date,
		Address:      ToDBCreateOrderAddress(order.Address),
		DeliveryTime: order.DeliveryTime,
	}
}

//ToEntityUpdateOrder transforms ReqUpdateOrder to Order
func ToEntityUpdateOrder(order *ReqUpdateOrder) *entity.Order {
	return &entity.Order{
		ID:     order.ID,
		Status: order.Status,
	}
}
