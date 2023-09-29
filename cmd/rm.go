package cmd

import (
	"fmt"
	"strconv"

	"github.com/Ed-cred/tasker/db"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a task without marking it as completed.",
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
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to remove task \"%d\". Error: %s\n", id, err)
			} else {
				fmt.Printf("Deleted task \"%d\".\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
