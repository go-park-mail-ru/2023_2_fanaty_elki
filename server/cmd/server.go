package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"server/internal/delivery"
	"server/internal/usecases"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/
// const allowedOrigin = "http://84.23.53.216"


const PORT = ":3333"
const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"

func main() {
	router := mux.NewRouter()
	// api := &Handler{
	//  	restaurantstore: store.NewRestaurantStore(),
	//  	userstore:       store.NewUserStore(),
	//  	sessions:        make(map[string]uint, 10),
	// }
	
	
	restaurants := usecases.NewRestaurantUsecase(nil)
	users := usecases.NewUserUsecase(nil)

	restaurants_api := delivery.NewRestaurantHandler(restaurants)
	users_api := delivery.NewUserHandler(users)
	router.HandleFunc("/restaurants", restaurants_api.GetRestaurantList).Methods(GET)
	router.HandleFunc("/users", users_api.SignUp).Methods(POST)
	router.HandleFunc("/login", users_api.Login).Methods(POST)
	router.HandleFunc("/logout", users_api.Logout).Methods(DELETE)
	router.HandleFunc("/auth", users_api.Auth).Methods(GET)
	
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Println("Server start")
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
