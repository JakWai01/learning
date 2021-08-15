package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	mux sync.Mutex
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

	exit := false

	for exit == false {
		fmt.Println("-----------------")
		fmt.Println("TODOGO")
		fmt.Println("1. Display togos") // Done
		fmt.Println("2. Add togo")      // Done
		fmt.Println("3. Finish togo")   // Done
		fmt.Println("4. Start pomo")    // Done
		fmt.Println("5. Exit")          // Done
		fmt.Println("-----------------")

		reader := bufio.NewReader(os.Stdin)
		operation, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		switch operation {
		case "1\n":
			fallthrough
		case "d\n":
			browseEntries(database)
		case "2\n":
			fallthrough
		case "a\n":
			fmt.Println("Enter content of new togo: ")
			todoContent, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			addEntry(statement, *database, todoContent, err)
		case "3\n":
			fallthrough
		case "f\n":
			browseEntries(database)

			fmt.Println("Enter id of togo you want to finish: ")

			idString, err := reader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}
			idString = idString[0:(len(idString) - 1)]
			id, err := strconv.Atoi(idString)
			if err != nil {
				log.Fatal(err)
			}

			removeEntries(statement, *database, id, err)
		case "4\n":
			fallthrough
		case "s\n":
			startPomo()
		case "5\n":
			fallthrough
		case "e\n":
			exit = true

		default:
			fmt.Println("Please enter valid input")
		}
	}
}

func addEntry(statement *sql.Stmt, database sql.DB, content string, err error) {
	mux.Lock()
	statement, err = database.Prepare("INSERT INTO todos (todo) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(content)
	mux.Unlock()
}

func browseEntries(database *sql.DB) {
	mux.Lock()
	if fileExists("todos.db") {
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
	} else {
		fmt.Println("The database was not created yet! You have to add a togo first")
	}
	mux.Unlock()
}

func removeEntries(statement *sql.Stmt, database sql.DB, id int, err error) {
	mux.Lock()
	statement, err = database.Prepare("DELETE FROM todos WHERE id=(?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(id)
	mux.Unlock()
}

func startPomo() {
	fmt.Println("25 minute timer startet..")
	time.Sleep(time.Minute * 25)
	fmt.Println("Time's up! Take a break and remanage your done togos")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
