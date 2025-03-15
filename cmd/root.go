package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomodoro-cli",
	Short: "A CLI for Pomodoro task management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome To Pomodoro CLI!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
