package entity

import "time"

type Promo struct {
	ID           uint
	Code         string
	PromoType    int
	Sale         uint
	RestaurantId uint
	ActiveFrom   time.Time
	ActiveTo     time.Time
}

type UserPromo struct {
	UserId  uint
	PromoId uint
}
