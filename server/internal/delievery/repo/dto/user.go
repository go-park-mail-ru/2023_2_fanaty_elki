package dto

type CreateUserDTO struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Birthday string `json:"Birthday"`
	Email    string `json:"Email"`
}
