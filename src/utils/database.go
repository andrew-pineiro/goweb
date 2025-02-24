package utils

import (
	"database/sql"
	"fmt"
	"goweb/models"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE = "./data/goweb.db"

func CheckUserByName(username string) models.User {
	var user models.User
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Printf("ERROR: Could not open database %s", err)
		return user
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM Users WHERE username = '%s'", username))
	if err != nil {
		log.Printf("ERROR: Could not query db  %s", err)
		return user
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Guid, &user.Username, &user.Email, &user.Password, &user.LastLoginDate, &user.LastChangeDate)
		if err != nil {
			log.Printf("ERROR: could not scan row %s", err)
			continue
		}
	}
	return user
}
func UpdateLastLogin(guid string) {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Printf("ERROR: Could not open database %s", err)
		return
	}
	defer db.Close()

	query, _ := db.Prepare("UPDATE Users set last_login_date=? where guid=?")

	_, err = query.Exec(time.Now().String(), guid)
	if err != nil {
		log.Printf("ERROR: unable to update record %s", err)
	}
}
