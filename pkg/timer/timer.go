package timer

import (
	"bufio"
	"fmt"
	"os"
	"pomodoro-cli/internal/task"
	"strings"
	"time"
)

const (
	POMODORO_DURATION = 25 * time.Minute
	SHORT_BREAK       = 5 * time.Minute
	LONG_BREAK        = 15 * time.Minute
)

func StartPomodoro() {
	task.ReadTasks()

	var index int
	fmt.Print("Choose your task to complete: ")
	_, err := fmt.Scanln(&index)
	if err != nil || index < 0 || index >= len(task.Tasks) {
		fmt.Println("Invalid task choice.")
		return
	}

	cancelCh := make(chan struct{})

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err == nil {
				if strings.TrimSpace(strings.ToLower(input)) == "cancel" {
					cancelCh <- struct{}{}
					return
				}
			}
		}
	}()

	for {
		fmt.Printf("\nStarting Pomodoro for %v...\n", POMODORO_DURATION)
		remaining := int(POMODORO_DURATION.Seconds())

		for remaining > 0 {
			minutes := remaining / 60
			seconds := remaining % 60
			fmt.Printf("\rTime remaining: %02d:%02d  (type 'cancel' to abort)", minutes, seconds)
			select {
			case <-time.After(1 * time.Second):
				remaining--
			case <-cancelCh:
				fmt.Println("\nPomodoro cancelled.")
				return
			}
		}

		var choice string
		fmt.Println("\nTime's up! Mark task as completed? (y/n): ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		if strings.ToLower(choice) == "y" {
			task.LoadTasks()[index].Completed = true
		}

		var breakTime time.Duration
		if index%4 == 0 {
			breakTime = LONG_BREAK
		} else {
			breakTime = SHORT_BREAK
		}

		fmt.Printf("Taking a break for %v...\n", breakTime)
		time.Sleep(breakTime)

		var cont string
		fmt.Print("Do you want to continue? (y/n): ")
		_, err = fmt.Scanln(&cont)
		if err != nil || strings.ToLower(cont) != "y" {
			break
		}
	}
}
