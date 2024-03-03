package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Task struct {
	Task        string
	Description string
	Complete    string
}

func convertLinetoTask(line string) Task {
	r := strings.Split(line, ",")
	task := Task{
		Task:        r[0],
		Description: r[1],
		Complete:    r[2],
	}
	return task
}
func readTaskLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Print("ERROR: can't open", &file)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
