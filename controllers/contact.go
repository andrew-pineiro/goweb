package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	MessageFile = "messages.json"
	DataDir     = "data"
)

type Message struct {
	Name    string
	Email   string
	Message string
	Sender  string
}

func checkExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func loadCurrentMsgs(file string) ([]Message, error) {
	var msgs []Message
	currData, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(currData, &msgs)
	return msgs, nil
}
func RecvMessage(msg Message, r *http.Request) error {
	//TODO(#3): Implement storing in DB
	var msgs []Message
	var err error
	log.Println(r.RemoteAddr, "RECEIVED:", msg)

	msg.Sender = r.RemoteAddr
	if !checkExists(DataDir) {
		os.Mkdir(DataDir, os.ModePerm)
	}

	fileName := filepath.Join(DataDir, MessageFile)
	if checkExists(fileName) {
		msgs, err = loadCurrentMsgs(fileName)
		if err != nil {
			return err
		}
	}

	msgs = append(msgs, msg)

	data, err := json.Marshal(msgs)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
