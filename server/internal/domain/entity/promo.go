package entity

import "time"

//Promo entity
type Promo struct {
	ID           uint
	Code         string
	PromoType    int
	Sale         uint
	RestaurantID uint
	ActiveFrom   time.Time
	ActiveTo     time.Time
}

//UserPromo entity
type UserPromo struct {
	UserID  uint
	PromoID uint
}
