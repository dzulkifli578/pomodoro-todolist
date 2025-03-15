package task

import (
	"fmt"
	"pomodoro-cli/internal/models"
	"pomodoro-cli/internal/storage"
)

var Tasks []models.Task

func LoadTasks() []models.Task {
	Tasks = storage.LoadTasksFromFile()
	return Tasks
}

func CreateTask() {
	LoadTasks()
	var taskName string

	fmt.Print("Enter task name: ")
	_, err := fmt.Scanln(&taskName)
	if err != nil {
		fmt.Println("Error reading task name:", err)
		return
	}

	newTask := models.Task{
		ID:        uint(len(Tasks)),
		Name:      taskName,
		Completed: false,
	}

	Tasks = append(Tasks, newTask)
	storage.SaveTasksToFile(Tasks)
	fmt.Println("Task added successfully!")
}

func ReadTasks() {
	LoadTasks()
	if len(Tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	fmt.Println("Your Tasks:")
	for i, task := range Tasks {
		status := "Pending"
		if task.Completed {
			status = "Done"
		}
		fmt.Printf("%v. %v [%v]\n", i, task.Name, status)
	}
}

func DeleteTask() {
	LoadTasks()
	var index int
	fmt.Print("Enter task ID to delete: ")
	fmt.Scanln(&index)

	if index < 0 || index >= len(Tasks) {
		fmt.Println("Invalid index. Task cannot be deleted")
		return
	}

	Tasks = append(Tasks[:index], Tasks[index+1:]...)
	storage.SaveTasksToFile(Tasks)
	fmt.Println("Task deleted successfully!")
}

func MarkTaskAsCompleted(index int) {
	if index < 0 || index >= len(Tasks) {
		fmt.Println("Invalid index. Task cannot be marked as complete")
		return
	}

	Tasks[index].Completed = true
	storage.SaveTasksToFile(Tasks)
}
