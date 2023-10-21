package db

import (
	"database/sql"
	"fmt"
)

func GetPostgres() (*sql.DB, error) {
	var (
		host     = "localhost"
		port     = 5432
		user     = 	User.username
		password = 	User.password
		dbname   = "prinesy-poday"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
