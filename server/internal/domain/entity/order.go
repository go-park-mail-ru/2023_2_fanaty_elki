package entity

import (
	
	"time"
)

type Order struct {
	ID 			 uint
	Status 		 uint8
	UserId 		 uint
	Date 		 time.Time
	Products 	 []*CartProduct
	Address 	 *Address
	Price 		 uint
	DeliveryTime uint8
}