package controllers

import (
	"log"
	"net/http"
)

type Message struct {
	Name    string
	Email   string
	Message string
}

func RecvMessage(msg Message, r *http.Request) error {
	var err error
	log.Println(r.RemoteAddr, "RECIEVED:", msg)
	return err
}
