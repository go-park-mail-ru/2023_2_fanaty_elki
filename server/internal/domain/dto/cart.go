package dto

import (
	"server/internal/domain/entity"
)

type CartProduct struct {
	Product   *entity.Product `json:"Product"`
	ItemCount int             `json:"ItemCount"`
}

type CartWithRestaurant struct {
	Restaurant *entity.Restaurant `json:"Restaurant"`
	Products   []*CartProduct     `json:"Products"`
	Promo      *RespPromo         `json:"Promo"`
}

type ReqProductID struct {
	Id uint
}

type Result struct {
	Body interface{} `json:"Body"`
}

//easyjson:json
type ProductSlice []*entity.Product
