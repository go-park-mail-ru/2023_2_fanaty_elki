package entity

import (
	"time"
)

//Order entity
type Order struct {
	ID           uint
	Status       uint8
	UserID       uint
	Date         time.Time
	Products     []*CartProduct
	Address      *Address
	Price        uint
	DeliveryTime uint8
}
