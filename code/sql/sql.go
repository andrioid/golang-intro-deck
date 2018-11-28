// https://github.com/golang/go/wiki/SQLInterface

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sql.db")
	if err != nil {
		log.Fatal(err) // bad practice, dont do this
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS foo (id integer not null primary key, name text);
	DELETE FROM foo;
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO foo (id, name) VALUES (42, 'Douglas Adams')`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name FROM foo")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

}
