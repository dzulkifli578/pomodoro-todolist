package operations

import (
	"fmt"
	"strings"
	"time"
)

const (
	MAX_TASKS         = 10
	POMODORO_DURATION = 25
	SHORT_BREAK       = 5
	LONG_BREAK        = 15
)

type Task struct {
	id        uint
	name      string
	completed bool
}

var tasks []Task

func CreateTask() {
	if len(tasks) >= MAX_TASKS {
		println("Task list is full.")
		return
	}

	var taskName string
	fmt.Print("Enter task name: ")
	_, err := fmt.Scanln(&taskName)
	if err != nil {
		fmt.Printf("Error reading task name: %v", err)
		return
	}

	task := Task{
		id:        uint(len(tasks)),
		name:      taskName,
		completed: false,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}

func ReadTasks() {
	if len(tasks) == 0 {
		println("No tasks available.")
		return
	}

	fmt.Println("Your Tasks:")
	for i, task := range tasks {
		status := "Pending"
		if task.completed {
			status = "Done"
		}
		fmt.Printf("%v. %v [%v]\n", i+1, task.name, status)
	}
}

func ReadTaskById(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid index. Task cannot be deleted.")
		return
	}
}

func DeleteTask(index int) {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid index. Task cannot be deleted.")
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task deleted successfully!")
}

func StartPomodoro() {
	if len(tasks) == 0 {
		println("No tasks available.")
		return
	}

	ReadTasks()
	var index int
	fmt.Print("Choose your task to complete: ")
	fmt.Scanln(&index)

	var cycles int

	for {
		fmt.Printf("\nStarting Pomodoro for %v minutes...\n", POMODORO_DURATION)

		for i := POMODORO_DURATION * 60; i > 0; i-- {
			minutes := i / 60
			seconds := i % 60
			fmt.Printf("\rTime remaining: %02v:%02v", minutes, seconds)
			time.Sleep(time.Second)
		}

		for {
			fmt.Println("\nTime's up! Mark task as completed? (y/n): ")
			var choice string
			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Error reading input: ", err)
				continue
			}

			switch strings.ToLower(choice) {
			case "y":
				tasks[index].completed = true
			case "n":
				tasks[index].completed = false
			default:
				fmt.Println("Input only (y/n)")
				break
			}
		}

		cycles++
		if cycles%4 == 0 {
			fmt.Printf("Taking a long break for %v minutes...\n", LONG_BREAK)
			time.Sleep(LONG_BREAK * time.Second)
		} else {
			fmt.Printf("Taking a short break for %d minutes...\n", SHORT_BREAK)
			time.Sleep(SHORT_BREAK * time.Second)
		}

		fmt.Print("Do you want to continue? (y/n): ")
		var cont string
		_, err := fmt.Scanln(&cont)
		if err != nil || (cont != "y" && cont != "Y") {
			break
		}
	}
}
