package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db_user := "lt"
	db_name := "lt"
	db_server := "localhost"
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgresql://%s@%s:26257/%s?sslmode=disable", db_user, db_server, db_name))
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}
	defer db.Close()

	//for i := 0; i < 200; i++ {
		// Insert two rows into the "accounts" table.
		//fmt.Printf(time.Now().Format("2006-01-02T15:04:05.000000-07:00"))
		if _, err := db.Exec(
			//"INSERT INTO lt.lasts (last, tag) VALUES ($1, 'test')", time.Now().Format("2006-01-02T15:04:05.000000-07:00")); err != nil {
			"INSERT INTO lt.lasts (last, tag) VALUES (NOW(), 'test')"); err != nil {
			log.Fatal(err)
		}
	//}

	// Print out the balances.
	rows, err := db.Query("SELECT * FROM lt.lasts ORDER BY last")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//fmt.Println("Initial balances:")
	for rows.Next() {
		var id, last, tag string
		if err := rows.Scan(&id, &last, &tag); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %s %s\n", id, last, tag)
	}

}
