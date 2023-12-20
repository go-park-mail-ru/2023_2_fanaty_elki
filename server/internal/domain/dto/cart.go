package dto

import (
	"server/internal/domain/entity"
)

//CartProduct dto
type CartProduct struct {
	Product   *entity.Product `json:"Product"`
	ItemCount int             `json:"ItemCount"`
}

//CartWithRestaurant dto
type CartWithRestaurant struct {
	Restaurant *entity.Restaurant `json:"Restaurant"`
	Products   []*CartProduct     `json:"Products"`
	Promo      *RespPromo         `json:"Promo"`
}

//ReqProductID dto
type ReqProductID struct {
	ID uint
}

type Result struct {
	Body interface{} `json:"Body"`
}

//easyjson:json
type ProductSlice []*entity.Product
