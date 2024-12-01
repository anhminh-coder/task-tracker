package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"task-tracker/internal"
)

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		Long: `Add a task to the task list by provide a description
Example:
task-tracker add description`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}

	return addCmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("you must provide a task description")
	}

	return internal.AddTask(args[0])
}
