package dto

import "server/internal/domain/entity"

type CartProduct struct {
	Product   *entity.Product
	ItemCount int
}

type ReqProductID struct {
	Id uint
}
