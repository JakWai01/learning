package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database created successfully")
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS person (id INTEGER PRIMARY KEY, firstName text, name text)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO person (firstName, name) VALUES ('Jakob', 'Waibel');")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("DELETE FROM person WHERE id = 22")
	statement.Exec()
	// statement, err = database.Prepare("DELETE FROM person WHERE id = 7")
	// statement.Exec()
	// statement, err = database.Prepare("DELETE FROM person WHERE id = 8")
	// statement.Exec()
	// statement, err = database.Prepare("DELETE FROM person WHERE id = 9")
	// statement.Exec()
	// statement, err = database.Prepare("DELETE FROM person WHERE id = 10")
	// statement.Exec()

	rows, err := database.Query("SELECT id, firstName, name FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var firstName string
	var name string

	for rows.Next() {
		rows.Scan(&id, &firstName, &name)
		fmt.Println(strconv.Itoa(id) + ": " + firstName + " " + name)
	}

}
