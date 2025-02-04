package controllers

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"goweb/models"
	"goweb/utils"
)

const PW_FILE = "data/logins.json"

var UsersList models.Users

func LoadUsers() {
	bytes, err := os.ReadFile(PW_FILE)
	if err != nil {
		log.Printf("WARNING: Could not load users - %s\n", err)
		return
	}
	if err := json.Unmarshal(bytes, &UsersList); err != nil {
		log.Printf("ERROR: Could not decode users json - %s", err)
	}
}
func DumpUsers() {
	jsonData, err := json.MarshalIndent(UsersList, "", "    ")
	if err != nil {
		log.Printf("CRITICAL: Unable to dump users from memory - %s", err)
		return
	}

	if err := os.WriteFile(PW_FILE, jsonData, 0644); err != nil {
		log.Printf("CRITICAL: Unable to dump users from memory - %s", err)
		return
	}
}
func retrieveUser(u string, p string) int {
	for i, user := range UsersList.Users {
		if u == user.Username && utils.CheckPasswordHash(p, user.PasswordHash) {
			UsersList.Users[i].LastLogon = time.Now()
			if time.Now().Compare(UsersList.Users[i].AuthExpires.AddDate(0, 0, 1)) > 0 {
				UsersList.Users[i].AuthToken = utils.GenerateToken(user)
				UsersList.Users[i].AuthExpires = time.Now()
			}
			return i
		}
	}
	return -1
}
func CheckAuthToken(token string) bool {
	for _, user := range UsersList.Users {
		if token == user.AuthToken {
			return true
			//return UsersList.Users[i]
		}
	}
	return false
}
func validateLogin(u string, p string) models.User {
	index := retrieveUser(u, p)
	if index >= 0 {
		DumpUsers()
		return UsersList.Users[index]
	}
	return models.User{Id: -1}
}

func Login(u string, p string) models.User {
	return validateLogin(u, p)
}
