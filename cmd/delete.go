package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"task-tracker/internal"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a task",
		Long: `Delete a task by provide an id of task

Example: 	
task-tracker delete 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}

	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return errors.New("the task ID is required")
	}

	return internal.DeleteTask(args[0])
}
