package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	User     = "laon"
	Password = "nvidia"
	Database = "laon"
	Host     = "127.0.0.1"
	port     = 5432
)

//("postgres", `host= localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable`)
//	

func main() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", User, Password, Database)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	fmt.Println("postgresql connection")
}
