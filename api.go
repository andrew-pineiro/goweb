package main

import "log"

const (
	Token        = "abc123"
	TaskFilePath = "data/tasks.csv"
)

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
func validToken(token string) bool {
	return token == Token
}
