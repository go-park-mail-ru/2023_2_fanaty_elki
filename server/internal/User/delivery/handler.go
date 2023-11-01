package delivery

import (
	userUsecase "server/internal/User/usecase"
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type UserHandler struct {
	users userUsecase.UsecaseI
}

func NewUserHandler(users userUsecase.UsecaseI) *UserHandler {
	return &UserHandler{
		users: users,
	}
}
