package controllers

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

const PW_FILE = "data/logins.json"

type User struct {
	Id           int32     `json:"Id"`
	Username     string    `json:"Username"`
	Password     string    `json:"Password"`
	PasswordHash string    `json:"PasswordHash"`
	IsBanned     bool      `json:"IsBanned"`
	LastLogon    time.Time `json:"LastLogon"`
	PwLastSet    string    `json:"PwLastSet"`
	AuthToken    string    `json:"AuthToken"`
	AuthExpires  time.Time `json:"AuthExpires"`
}

type Users struct {
	Users []User `json:"Users"`
}

var UsersList Users

func LoadUsers() {
	bytes, err := os.ReadFile(PW_FILE)
	if err != nil {
		log.Printf("ERROR: Could not load users - %s\n", err)
		return
	}
	if err := json.Unmarshal(bytes, &UsersList); err != nil {
		log.Printf("ERROR: Could not decode json - %s", err)
	}
}
func DumpUsers() {
	jsonData, err := json.MarshalIndent(UsersList, "", "    ")
	if err != nil {
		log.Printf("CRITICAL ERROR: Unable to dump users from memory - %s", err)
		return
	}

	if err := os.WriteFile(PW_FILE, jsonData, 0644); err != nil {
		log.Printf("CRITICAL ERROR: Unable to dump users from memory - %s", err)
		return
	}
}
func retrieveUser(u string, p string) int {
	for i, user := range UsersList.Users {
		if u == user.Username && CheckPasswordHash(p, user.PasswordHash) {
			UsersList.Users[i].LastLogon = time.Now()
			if time.Now().Compare(UsersList.Users[i].AuthExpires.AddDate(0, 0, 1)) > 0 {
				UsersList.Users[i].AuthToken = GenerateToken(user)
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
func validateLogin(u string, p string) User {
	index := retrieveUser(u, p)
	if index >= 0 {
		DumpUsers()
		return UsersList.Users[index]
	}
	return User{Id: -1}
}

func Login(u string, p string) User {
	return validateLogin(u, p)
}
