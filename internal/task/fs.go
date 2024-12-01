package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func taskFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
	}

	return path.Join(cwd, "tasks.json")
}

func ReadTasksFromFile() ([]Task, error) {
	filePath := taskFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file: ", err)
		}
		defer file.Close()
		return []Task{}, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding file: ", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath := taskFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding file:", err)
		return err
	}

	return nil
}
