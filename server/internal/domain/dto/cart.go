package dto

import (
	"server/internal/domain/entity"
)

type CartProduct struct {
	Product   *entity.Product
	ItemCount int
}

type CartWithRestaurant struct {
	Restaurant *entity.Restaurant
	Products   []*CartProduct
}

type ReqProductID struct {
	Id uint
}
