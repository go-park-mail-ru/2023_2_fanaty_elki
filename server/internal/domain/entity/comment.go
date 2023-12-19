package entity

import (
	"time"
)

//Comment entity
type Comment struct {
	ID           uint
	Text         string
	RestaurantID uint
	UserID       uint
	Rating       uint8
	Date         time.Time
}
