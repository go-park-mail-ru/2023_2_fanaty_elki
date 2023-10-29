package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	sessionDev "server/internal/Session/delivery"
//	userDev "server/internal/User/delivery"
	restaurantDev "server/internal/Restaurant/delivery"
	sessionUC "server/internal/Session/usecase"
	userUC "server/internal/User/usecase"
	restaurantUC "server/internal/Restaurant/usecase"
	sessionRep "server/internal/Session/repository/postgres"
	userRep "server/internal/User/repository/postgres"
	restaurantRep "server/internal/Restaurant/repository/postgres"
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

	userRepo := userRep.NewUserRepo(db)
	restaurantRepo := restaurantRep.NewRestaurantRepo(db)
	sessionRepo := sessionRep.NewSessionManager(redisConn)
	
	users := userUC.NewUserUsecase(userRepo)
	restaurants := restaurantUC.NewRestaurantUsecase(restaurantRepo)
	sessions := sessionUC.NewSessionUsecase(sessionRepo, userRepo)

	//usersHandler := userDev.NewUserHandler(users)
	restaurantsHandler := restaurantDev.NewRestaurantHandler(restaurants)
	sessionsHandler := sessionDev.NewSessionHandler(sessions, users)
	

	router.HandleFunc("/api/restaurants", restaurantsHandler.GetRestaurantList).Methods(GET)
	router.HandleFunc("/api/users", sessionsHandler.SignUp).Methods(POST)
	router.HandleFunc("/api/login", sessionsHandler.Login).Methods(POST)
	router.HandleFunc("/api/logout", sessionsHandler.Logout).Methods(DELETE)
	router.HandleFunc("/api/auth", sessionsHandler.Auth).Methods(GET)
	
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
