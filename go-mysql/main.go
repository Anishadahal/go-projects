package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Drivers: ", sql.Drivers())

	//open connection
	db, err := sql.Open("mysql", "root@tcp(localhost)/testdb")
	if err != nil {
		log.Fatal("Unable to open connection to DB")
	}
	defer db.Close()

	results, err := db.Query("select * from table_test;")
	if err != nil {
		log.Fatal("Error when fetching table_test table rows", err)
	}
	defer results.Close()

	for results.Next() {
		var (
			id    int
			name  string
			price int
		)
		err = results.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal("Unable to parse row")
		}
		fmt.Printf("ID: %d, NAME: %s, PRICE: %d\n", id, name, price)
	}
	var (
		id    int
		name  string
		price int
	)
	err = db.QueryRow("Select * from table_test where id=1").Scan(&id, &name, &price)

	if err != nil {
		log.Fatal("Unable to parse row: ", err)
	}
	fmt.Printf("ID: %d, NAME: %s, PRICE: %d\n", id, name, price)

	products := []struct {
		name  string
		price int
	}{
		{"light", 10},
		{"mic", 100},
		{"table", 30},
	}

	stmt, err := db.Prepare("INSERT INTO table_test(name, price) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Unable to prepare statements: ", err)
	}
	defer stmt.Close()

	for _, product := range products {
		_, err := stmt.Exec(product.name, product.price)
		if err != nil {
			log.Fatal("Unable to execute statement: ", err)
		}

	}
}
