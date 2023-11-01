package delivery

import (
//	"server/internal/domain/entity"
	userUsecase "server/internal/User/usecase"
	// "database/sql"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "time"
)


//const allowedOrigin = "http://84.23.53.216"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type UserHandler struct {
	users userUsecase.UsecaseI
	//sessManager  usecases.SessionUsecase
}

func NewUserHandler(users userUsecase.UsecaseI) *UserHandler{
	return &UserHandler{
		users: users,
		//sessManager: *sessManager,
	}
}

