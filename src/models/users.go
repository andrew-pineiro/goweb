package models

import "time"

type User struct {
	Id             int32
	Guid           string
	Username       string
	Email          string
	Password       string
	LastLoginDate  time.Time
	LastChangeDate time.Time
}

type Users struct {
	Users []User `json:"Users"`
}
