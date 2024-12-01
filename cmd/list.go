package cmd

import (
	"github.com/spf13/cobra"
	"task-tracker/internal"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long: `List all tasks. You can filter tasks by status

Example:
task-tracker list
task-tracker list to-do
task-tracker list in-progress
task-tracker list done`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListCmd(args)
		},
	}

	return cmd
}

func RunListCmd(args []string) error {
	statusFilter := "all"
	if len(args) == 1 {
		statusFilter = args[0]
	}

	return internal.ListTasks(internal.TaskStatus(statusFilter))
}
