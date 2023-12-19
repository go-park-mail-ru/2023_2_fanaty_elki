package dto

import (
	"server/internal/domain/entity"
)

//CartProduct dto
type CartProduct struct {
	Product   *entity.Product
	ItemCount int
}

//CartWithRestaurant dto
type CartWithRestaurant struct {
	Restaurant *entity.Restaurant
	Products   []*CartProduct
	Promo      *RespPromo
}

//ReqProductID dto
type ReqProductID struct {
	ID uint
}
