package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "task",
	Short: "Tasker is a CLI task manager built in Go",
}