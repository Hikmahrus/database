package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Staff = iota
	Supervisor
	Manager
)

type Employee struct {
	Id           uint
	Name         string
	Address      string
	Age          int
	Birthdate    string
	Level        uint
	IdDepartment int
}

type Department struct {
	Id             uint
	NamaDepartment string
}

func main() {
	dsn := "host=localhost user=postgres password=hikmah34 dbname=learn_gorm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Employee{})
	db.AutoMigrate(&Department{})

	db.Create(&Department{NamaDepartment: "newer department"})
	db.Create(&Department{NamaDepartment: "new department"})
	db.Create(&Department{NamaDepartment: "old department"})
	db.Create(&Department{NamaDepartment: "older department"})
	db.Create(&Department{NamaDepartment: "oldest department"})
	db.Create(&Employee{Name: "Deni", Address: "Surabaya", Age: 23, Birthdate: "2000-02-02", Level: Staff, IdDepartment: 1})
	db.Create(&Employee{Name: "Jono", Address: "Malang", Age: 24, Birthdate: "1999-02-02", Level: Supervisor, IdDepartment: 2})
	db.Create(&Employee{Name: "Andre", Address: "Sidoarjo", Age: 20, Birthdate: "2002-02-02", Level: Staff, IdDepartment: 1})
	db.Create(&Employee{Name: "Dono", Address: "Malang", Age: 33, Birthdate: "1990-02-02", Level: Supervisor, IdDepartment: 3})
	db.Create(&Employee{Name: "Joko", Address: "Malang", Age: 30, Birthdate: "1993-02-02", Level: Supervisor, IdDepartment: 3})

	var employees []Employee
	db.Model(&Employee{}).Scan(&employees)

	fmt.Println(employees)

	var deparments []Department
	db.Model(&Department{}).Scan(&deparments)

	fmt.Println(deparments)

	db.Model(&Employee{}).Where("Name = ?", "Deni").Update("Address", "Sidoarjo")
	db.Delete(&Department{}, 3)

	db.Model(&Employee{}).Scan(&employees)

	fmt.Println(employees)

	db.Model(&Department{}).Scan(&deparments)

	fmt.Println(deparments)
}
