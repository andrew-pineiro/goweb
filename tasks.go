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
func getAllTasks() []Task {
	var tasks []Task
	lines, err := readTaskLines(TaskFilePath)
	if err != nil {
		log.Print("ERROR: cannot read lines in file", TaskFilePath)
	}
	for _, line := range lines[1:] {
		tasks = append(tasks, convertLinetoTask(line))
	}
	return tasks
}
