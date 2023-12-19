package db

import (
	"database/sql"
)

//GetPostgres gets postgres connection
func GetPostgres(psqlInfo string) (*sql.DB, error) {

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//fmt.Println("Successfully connected!")
	return db, nil
}
