package models

import "time"

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
