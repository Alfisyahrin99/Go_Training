package main

import (
	"database/sql"
	"fmt"

	_ "godb/connection"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id      int
	Name    string
	Balance int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = 123
	dbName   = "lapakin_db"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	newCst := Customer{
		Name:    "Juan",
		Balance: 15000,
	}

	insertCustomer(db, newCst)
}

func insertCustomer(db *sql.DB, newCustomer Customer) {
	query := fmt.Sprintf("insert into mst_customer(name,balance)values($1,$2)")
	result, errExec := db.Exec(query, newCustomer.Name, newCustomer.Balance)

	if errExec != nil {
		fmt.Println(errExec.Error())
	}

	fmt.Println(result)
}
