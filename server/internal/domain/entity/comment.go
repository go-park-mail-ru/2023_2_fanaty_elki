package entity

import (
	"time"
)

type Comment struct {
	Id			 uint
	Text 		 string
	RestaurantId uint
	UserId		 uint
	Rating 		 uint8
	Date		 time.Time
}