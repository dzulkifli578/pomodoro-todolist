package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"pomodoro-cli/internal/models"
)

const taskFile = "tasks.json"

func LoadTasksFromFile() []models.Task {
	var tasks []models.Task
	file, err := os.Open(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks
		}
		fmt.Println("Error reading tasks file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding task:", err)
	}
	return tasks
}

func SaveTasksToFile(tasks []models.Task) {
	file, err := os.Create(taskFile)
	if err != nil {
		fmt.Println("Error creating tasks file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding task:", err)
	}
}
