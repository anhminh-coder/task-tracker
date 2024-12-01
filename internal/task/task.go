package internal

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        = "todo"
	TASK_STATUS_IN_PROGRESS = "in-progress"
	TASK_STATUS_DONE        = "done"
)

type Task struct {
	ID          string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(description string) *Task {
	return &Task{
		ID:          ulid.Make().String(),
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	var filteredTasks []Task
	if status != "all" {
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	} else {
		filteredTasks = tasks
	}

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Status", "Created_At", "Updated_At"})
	for _, task := range filteredTasks {
		table.Append([]string{
			task.ID,
			task.Description,
			string(task.Status),
			task.CreatedAt.String(),
			task.UpdatedAt.String(),
		})
	}

	table.Render()
	return nil
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	newTask := NewTask(description)
	tasks = append(tasks, *newTask)

	fmt.Println("Task added successfully with id: ", newTask.ID)
	return WriteTasksToFile(tasks)
}

func DeleteTask(id string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}
	if len(tasks) == len(updatedTasks) {
		fmt.Printf("Task %s not found", id)
	}
	fmt.Printf("Task %s deleted successfully", id)
	return WriteTasksToFile(updatedTasks)
}

func UpdateDescription(id, description string) error {
	tasks, err := ReadTasksFromFile()
	found := false
	if err != nil {
		return err
	}
	for index, _ := range tasks {
		if tasks[index].ID == id {
			found = true
			tasks[index].Description = description
			tasks[index].UpdatedAt = time.Now()
			break
		}
	}
	if !found {
		fmt.Printf("Task %s not found.", id)
		return nil
	}
	fmt.Printf("Task %s updated successfully", id)
	return WriteTasksToFile(tasks)
}

func UpdateStatus(id string, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	found := false
	if err != nil {
		return err
	}
	for index, _ := range tasks {
		if tasks[index].ID == id {
			found = true
			tasks[index].Status = status
			tasks[index].UpdatedAt = time.Now()
			break
		}
	}
	if !found {
		fmt.Printf("Task %s not found.", id)
		return nil
	}
	fmt.Printf("Task %s is marked to %s", id, status)
	return WriteTasksToFile(tasks)
}
