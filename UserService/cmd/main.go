package main

import (
	userDel "UserService/internal/User/delivery"
	userRep "UserService/internal/User/repository/postgres"
	userUC "UserService/internal/User/usecase"
	proto "UserService/proto"
	"database/sql"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"fmt"
	"log"
	"net"
	"time"
)

//PORT of microservice
const PORT = ":8083"

var (
	host     = "test_postgres"
	port     = 5432
	user     = "uliana"
	password = "uliana"
	dbname   = "prinesy-poday"

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

//GetPostgres gets connections of postgresql
func GetPostgres(psqlInfo string) (*sql.DB, error) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	db, err := GetPostgres(psqlInfo)
	if err != nil {
		fmt.Println(err, " ", psqlInfo)
		log.Fatalf("cant connect to postgres")
		return
	}
	defer db.Close()

	server := grpc.NewServer()

	userRepo := userRep.NewUserRepo(db)
	userUC := userUC.NewUserUsecase(userRepo)
	proto.RegisterUserRPCServer(server, userDel.NewUserManager(userUC))

	log.Println("starting server at ", PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
