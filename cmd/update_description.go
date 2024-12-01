package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task-tracker/internal"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update task description",
		Long: `Update task description by providing id and new description

Example:
task-tracker update 1 new-description`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}

	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) != 2 {
		fmt.Println("you must provide id and new description of a task")
	}

	return internal.UpdateDescription(args[0], args[1])
}
