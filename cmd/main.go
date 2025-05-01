package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"taskmgr/taskmgr"
)

var taskList = []taskmgr.Task{}

func DisplayTask(task taskmgr.Task) {
	fmt.Println("Task Name: ", task.Name)
	fmt.Println("Task Description: ", task.Desc)
	if task.Note != "" {
		fmt.Println("Task Note: ", task.Note)
	}
	fmt.Println("Task Status: ", taskmgr.StatusStr[task.Status])
	fmt.Println("Task Tags: ", task.Tags)
}

func showHelp() {
	helpstr := `This is a simple task manager for managing tasks.
Each task has a mandatory non-empty NAME and a status (Not Started, In Progress, Deferred, Completed)
Additionally, each task may have an ALIAS, a DESCRIPTION, and a set of TAG(S). 
	
usage: taskmgr <command> [options] [args]

The commands to interact with the tasks are:
  Add a new task:		add <NAME> <DESCRIPTION> <ALIAS>    
  Edit a task:			edit <NAME> | <ALIAS>   	
  Change the status:		update <NAME> | <ALIAS> <STATUS>
  Delete a task:		delete <NAME> | <ALIAS>   
  
The commands to display the tasks are:
  Display a specific task: 	display <NAME> | <ALIAS>
  Display all tasks: 		list
`

	fmt.Println(helpstr)
}

func addTask(reader *bufio.Reader) {
	// Add a new task
	var name string

	// Loop until a valid task name is provided
	for {
		fmt.Print("  Enter task name: ")
		name, _ = reader.ReadString('\n')
		name = name[:len(name)-1] // Remove the newline character
		if name == "" {
			fmt.Println("  Task name cannot be empty.")
		} else {
			break
		}
	}

	fmt.Print("  Enter an optional task description: ")
	desc, _ := reader.ReadString('\n')
	desc = desc[:len(desc)-1] // Remove the newline character

	// Each task has a unique ID
	id := GenerateID()

	// Create the taskmgr.Task object and add it to the task list
	newTask := taskmgr.MakeTask(id, name, desc, "", []int{})
	taskList = append(taskList, newTask)
	fmt.Println("  Task created with ID:", id)
}

func listTask() {
	// List all tasks
	if len(taskList) == 0 {
		fmt.Println("No tasks exist.")
		return
	}

	fmt.Printf("ID\t|\tName\t\t|\tStatus\t\t|\tTags\n")
	fmt.Println("----------------------------------------------------")
	for _, task := range taskList {
		fmt.Printf("%d\t|\t%s\t\t|\t%s\t\t|\n", task.ID, task.Name, taskmgr.StatusStr[task.Status])
	}
	fmt.Println("----------------------------------------------------")
}

func GenerateID() taskmgr.TaskID {
	// Generate a unique ID for the task
	return taskmgr.TaskID(len(taskList))
}

func mainLoop() {
	// Main loop for the task manager
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("taskmgr> ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character

		switch input {
		case "add", "a":
			addTask(reader)
		case "edit":
			// Edit a task
		case "update":
			// Update a task
		case "delete":
			// Delete a task
		case "display":
			// Display a specific task
		case "list", "l":
			listTask()
		case "help":
			showHelp()
		case "exit", "quit", "q":
			fmt.Println("Exiting task manager.")
			return
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}

func main() {
	// Set up logging
	log.SetOutput(os.Stdout)

	if len(os.Args) < 2 {
		// No arguments provided, start the main loop
		fmt.Println("Initializing the Go Task Manager")
		fmt.Println("Type 'help' for a list of commands.")
		mainLoop()
	}

	// Called without arguments
	if len(os.Args) == 2 {
		if os.Args[1] != "help" {
			fmt.Println("Unknown command:", os.Args[1])
		}
		showHelp()
		return
	}

	// switch os.Args[1] {
	// case "help":
	// 	showHelp()
	// case "add":
	// 	// Add a new task
	// case "edit":
	// 	// Edit a task
	// case "update":
	// 	// Update a task
	// case "delete":
	// 	// Delete a task
	// case "display":
	// 	// Display a specific task
	// case "list":
	// 	// Display all tasks
	// default:
	// 	fmt.Println("Unknown command:", os.Args[1])
	// 	showHelp()
	// 	return
	// }

	// bufio.NewReader(os.Stdin)

	// fmt.Println("Welcome")
	// // Init()
	// fmt.Println("Creating a new task...")
	// // name, _ := ReadStr("Enter task name: ")
	// // desc, _ := ReadStr("Enter task description: ")
	// // note, _ := ReadStr("Enter additional notes: ")

	// name := "Test Task"
	// desc := "This is a test task"
	// note := "This is a test note"

	// var taskList []tasks.Task
	// newtask := tasks.MakeTask(name, desc, note, []int{})
	// taskList = append(taskList, newtask)
	// for task := range taskList {
	// 	fmt.Println()
	// 	DisplayTask(taskList[task])
	// }
}
