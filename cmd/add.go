package cmd

import (
	"fmt"
	"strings"

	"github.com/Ed-cred/tasker/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Unable to create task:", err)
			return 
		}
		fmt.Printf("Added \"%s\" task to the list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
