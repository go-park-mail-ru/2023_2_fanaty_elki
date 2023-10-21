package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"server/internal/delivery"
	"server/internal/usecases"
	"server/db"
	"flag"
	_"github.com/lib/pq"
	"github.com/gomodule/redigo/redis"
	"log"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/

const PORT = ":3333"
const GET = "GET"
const POST = "POST"
const DELETE = "DELETE"

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")
)


func main() {
	flag.Parse()
	router := mux.NewRouter()
	
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatalf("cant connect to redis")
	}

	db, err := db.GetPostgres()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	
	restaurants := usecases.NewRestaurantUsecase(db)
	users := usecases.NewUserUsecase(db)
	sessions := usecases.NewSessionUsecase(redisConn)

	restaurants_api := delivery.NewRestaurantHandler(restaurants)
	users_api := delivery.NewUserHandler(users, sessions)

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
	err = server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
