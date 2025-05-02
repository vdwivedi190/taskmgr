package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"taskmgr/taskmgr"
)

var TaskFile, TaskDir string

func initTaskFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Define the directory and file paths in the user's home directory
	TaskDir = homeDir + "/.taskmgr"
	TaskFile = TaskDir + "/tasks.json"

	// Create the directory if it doesn't exist
	_, statErr := os.Stat(TaskDir)
	if os.IsNotExist(statErr) {
		// Create the directory if it doesn't exist
		err := os.MkdirAll(TaskDir, 0777)
		if err != nil {
			return err
		}
	}

	// Create a JSON file for tasks if it doesn't exist
	_, statErr = os.Stat(TaskFile)
	if os.IsNotExist(statErr) {
		_, fileErr := os.Create(TaskFile)
		if fileErr != nil {
			return err
		}

		// Initialize with an empty task list
		taskList := []taskmgr.Task{}
		err = SaveTaskList(taskList)
		if err != nil {
			return err
		}
	}
	return nil
}

// Function to read the task list from the JSON file
func GetTaskList() ([]taskmgr.Task, error) {
	err := initTaskFile()
	if err != nil {
		return nil, err
	}

	// Read the tasks from the JSON file
	file, err := os.Open(TaskFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize the task list
	taskList := []taskmgr.Task{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&taskList)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return []taskmgr.Task{}, err
	}

	return taskList, nil
}

// Function to write the task list to the JSON file
func SaveTaskList(taskList []taskmgr.Task) error {
	file, err := os.Create(TaskFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(taskList)
	if err != nil {
		return err
	}

	return nil
}
