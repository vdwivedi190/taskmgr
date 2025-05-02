package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"taskmgr/taskmgr"
)

// Maximum string length for formatting the task display
const (
	idLen   = 3
	nameLen = 20
	statLen = 12
	tagLen  = 10
)

// Function to display the help message
func showHelp() {
	helpstr := `This is a simple CLI-based task manager. Each task has a mandatory non-empty NAME and a STATUS (Not Started, In Progress, Completed, or Deferred). Additionally, each task may have an ALIAS, a DESCRIPTION, and a set of TAG(S). 

  a, add: 	Add a new task
  e, edit: 	Edit an existing task
  s, stat: 	Change the status of a task
  f, fin: 	Change the status of a task
  r, rm: 	Delete a task
  d. disp: 	Display a specific task
  l, list: 	List all existing tasks
  h, help: 	Show this help message
  q, quit: 	Exit the task manager  
`
	fmt.Println(helpstr)
}

// Function to generate a unique ID for a new task
func genID(taskList []taskmgr.Task) taskmgr.TaskID {
	// Generate a unique ID for the task
	return taskmgr.TaskID(len(taskList))
}

// Function to obtain a task from the user and add it to the task list
func addTask(taskList []taskmgr.Task, reader *bufio.Reader) []taskmgr.Task {
	var name string

	// Loop until a valid task name is provided
	for {
		fmt.Print("  Enter task name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name[:len(name)-1]) // Remove the newline character
		if name == "" {
			fmt.Println("  Task name cannot be empty.")
		} else {
			break
		}
	}

	fmt.Print("  Enter an optional task description: ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc[:len(desc)-1]) // Remove the newline character

	// Each task has a unique ID
	id := genID(taskList)

	// Create the taskmgr.Task object and add it to the task list
	fmt.Println("  Creating a new task with ID:", id)
	newTask := taskmgr.MakeTask(id, name, desc, "", []int{})
	taskList = append(taskList, newTask)
	return taskList
}

// Search the list to see if the task ID exists
// Returns a boolean indicating if the ID exists and the index of the task
func existsID(taskList []taskmgr.Task, id taskmgr.TaskID) (bool, int) {
	for index, task := range taskList {
		if task.ID == taskmgr.TaskID(id) {
			return true, index
		}
	}
	return false, -1
}

// Function to obtain a valid task ID from the user
// Returns the task ID and the index of the task in the list
// If the user enters an empty string, it returns an error
func getID(taskList []taskmgr.Task, reader *bufio.Reader) (taskmgr.TaskID, int, error) {
	var id taskmgr.TaskID
	var tmpid, index int
	var err error
	var exists bool

	for {
		fmt.Print("  Enter the ID of task to edit (leave empty to quit): ")
		idStr, _ := reader.ReadString('\n')
		idStr = idStr[:len(idStr)-1] // Remove the newline character

		if idStr == "" {
			return 0, -1, errors.New("exit")
		}

		tmpid, err = strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("  Task ID must be an integer!")
		}
		id = taskmgr.TaskID(tmpid)
		exists, index = existsID(taskList, id)
		if !exists {
			fmt.Printf("  A Task with ID %d does not exist!\n", tmpid)
		} else {
			break
		}
	}
	return id, index, nil
}

// Function to obtain a task from the user and add it to the task list
func editTask(taskList []taskmgr.Task, reader *bufio.Reader) {
	id, index, err := getID(taskList, reader)
	if err != nil {
		fmt.Println("  Exiting edit mode...")
		return
	}

	fmt.Println("  Editing task with ID:", id)
	displaySingleTask(taskList[index])

	fmt.Print("  Enter the new name (leave empty to retain the old name): ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1] // Remove the newline character
	if name != "" {
		taskList[index].Name = name
	}

	fmt.Print("  Enter an optional task description (leave empty to retain the old name): ")
	desc, _ := reader.ReadString('\n')
	desc = desc[:len(desc)-1] // Remove the newline character
	if desc != "" {
		taskList[index].Desc = desc
	}

	fmt.Print("  Enter an optional task note (leave empty to retain the old name): ")
	note, _ := reader.ReadString('\n')
	note = note[:len(note)-1] // Remove the newline character
	if note != "" {
		taskList[index].Note = note
	}
}

// Function to obtain a task from the user and add it to the task list
func statTask(taskList []taskmgr.Task, reader *bufio.Reader) {
	id, index, err := getID(taskList, reader)
	if err != nil {
		fmt.Println("  Exiting stat mode...")
		return
	}

	fmt.Printf("  The current status of Task # %d (%s) is \"%s\"\n", id, taskList[id].Name, taskmgr.StatusStr[taskList[id].Status])

	fmt.Print("  Select the new status (0 = Not Started, 1 = In Progress, 2 = Completed, 3 = Deferred, leave empty to retain the old status): ")
	var statStr string
	var stat int
	var convErr error

	for {
		statStr, _ = reader.ReadString('\n')
		statStr = statStr[:len(statStr)-1] // Remove the newline character
		if statStr == "" {
			fmt.Println("  Retaining the old status.")
			return
		}
		stat, convErr = strconv.Atoi(statStr)
		if convErr != nil {
			fmt.Println("  Status must be an integer!")
			continue
		}
		if stat < 0 || stat > 3 {
			fmt.Println("  Status must be between 0 and 3!")
			continue
		} else {
			break
		}
	}
	taskList[id].Status = taskmgr.TaskStatus(stat)
	fmt.Printf("  Status of Task # %d (%s) updated to \"%s\"\n", id, taskList[index].Name, taskmgr.StatusStr[taskList[index].Status])
}

// Function to obtain a task from the user and add it to the task list
func finTask(taskList []taskmgr.Task, reader *bufio.Reader) {
	id, index, err := getID(taskList, reader)
	if err != nil {
		fmt.Println("  Exiting stat mode...")
		return
	}

	taskList[id].Status = taskmgr.Done
	fmt.Printf("  Status of Task # %d (%s) updated to \"%s\"\n", id, taskList[index].Name, taskmgr.StatusStr[taskList[index].Status])
}

// Function to delete a task from the list
func rmTask(taskList []taskmgr.Task, reader *bufio.Reader) []taskmgr.Task {
	id, index, err := getID(taskList, reader)
	if err != nil {
		fmt.Println("  Exiting stat mode...")
		return taskList
	}

	fmt.Println("  Deleting task with ID:", id)
	displaySingleTask(taskList[index])
	return append(taskList[:index], taskList[index+1:]...)
}

// Function to delete a task from the list
func dispTask(taskList []taskmgr.Task, reader *bufio.Reader) {
	_, index, err := getID(taskList, reader)
	if err != nil {
		return
	}
	displaySingleTask(taskList[index])
}

// Function to format a string to a specific length with a given alignment
// The string is truncated (with ellipsis) if it exceeds the maximum length
// The alignment can be given as "left", "right", or "center" or "l", "r", "c"
func FormatStr(str string, maxLen int, align string) string {
	if len(str) > maxLen {
		return str[:(maxLen-3)] + "..."
	} else {
		switch align {
		case "left", "l":
			return fmt.Sprintf("%-*s", maxLen, str)
		case "right", "r":
			return fmt.Sprintf("%*s", maxLen, str)
		case "center", "c":
			// Center the string
			lpad := (maxLen - len(str)) / 2
			tempStr := fmt.Sprintf("%*s", lpad+len(str), str)
			return fmt.Sprintf("%-*s", maxLen, tempStr)
		default:
			// Default to left alignment
			return fmt.Sprintf("%-*s", maxLen, str)
		}
	}
}

// Function to format a task into a string for display
func formatTaskStr(task taskmgr.Task) string {
	// Construct a formatted string for the task
	idStr := FormatStr(strconv.Itoa(int(task.ID)), idLen, "r")
	nameStr := FormatStr(task.Name, nameLen, "l")
	statStr := FormatStr(taskmgr.StatusStr[task.Status], statLen, "l")
	return idStr + "    " + nameStr + "   " + statStr + "   "
}

// Function to display a single task
func displaySingleTask(task taskmgr.Task) {
	fmt.Println("   Name: ", task.Name)
	fmt.Println("   Description: ", task.Desc)
	if task.Note != "" {
		fmt.Println("   Note: ", task.Note)
	}
	fmt.Println("   Status: ", taskmgr.StatusStr[task.Status])
	// fmt.Println("   Tags: ", task.Tags)
}

// Function to list all tasks in the task list
func listTask(taskList []taskmgr.Task) {
	// List all tasks
	if len(taskList) == 0 {
		fmt.Println("No tasks exist.")
		return
	}

	idStr := FormatStr("ID", idLen, "r")
	nameStr := FormatStr("TASK", nameLen, "l")
	statStr := FormatStr("STATUS", statLen, "l")
	fmt.Println(idStr + "    " + nameStr + "   " + statStr + "   ")

	for _, task := range taskList {
		fmt.Println(formatTaskStr(task))
	}
}

func CLILoop(taskList []taskmgr.Task) []taskmgr.Task {
	// Main loop for the task manager
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("taskmgr> ")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character

		switch input {
		case "a", "add":
			taskList = addTask(taskList, reader)
		case "e", "edit":
			editTask(taskList, reader)
		case "s", "stat":
			// Update a task status
			statTask(taskList, reader)
		case "f", "fin":
			// Update a task status
			finTask(taskList, reader)
		case "r", "rm":
			taskList = rmTask(taskList, reader)
		case "d", "disp":
			dispTask(taskList, reader)
		case "l", "list":
			listTask(taskList)
		case "h", "help":
			showHelp()
		case "q", "quit":
			return taskList
		default:
			fmt.Println("Unknown command:", input)
		}
	}
}

// Function to get the task list (for testing purposes)
// This function should be replaced with a call to the storage package
// to retrieve the task list from a JSON file or database
func GetTaskList() ([]taskmgr.Task, error) {

	// Initialize the task list
	taskList := []taskmgr.Task{}

	id := genID(taskList)
	newTask := taskmgr.MakeTask(id, "Task #1", "", "", []int{})
	taskList = append(taskList, newTask)

	id = genID(taskList)
	newTask = taskmgr.MakeTask(id, "Looooooooooooooonger Task", "", "", []int{})
	taskList = append(taskList, newTask)

	return taskList, nil
}
