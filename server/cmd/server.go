package main

import (
	"errors"
	"fmt"
	"net/http"
	sessionDev "server/internal/Session/delivery"

	"github.com/gorilla/mux"

	//	userDev "server/internal/User/delivery"
	//	userDev "server/internal/User/delivery"
	"flag"
	"log"
	"server/db"
	productRep "server/internal/Product/repository/postgres"
	restaurantDev "server/internal/Restaurant/delivery"
	restaurantRep "server/internal/Restaurant/repository/postgres"
	restaurantUsecase "server/internal/Restaurant/usecase"
	sessionRep "server/internal/Session/repository/redis"
	sessionUsecase "server/internal/Session/usecase"
	userRep "server/internal/User/repository/postgres"
	userUsecase "server/internal/User/usecase"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/

const PORT = ":3333"

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")

	host     = "localhost"
	port     = 5432
	user     = db.User.Username
	password = db.User.Password
	dbname   = "prinesy-poday"
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

func main() {
	flag.Parse()
	router := mux.NewRouter()

	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatalf("cant connect to redis")
	}

	db, err := db.GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	userRepo := userRep.NewUserRepo(db)
	restaurantRepo := restaurantRep.NewRestaurantRepo(db)
	productRepo := productRep.NewProductRepo(db)
	sessionRepo := sessionRep.NewSessionManager(redisConn)

	userUC := userUsecase.NewUserUsecase(userRepo)
	restaurantUC := restaurantUsecase.NewRestaurantUsecase(restaurantRepo, productRepo)
	sessionUC := sessionUsecase.NewSessionUsecase(sessionRepo, userRepo)

	restaurantsHandler := restaurantDev.NewRestaurantHandler(restaurantUC)
	sessionsHandler := sessionDev.NewSessionHandler(sessionUC, userUC)

	router.HandleFunc("/api/restaurants", restaurantsHandler.GetRestaurantList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{id}", restaurantsHandler.GetRestaurantById).Methods(http.MethodGet)
	router.HandleFunc("/api/users", sessionsHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/api/login", sessionsHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/logout", sessionsHandler.Logout).Methods(http.MethodDelete)
	router.HandleFunc("/api/auth", sessionsHandler.Auth).Methods(http.MethodDelete)
	router.HandleFunc("/api/me", sessionsHandler.Profile).Methods(http.MethodGet)
	router.HandleFunc("/api/me", sessionsHandler.UpdateProfile).Methods(http.MethodPatch)

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
