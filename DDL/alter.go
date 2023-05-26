package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hikmah34"
	dbname   = "learn_connection"
)

func ConnectDB() (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`ALTER TABLE employee RENAME TO employees`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table employee renamed to employees")
}
