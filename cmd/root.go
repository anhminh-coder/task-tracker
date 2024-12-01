package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
		Long: `Task Tracker is a CLI tool for managing tasks. It allows you to create, list, and delete tasks.
				You can also mark tasks as completed and update their status.`,
	}
	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewMarkInProgressCmd())
	cmd.AddCommand(NewMarkDoneCmd())

	return cmd
}
