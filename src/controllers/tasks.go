package controllers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const TaskFilePath = "data/tasks.csv"

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
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func GetAllTasks() ([]Task, error) {
	var tasks []Task
	lines, err := readTaskLines(TaskFilePath)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return tasks, err
	}
	for _, line := range lines[1:] {
		tasks = append(tasks, convertLinetoTask(line))
	}
	return tasks, err
}
