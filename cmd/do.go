package cmd

import (
	"fmt"
	"strconv"

	"github.com/Ed-cred/tasker/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("failed to parse argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id - 1]
			err := db.CompleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark task \"%d\" as complete. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked task \"%d\" as complete.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
