package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/db"
	addressDev "server/internal/Address/delivery"
	addressRep "server/internal/Address/repository/postgres"
	addressUsecase "server/internal/Address/usecase"
	cartDev "server/internal/Cart/delivery"
	cartRep "server/internal/Cart/repository/postgres"
	cartUsecase "server/internal/Cart/usecase"
	commentDev "server/internal/Comment/delivery"
	commentRep "server/internal/Comment/repository/postgres"
	commentUsecase "server/internal/Comment/usecase"
	orderDev "server/internal/Order/delivery"
	orderRep "server/internal/Order/repository/postgres"
	orderUsecase "server/internal/Order/usecase"
	productDev "server/internal/Product/delivery"
	productRep "server/internal/Product/repository/microservice"
	productUsecase "server/internal/Product/usecase"
	promoDev "server/internal/Promo/delivery"
	promoRep "server/internal/Promo/repository/postgres"
	promoUsecase "server/internal/Promo/usecase"
	restaurantDev "server/internal/Restaurant/delivery"
	restaurantRep "server/internal/Restaurant/repository/postgres"
	restaurantUsecase "server/internal/Restaurant/usecase"
	sessionDev "server/internal/Session/delivery"
	sessionRep "server/internal/Session/repository/microservice"
	sessionUsecase "server/internal/Session/usecase"
	userRep "server/internal/User/repository/microservice"
	userUsecase "server/internal/User/usecase"
	"server/internal/domain/entity"
	"server/internal/middleware"
	auth "server/proto/auth"
	product "server/proto/product"
	userP "server/proto/user"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/

//PORT of main app
const PORT = ":8080"

// var (
// 	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")

// 	host     = "localhost"
// 	port     = 5432
// 	user     = db.User.Username
// 	password = db.User.Password
// 	dbname   = "prinesy-poday"
// 	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// )

var (
	redisAddr = flag.String("addr", "redis://redis-session:6379/0", "redis addr")

	host     = "test_postgres"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "prinesy-poday"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

func main() {
	flag.Parse()
	router := mux.NewRouter()
	authRouter := mux.NewRouter()
	corsRouter := mux.NewRouter()

	grpcConnAuth, err := grpc.Dial(
		"auth_mvs:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnAuth.Close()
	authManager := auth.NewSessionRPCClient(grpcConnAuth)

	grpcConnProduct, err := grpc.Dial(
		"product_mvs:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnProduct.Close()
	productManager := product.NewProductRPCClient(grpcConnProduct)

	grpcConnUser, err := grpc.Dial(
		"user_mvs:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcConnUser.Close()
	userManager := userP.NewUserRPCClient(grpcConnUser)

	time.Sleep(5 * time.Second)

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

	var OKHitCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ok_request_count",
			Help: "200 status counter",
		},
	)

	var InternalServerErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "internal_server_error_request_count",
			Help: "500 status counter",
		},
	)

	var NotFoundErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "no_found_server_error_request_count",
			Help: "400 status counter",
		},
	)

	var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	var timerhits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timerhits",
	}, []string{"status", "path"})

	hitstats := &entity.HitStats{
		Ok:                  OKHitCounter,
		InternalServerError: InternalServerErrorCounter,
		NotFoundError:       NotFoundErrorCounter,
		URLMetric:           *hits,
		Timing:              *timerhits,
	}

	prometheus.MustRegister(OKHitCounter)
	prometheus.MustRegister(InternalServerErrorCounter)
	prometheus.MustRegister(NotFoundErrorCounter)
	prometheus.MustRegister(hits)
	prometheus.MustRegister(timerhits)

	logger := middleware.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)

	userRepo := userRep.NewUserMicroService(userManager)
	restaurantRepo := restaurantRep.NewRestaurantRepo(db)
	productRepo := productRep.NewProductMicroService(productManager)
	cartRepo := cartRep.NewCartRepo(db)
	sessionRepo := sessionRep.NewMicroService(authManager)
	orderRepo := orderRep.NewOrderRepo(db)
	commentRepo := commentRep.NewCommentRepo(db)
	promoRepo := promoRep.NewPromoRepo(db)
	addressRepo := addressRep.NewAddressRepo(db)

	userUC := userUsecase.NewUserUsecase(userRepo, cartRepo)
	sessionUC := sessionUsecase.NewSessionUsecase(sessionRepo, userRepo, addressRepo)
	restaurantUC := restaurantUsecase.NewRestaurantUsecase(restaurantRepo, productRepo, sessionRepo, orderRepo)
	cartUC := cartUsecase.NewCartUsecase(cartRepo, productRepo, sessionRepo, restaurantRepo, promoRepo)
	orderUC := orderUsecase.NewOrderUsecase(orderRepo, cartRepo, productRepo)
	productUC := productUsecase.NewProductUsecase(productRepo)
	commentUC := commentUsecase.NewCommentUsecase(commentRepo, userRepo, restaurantRepo)
	promoUC := promoUsecase.NewPromoUsecase(cartRepo, promoRepo, sessionRepo, restaurantRepo)
	addressUC := addressUsecase.NewAddressUsecase(addressRepo, sessionRepo)

	restaurantsHandler := restaurantDev.NewRestaurantHandler(restaurantUC, logger)
	cartsHandler := cartDev.NewCartHandler(cartUC, logger)
	sessionsHandler := sessionDev.NewSessionHandler(sessionUC, userUC, logger)
	orderHandler := orderDev.NewOrderHandler(orderUC, sessionUC, logger)
	productHandler := productDev.NewProductHandler(productUC, logger)
	commentHandler := commentDev.NewCommentHandler(commentUC, sessionUC, logger)
	promoHandler := promoDev.NewPromoHandler(promoUC, logger)
	addressHandler := addressDev.NewAddressHandler(addressUC, sessionUC, logger)

	authMW := middleware.NewSessionMiddleware(sessionUC, logger)

	router.PathPrefix("/api/login").Handler(corsRouter)
	router.PathPrefix("/api/logout").Handler(authRouter)
	router.PathPrefix("/api/auth").Handler(authRouter)
	router.PathPrefix("/api/cart").Handler(authRouter)
	router.PathPrefix("/api/users/me").Handler(authRouter)
	router.PathPrefix("/api/orders").Handler(authRouter)
	router.PathPrefix("/api/csrf").Handler(authRouter)
	router.PathPrefix("/api/users").Handler(corsRouter)
	router.PathPrefix("/api/comments").Handler(authRouter).Methods(http.MethodPost, http.MethodOptions)
	router.PathPrefix("/api/promo").Handler(authRouter)

	router.Use(middleware.PanicMiddleware)
	router.Use(logger.ACLogMiddleware)
	router.Use(middleware.CorsMiddleware)
	corsRouter.Use(middleware.CorsCredentionalsMiddleware)
	authRouter.Use(authMW.AuthMiddleware)

	restaurantsHandler.RegisterHandler(router)
	productHandler.RegisterHandler(router)
	cartsHandler.RegisterHandler(authRouter)
	sessionsHandler.RegisterCorsHandler(corsRouter)
	sessionsHandler.RegisterAuthHandler(authRouter)
	orderHandler.RegisterHandler(authRouter)
	productHandler.RegisterHandler(router)
	commentHandler.RegisterPostHandler(authRouter)
	commentHandler.RegisterGetHandler(router)
	promoHandler.RegisterHandler(authRouter)
	addressHandler.RegisterHandler(authRouter)

	router.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		promhttp.Handler().ServeHTTP(w, r)
	}).Methods(http.MethodGet)

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
