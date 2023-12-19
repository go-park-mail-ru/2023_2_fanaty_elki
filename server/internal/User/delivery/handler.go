package delivery

import (
	userUsecase "server/internal/User/usecase"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//UserHandler handles requests of user's api
type UserHandler struct {
	users userUsecase.Iusecase
}

//NewUserHandler creates User handler
func NewUserHandler(users userUsecase.Iusecase) *UserHandler {
	return &UserHandler{
		users: users,
	}
}

