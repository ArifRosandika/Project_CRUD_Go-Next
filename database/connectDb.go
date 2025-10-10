package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	con := "postgres://postgres:1307@localhost:5432/CRUDGo_db?sslmode=disable"

	db, err := sql.Open("postgres", con)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connect to postgres")
	return db
}