// package main

// import (
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	statement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = 'yourDBName');`

// 	row := db.QueryRow(statement)
// 	var exists bool
// 	err = row.Scan(&exists)
// 	check(err)

// 	if exists == false {
// 		statement = `CREATE DATABASE yourDBName;`
// 		_, err = db.Exec(statement)
// 		check(err)
// 	}
// }

// func check(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
