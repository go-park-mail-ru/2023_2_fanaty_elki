package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"server/config"
	"server/db"
	adminDev "server/internal/Admin/delivery"
	adminRep "server/internal/Admin/repository/redis"
	adminUsecase "server/internal/Admin/usecase"
	cartDev "server/internal/Cart/delivery"
	cartRep "server/internal/Cart/repository/postgres"
	cartUsecase "server/internal/Cart/usecase"
	csatDev "server/internal/Csat/delivery"
	csatRep "server/internal/Csat/repository/postgres"
	csatUsecase "server/internal/Csat/usecase"
	orderDev "server/internal/Order/delivery"
	orderRep "server/internal/Order/repository/postgres"
	orderUsecase "server/internal/Order/usecase"
	productDev "server/internal/Product/delivery"
	productRep "server/internal/Product/repository/postgres"
	productUsecase "server/internal/Product/usecase"
	restaurantDev "server/internal/Restaurant/delivery"
	restaurantRep "server/internal/Restaurant/repository/postgres"
	restaurantUsecase "server/internal/Restaurant/usecase"
	sessionDev "server/internal/Session/delivery"
	sessionRep "server/internal/Session/repository/redis"
	sessionUsecase "server/internal/Session/usecase"
	userRep "server/internal/User/repository/postgres"
	userUsecase "server/internal/User/usecase"
	"server/internal/middleware"
	//"time"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/

const PORT = ":8080"

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

// var (
// 	redisAddr = flag.String("addr", "redis://redis-session:6379/0", "redis addr")

// 	host     = "test_postgres"
// 	port     = 5432
// 	user     = db.User.Username
// 	password = db.User.Password
// 	dbname   = "prinesy-poday"

// 	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// )

func main() {
	flag.Parse()
	router := mux.NewRouter()
	authRouter := mux.NewRouter()
	corsRouter := mux.NewRouter()
	adminRouter := mux.NewRouter()

	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	//time.Sleep(5 * time.Second)

	db, err := db.GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	logger := middleware.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())

	userRepo := userRep.NewUserRepo(db)
	restaurantRepo := restaurantRep.NewRestaurantRepo(db)
	productRepo := productRep.NewProductRepo(db)
	cartRepo := cartRep.NewCartRepo(db)
	sessionRepo := sessionRep.NewSessionManager(redisConn)
	orderRepo := orderRep.NewOrderRepo(db)
	csatRepo := csatRep.NewCsatRepo(db)
	adminRepo := adminRep.NewadminManager(redisConn)

	userUC := userUsecase.NewUserUsecase(userRepo, cartRepo)
	restaurantUC := restaurantUsecase.NewRestaurantUsecase(restaurantRepo, productRepo)
	cartUC := cartUsecase.NewCartUsecase(cartRepo, productRepo, sessionRepo)
	sessionUC := sessionUsecase.NewSessionUsecase(sessionRepo, userRepo)
	orderUC := orderUsecase.NewOrderUsecase(orderRepo, cartRepo, productRepo)
	productUC := productUsecase.NewProductUsecase(productRepo)
	csatUC := csatUsecase.NewCsatUsecase(csatRepo)
	adminUC := adminUsecase.NewadminUsecase(adminRepo, userRepo)

	restaurantsHandler := restaurantDev.NewRestaurantHandler(restaurantUC, logger)
	cartsHandler := cartDev.NewCartHandler(cartUC, logger)
	sessionsHandler := sessionDev.NewSessionHandler(sessionUC, userUC, logger)
	orderHandler := orderDev.NewOrderHandler(orderUC, sessionUC, logger)
	productHandler := productDev.NewProductHandler(productUC, logger)
	csatHandler := csatDev.NewCsatHandler(csatUC, logger)
	adminHandler := adminDev.NewAdminHandler(adminUC, userUC, logger)
	authMW := middleware.NewSessionMiddleware(sessionUC, logger)
	adminMW := middleware.NewAdminMiddleware(adminUC, logger)

	router.PathPrefix("/api/login").Handler(corsRouter)
	router.PathPrefix("/api/logout").Handler(authRouter)
	router.PathPrefix("/api/auth").Handler(authRouter)
	router.PathPrefix("/api/cart").Handler(authRouter)
	router.PathPrefix("/api/users/me").Handler(authRouter)
	router.PathPrefix("/api/orders").Handler(authRouter)
	router.PathPrefix("/api/csrf").Handler(authRouter)
	router.PathPrefix("/api/users").Handler(corsRouter)
	router.PathPrefix("/api/csat/admin").Handler(adminRouter)
	
	router.Use(logger.ACLogMiddleware)
	router.Use(middleware.PanicMiddleware)
	router.Use(middleware.CorsMiddleware)
	corsRouter.Use(middleware.CorsCredentionalsMiddleware)
	authRouter.Use(authMW.AuthMiddleware)
	adminRouter.Use(adminMW.AuthMiddleware)

	restaurantsHandler.RegisterHandler(router)
	productHandler.RegisterHandler(router)
	cartsHandler.RegisterHandler(authRouter)
	sessionsHandler.RegisterCorsHandler(corsRouter)
	sessionsHandler.RegisterAuthHandler(authRouter)
	orderHandler.RegisterHandler(authRouter)
	productHandler.RegisterHandler(router)
	csatHandler.RegisterHandler(router)
	csatHandler.RegisterAdminHandler(adminRouter)
	adminHandler.RegisterAdminHandler(adminRouter)
	adminHandler.RegisterCorsHandler(adminRouter)
	
	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Println("Server start at port", PORT[1:])
	err = server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
