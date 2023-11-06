package entity

import "database/sql"

type User struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}
