package db

import (
	//"database/sql"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dsn = "user=uliana dbname=prinesy-poday password=uliana host=localhost port=5432 sslmode=disable pool_max_conns=20"
)

func GetPostgres() (*pgxpool.Pool, error) {
	// var (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "uliana"
	// 	password = "uliana"
	// 	dbname   = "prinesy-poday"
	// )

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	return nil, err
	// }

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// err = conn.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err == nil {
		fmt.Println("Ping sucessful")
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
