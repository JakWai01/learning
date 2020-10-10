package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Initialize database
	database, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, todo TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	addEntry(statement, *database, "test", err)
	// statement, err = database.Prepare("INSERT INTO todos (todo) VALUES (?)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// statement.Exec("digital decluttering")

	// Prints entries - Browse component
	browseEntries(database)
	// rows, err := database.Query("SELECT id, todo FROM todos")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var id int
	// var todo string

	// for rows.Next() {
	// 	rows.Scan(&id, &todo)
	// 	fmt.Println(strconv.Itoa(id) + ": " + todo)
	// }

	fmt.Println("OK")

	// Flags

	// Todo list with simple structure
	exit := false
	// Menu

	for exit == false {
		fmt.Println("TODOGO")
		fmt.Println("1. Display togos")
		fmt.Println("2. Add togo")
		fmt.Println("3. Finish togo")
		fmt.Println("4. Start pomo")
		fmt.Println("5. Exit")

		reader := bufio.NewReader(os.Stdin)
		operation, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		switch operation {
		case "1\n":
			fallthrough
		case "d\n":

		case "2\n":
			fallthrough
		case "a\n":

		case "3\n":
			fallthrough
		case "f\n":

		case "4\n":
			fallthrough
		case "s\n":

		case "5\n":
			fallthrough
		case "e\n":
			exit = true

		default:
			fmt.Println("Please enter valid input")
		}
	}
	// write into database

	// read out of database
}

//func createDb() {

// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:8000)/test")
// if err != nil {
// 	log.Fatal(err)
// } else {
// 	fmt.Println("OK")
// }
// defer db.Close()

// _, err = db.Exec("CREATE DATABASE todos")
// if err != nil {
// 	fmt.Println(err.Error())
// } else {
// 	fmt.Println("Successfully created database..")
// }

// _, err = db.Exec("USE todos")
// if err != nil {
// 	log.Fatal(err)
// } else {
// 	fmt.Println("DB selected successfully")
// }

// stmt, err := db.Prepare("CREATE Table togos(id int NOT NULL AUTO_INCREMENT, todo varchar(50), PRIMARY KEY (id));")
// if err != nil {
// 	log.Fatal(err)
// }

// _, err = stmt.Exec()
// if err != nil {
// 	fmt.Println(err.Error())
// } else {
// 	fmt.Println("Table created successfully..")
// }

//defer db.Close()

//}

func addEntry(statement *sql.Stmt, database sql.DB, content string, err error) {
	statement, err = database.Prepare("INSERT INTO todos (todo) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(content)
}

func browseEntries(database *sql.DB) {
	rows, err := database.Query("SELECT id, todo FROM todos")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var todo string

	for rows.Next() {
		rows.Scan(&id, &todo)
		fmt.Println(strconv.Itoa(id) + ": " + todo)
	}
}

func removeEntries() {

}
