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
	Promo      *RespPromo
}

type ReqProductID struct {
	Id uint
}
