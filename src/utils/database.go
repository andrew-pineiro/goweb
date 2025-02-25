package utils

import (
	"database/sql"
	"fmt"
	"goweb/models"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const CURR_MIGR_VER = 1

var DB_FILE string

func InitializeDb(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	DB_FILE = fmt.Sprintf("%s/%s", path, "goweb.db")
	if _, err := os.Stat(DB_FILE); os.IsNotExist(err) {
		os.Create(DB_FILE)
	}
}
func CheckMigrationVersion() int {
	version := 0
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Println(err)
		return -1
	}
	defer db.Close()
	err = db.QueryRow("PRAGMA user_version").Scan(&version)

	if err == sql.ErrNoRows {
		return version
	} else if err != nil {
		log.Println(err)
		return -1
	}
	return version
}
func MigrationHandler(version int) {
	db, err := sql.Open("sqlite3", DB_FILE)
	if err != nil {
		log.Printf("ERROR: could not configure migration %s", err)
		return
	}
	defer db.Close()
	switch version + 1 {
	case 1:
		query, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Users (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`guid` VARCHAR(255) NOT NULL, `username` VARCHAR(255) NOT NULL,`email` VARCHAR(255) NOT NULL DEFAULT '',`password` VARCHAR(255) NOT NULL DEFAULT '',`last_login_date` DATETIME NOT NULL DEFAULT '',`last_change_date` DATETIME NOT NULL DEFAULT '');")
		_, err := query.Exec()
		if err != nil {
			log.Printf("ERROR: could not complete migration version %d %s", version, err)
			break
		}
		newUUID := uuid.New()
		query, _ = db.Prepare("INSERT INTO Users (guid, username, password, email) VALUES (?, ?, ?, ?)")

		//TODO: unhardcode the admin account
		_, err = query.Exec(newUUID, "admin", HashPassword("goweb25"), "admin@admin.com")
		if err != nil {
			log.Printf("ERROR: could not create default admin user")
			break
		}

		query, _ = db.Prepare("PRAGMA user_version = 1;")
		_, err = query.Exec()
		if err != nil {
			log.Printf("ERROR: could not update migration version %s", err)
			break
		}
		log.Printf("MIGRATION: Completed version 1 migration")
	}
}

// TODO: make these more generic and then create a repo file
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
