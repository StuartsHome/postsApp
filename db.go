package main

import (

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MYSQL Tutorial")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)first_instance")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully connected to MySQL database")
}wo