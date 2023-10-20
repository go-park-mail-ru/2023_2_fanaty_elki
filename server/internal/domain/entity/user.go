package entity

import (
	"database/sql"
)

type User struct {
	ID          uint           `json:"ID"`
	Username    string         `json:"Username"`
	Password    string         `json:"Password"`
	Birthday    sql.NullString `json:"Birthday"`
	PhoneNumber string         `json:"PhoneNumber"`
	Email       string         `json:"Email"`
	Icon        sql.NullString `json:"Icon"`
}