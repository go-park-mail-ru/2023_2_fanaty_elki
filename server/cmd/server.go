package main

import (
	"errors"
	"fmt"
	"net/http"

	"flag"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"server/db"
	cartDev "server/internal/Cart/delivery"
	cartRep "server/internal/Cart/repository/postgres"
	cartUsecase "server/internal/Cart/usecase"
	orderDev "server/internal/Order/delivery"
	orderRep "server/internal/Order/repository/postgres"
	orderUsecase "server/internal/Order/usecase"
	productRep "server/internal/Product/repository/postgres"
	restaurantDev "server/internal/Restaurant/delivery"
	restaurantRep "server/internal/Restaurant/repository/postgres"
	restaurantUsecase "server/internal/Restaurant/usecase"
	sessionDev "server/internal/Session/delivery"
	sessionRep "server/internal/Session/repository/redis"
	sessionUsecase "server/internal/Session/usecase"
	userRep "server/internal/User/repository/postgres"
	userUsecase "server/internal/User/usecase"
	"server/internal/middleware"
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
	
	corsRouter := mux.NewRouter()
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
	cartRepo := cartRep.NewCartRepo(db)
	sessionRepo := sessionRep.NewSessionManager(redisConn)
	orderRepo := orderRep.NewOrderRepo(db)

	userUC := userUsecase.NewUserUsecase(userRepo, cartRepo)
	restaurantUC := restaurantUsecase.NewRestaurantUsecase(restaurantRepo, productRepo)
	cartUC := cartUsecase.NewCartUsecase(cartRepo, productRepo, sessionRepo)
	sessionUC := sessionUsecase.NewSessionUsecase(sessionRepo, userRepo)
	orderUC := orderUsecase.NewOrderUsecase(orderRepo)

	restaurantsHandler := restaurantDev.NewRestaurantHandler(restaurantUC)
	cartsHandler := cartDev.NewCartHandler(cartUC)
	sessionsHandler := sessionDev.NewSessionHandler(sessionUC, userUC)
	orderHandler := orderDev.NewOrderHandler(orderUC, sessionUC)
	
	router.PathPrefix("/api/login").Handler(corsRouter)
	router.PathPrefix("/api/logout").Handler(corsRouter)
	router.PathPrefix("/api/auth").Handler(corsRouter)
	router.PathPrefix("/api/cart").Handler(corsRouter)
	router.PathPrefix("/api/me").Handler(corsRouter)
	router.PathPrefix("/api/orders").Handler(corsRouter)
	
	router.Use(middleware.CorsMiddleware)
	corsRouter.Use(middleware.CorsCredentionalsMiddleware)
	router.HandleFunc("/api/restaurants", restaurantsHandler.GetRestaurantList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{id}", restaurantsHandler.GetRestaurantById).Methods(http.MethodGet)
	corsRouter.HandleFunc("/api/cart", cartsHandler.GetCart).Methods(http.MethodGet)
	corsRouter.HandleFunc("/api/cart/add", cartsHandler.AddProductToCart).Methods(http.MethodPost)
	corsRouter.HandleFunc("/api/cart/delete", cartsHandler.DeleteProductFromCart).Methods(http.MethodPost)
	corsRouter.HandleFunc("/api/cart/update/up", cartsHandler.UpdateItemCountUp).Methods(http.MethodPatch)
	corsRouter.HandleFunc("/api/cart/update/down", cartsHandler.UpdateItemCountDown).Methods(http.MethodPatch)
	router.HandleFunc("/api/users", sessionsHandler.SignUp).Methods(http.MethodPost)
	corsRouter.HandleFunc("/api/login", sessionsHandler.Login).Methods(http.MethodPost)
	corsRouter.HandleFunc("/api/logout", sessionsHandler.Logout).Methods(http.MethodDelete)
	corsRouter.HandleFunc("/api/auth", sessionsHandler.Auth).Methods(http.MethodGet)
	corsRouter.HandleFunc("/api/me", sessionsHandler.Profile).Methods(http.MethodGet)
	corsRouter.HandleFunc("/api/me", sessionsHandler.UpdateProfile).Methods(http.MethodPatch)
	corsRouter.HandleFunc("/api/orders", orderHandler.CreateOrder).Methods(http.MethodPost)
	corsRouter.HandleFunc("/api/orders", orderHandler.UpdateOrder).Methods(http.MethodPatch)
	corsRouter.HandleFunc("/api/orders", orderHandler.GetOrders).Methods(http.MethodGet)
	corsRouter.HandleFunc("/api/orders/{id}", orderHandler.GetOrder).Methods(http.MethodGet)

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
