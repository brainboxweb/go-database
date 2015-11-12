package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql" //Imported for side-effects only
)

func main() {
	/*
		user: root the user
		password: carrot
		docker-machine ip: 192.168.99.100
		port: 3306
		database name: brainbox
	 */
	db, err := sql.Open("mysql", "root:carrot@tcp(192.168.99.100:3306)/brainbox")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS hello (world varchar(50))")
	if err != nil {
		log.Fatal(err)
	}
	res, err := db.Exec(
	"INSERT INTO hello(world) VALUES('hello world!')")
	if err != nil {
		log.Fatal(err)
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("inserted %d rows", rowCount)

	rows, err := db.Query("SELECT * FROM hello")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("found row containing %q", s)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
