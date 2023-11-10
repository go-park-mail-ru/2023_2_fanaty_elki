package entity

import (
	"time"
)

type Order struct {
	ID uint
	Status string
	UserId uint
	Date time.Time
	Address *Address
}