package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	createDb()
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

func createDb() {

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
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

	defer db.Close()

}
