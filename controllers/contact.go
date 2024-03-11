package controllers

import "log"

type Message struct {
	Name  string
	Email string
	Body  string
}

func RecvMessage(msg Message) error {
	var err error
	log.Println(msg)
	return err
}
