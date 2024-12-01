package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task-tracker/internal"
)

func NewMarkInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-in-progress",
		Short: "Mark status of a task to in-progress",
		Long: `Mark status of a task to in-progress by provide an id of task

Example:
task-tracker mark-in-progress 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, internal.TASK_STATUS_IN_PROGRESS)
		},
	}

	return cmd
}

func NewMarkDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mark-done",
		Short: "Mark status of a task to done",
		Long: `Mark status of a task to done by provide an id of task

Example:
task-tracker mark-done 1`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskStatusCmd(args, internal.TASK_STATUS_DONE)
		},
	}

	return cmd
}

func RunUpdateTaskStatusCmd(args []string, status internal.TaskStatus) error {
	if len(args) == 0 {
		fmt.Println("you must provide id of a task")
	}

	return internal.UpdateStatus(args[0], status)
}
