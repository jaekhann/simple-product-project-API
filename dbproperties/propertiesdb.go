package dbproperties

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "jaesql"
	dbname   = "dbwarehouse"
)

func ConnectToDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
