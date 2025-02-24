package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetDBRows() {
	db, err := sql.Open("sqlite3", "./data/goweb.db")
	if err != nil {
		log.Fatalf("ERROR: Could not open database %s", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Printf("ERROR: Could not query db  %s", err)
		return
	}
	for rows.Next() {
		var id int
		var guid string
		var username string
		var email string

		err = rows.Scan(&id, &guid, &username, &email)
		if err != nil {
			continue
		}
		fmt.Printf("Found user: %s", username)
	}

}
