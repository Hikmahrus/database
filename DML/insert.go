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

	_, err = db.Exec(`INSERT INTO employee VALUES (1, 'Rizky', 25, 'Jl. Kebon Jeruk', 2000000), 
		(2, 'Adni', 27, 'Jl. Kebon Sirih', 3000000),
		(3, 'Budi', 30, 'Jl. Kebon Melati', 4000000),
		(4, 'Caca', 32, 'Jl. Kebon Anggrek', 5000000),
		(5, 'Deni', 35, 'Jl. Kebon Mawar', 6000000)`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Insert Data Success")
}
