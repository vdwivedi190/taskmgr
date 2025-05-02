# A CLI-based task manager using Go

taskmgr is a lightweight task manager written in Go which provides a command line based interface to manage tasks. The tasks are described by a title, a description, a status (Not Started, In Progress, Completed, or Deferred). The tasks are stored in a JSON format in the user's home directory. 

## Usage 

The module can be invoked either by first building it using `go build` or directly invoked from the `taskmgr` directory using `go run` as 
```
vatsal@taskmgr>go run . 
Welcome to the Go Task Manager!
Imported 6 previous tasks.

Type 'help' for a list of commands.
taskmgr>
```
The `taskmgr` prompt supports a set of commands to interact with the stored tasks. 
```
taskmgr> help
This is a simple CLI-based task manager. Each task has a mandatory
non-empty NAME and a STATUS (Not Started, In Progress, Completed,
or Deferred). Additionally, each task may have a DESCRIPTION and a
set of NOTES. The task manager prompt supports the following commands:

  a, add:       Add a new task
  e, edit:      Edit an existing task
  s, stat:      Change the status of a task
  f, fin:       Change the status of a task
  r, rm:        Delete a task
  d. disp:      Display a specific task
  l, list:      List all existing tasks
  h, help:      Show this help message
  q, quit:      Exit the task manager  

taskmgr> 
```
Before the program exits, the tasks are saved to `~/taskmgr/tasks.json`. 
