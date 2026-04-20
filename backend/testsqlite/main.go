package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT sqlite_version()").Scan(&version)
	fmt.Println("SQLite version:", version)
}
