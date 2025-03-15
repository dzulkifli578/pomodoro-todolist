package cmd

import (
	"pomodoro-cli/internal/task"

	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage tasks",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		task.CreateTask()
	},
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		task.ReadTasks()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		task.DeleteTask()
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
	taskCmd.AddCommand(createCmd, readCmd, deleteCmd)
}
