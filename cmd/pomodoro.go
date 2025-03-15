package cmd

import (
	"pomodoro-cli/pkg/timer"

	"github.com/spf13/cobra"
)

var pomodoroCmd = &cobra.Command{
	Use:   "pomodoro",
	Short: "Start Pomodoro timer",
	Run: func(cmd *cobra.Command, args []string) {
		timer.StartPomodoro()
	},
}

func init() {
	rootCmd.AddCommand(pomodoroCmd)
}
