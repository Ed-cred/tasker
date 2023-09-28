package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("task has been added")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
