package main

import (
	sessDel "AuthService/internal/delivery"
	sessRep "AuthService/internal/repository/redis"
	sessUC "AuthService/internal/usecase"
	auth "AuthService/proto"
	"flag"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
	"log"
	"net"
)

//PORT port of auth service
const PORT = ":8081"

var redisAddr = flag.String("addr", "redis://redis-session:6379/0", "redis addr")

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	server := grpc.NewServer()

	sessionRepo := sessRep.NewSessionManager(redisConn)
	sessionUC := sessUC.NewSessionUsecase(sessionRepo)
	auth.RegisterSessionRPCServer(server, sessDel.NewAuthManager(sessionUC))

	log.Println("starting server at ", PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
