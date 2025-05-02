package main

import (
	"fmt"
	"log"

	"taskmgr/cli"
	"taskmgr/storage"
)

func main() {
	fmt.Println("Welcome to the Go Task Manager!")

	// Get the task list from the storage package
	taskList, err := storage.GetTaskList()
	if err != nil {
		log.Fatal("Error getting task list: ", err)
	}

	if len(taskList) == 0 {
		fmt.Println("No tasks found. Initializing an empty task list.")
	} else {
		fmt.Println("Imported ", len(taskList), "previously created tasks.")
	}

	// Drop to the taskmgr prompt
	// The function CLILoop() returns the final task list
	fmt.Printf("\nType 'help' for a list of commands.\n")
	taskList = cli.CLILoop(taskList)

	// Save the tasks
	err = storage.SaveTaskList(taskList)
	if err != nil {
		log.Fatal("Error saving task list: ", err)
	}
	fmt.Println("Tasks saved successfully.")

	// Called with a single argument
	// if len(os.Args) == 2 {
	// 	...
	// }

}
