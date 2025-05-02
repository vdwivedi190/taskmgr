package tmp

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TaskStatus int
type TaskID int

const (
	// Task Status Constants
	NotStarted = iota
	InProgress
	Deferred
	Completed
)

var StatusStr = map[TaskStatus]string{
	NotStarted: "Not Started",
	InProgress: "In Progress",
	Deferred:   "Deferred",
	Completed:  "Completed",
}

type Task struct {
	Name   string
	Desc   string
	Note   string
	ID     TaskID
	Status TaskStatus
	Tags   []int
}

func GenerateID() TaskID {
	// Generate a unique ID for the task
	return 1
}

func makeTask(name, desc, note string, tags []int) Task {
	// Generate a unique ID for the task
	id := GenerateID()

	return Task{
		Name:   name,
		Desc:   desc,
		Note:   note,
		Status: NotStarted,
		ID:     id,
		Tags:   tags,
	}
}

// func InitDir(notedir string) {
// 	dirErr := os.Mkdir(notedir, 0666)
// 	if dirErr != nil {
// 		log.Fatal("Error creating directory: ", dirErr)
// 	}

// 	// _, fileErr := os.Create(notedir + "/tasks.json")

// 	// if err != nil {
// 	// 	log.Fatal("Error creating directory: ", err)
// 	// }
// 	// fmt.Println(createErr)
// 	// if createErr != nil {
// 	// 	log.Fatal("Error creating file: ", createErr)
// 	// }
// 	// fmt.Println("File created successfully")
// }

func Init() {

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory. Aborting... ", err)
	}

	notedir := homedir + "/.gotaskmgr"
	taskfile := notedir + "/tasks.json"
	tagfile := notedir + "/tags.json"

	_, statErr := os.Stat(notedir)
	if os.IsNotExist(statErr) {
		// Create the directory if it doesn't exist
		log.Default().Println("Directory does not exist. Creating...")
		err := os.MkdirAll(notedir, 0666)
		if err != nil {
			log.Fatal("Error creating directory: ", err)
		}
	}

	// Create a JSON file for tasks and tags if they don't exist
	_, statErr = os.Stat(taskfile)
	if os.IsNotExist(statErr) {
		_, fileErr := os.Create(taskfile)
		if fileErr != nil {
			log.Fatal("Error creating tasks file: ", fileErr)
		}
		log.Default().Println("Tasks file initialized successfully")
	}

	_, statErr = os.Stat(tagfile)
	if os.IsNotExist(statErr) {
		_, fileErr := os.Create(tagfile)
		if fileErr != nil {
			log.Fatal("Error creating tasks file: ", fileErr)
		}
		log.Default().Println("Tasks file initialized successfully")
	}

	// fmt.Println("Task Manager Initialized")
}

func DisplayTask(task Task) {
	fmt.Println("Task Name: ", task.Name)
	fmt.Println("Task Description: ", task.Desc)
	if task.Note != "" {
		fmt.Println("Task Note: ", task.Note)
	}
	fmt.Println("Task Status: ", StatusStr[task.Status])
	fmt.Println("Task Tags: ", task.Tags)
}

func ReadStr(prompt string) (string, error) {
	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	str, err := reader.ReadString('\n')
	str = str[:len(str)-1] // Remove the newline character
	// Call the reader to read user's input
	return str, err
}

func main() {
	fmt.Println("Welcome")
	Init()
	fmt.Println("Creating a new task...")
	name, _ := ReadStr("Enter task name: ")
	// desc, _ := ReadStr("Enter task description: ")
	// note, _ := ReadStr("Enter additional notes: ")
	desc := "This is a test task"
	note := "This is a test note"

	var taskList []Task
	newtask := makeTask(name, desc, note, []int{})
	taskList = append(taskList, newtask)
	for task := range taskList {
		fmt.Println()
		DisplayTask(taskList[task])
	}
}

// usage: taskmgr <command> [options] [args]

// The commands to interact with the tasks are:
//   Add a new task:		add <NAME> <DESCRIPTION> <ALIAS>
//   Edit a task:			edit <NAME> | <ALIAS>
//   Change the status:		update <NAME> | <ALIAS> <STATUS>
//   Delete a task:		delete <NAME> | <ALIAS>

// The commands to display the tasks are:
//   Display a specific task: 	display <NAME> | <ALIAS>
//   Display all tasks: 		list
// `
