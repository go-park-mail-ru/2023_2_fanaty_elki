package entity

type User struct {
	ID       uint   `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Birthday string `json:"Birthday"`
	Email    string `json:"Email"`
}