package entity

type Restaurant struct {
	ID              uint
	Name            string
	Rating          float32
	CommentsCount   int
	Icon            string
	Category        string
	MinDeliveryTime int
	MaxDeliveryTime int
	DeliveryPrice   float32
}
