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

type Employee struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Age     int    `db:"age"`
	Salary  int    `db:"salary"`
}

func ConnectDB() (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func main() {
	var listEmployee []Employee

	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	rows, err := db.Exec(`SELECT * FROM employees`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var employee Employee

		// kita tampung setiap baris data ke struct Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Address, &employee.Age, &employee.Salary)
		if err != nil {
			panic(err)
		}

		// kemudian kita tambahkan struct Employee ke listEmployee
		listEmployee = append(listEmployee, employee)
	}

	fmt.Println("Update Data Success")
}
