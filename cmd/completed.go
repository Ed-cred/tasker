package cmd

import (
	"fmt"
	"os"

	"github.com/Ed-cred/tasker/db"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows a list of all tasks completed during the current day",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetCompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no completed tasks today. You can do this! ")
			return
		}
		fmt.Println("You have completed the following tasks today:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i + 1, task)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
