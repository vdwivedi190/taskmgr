package taskmgr

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

func MakeTask(id TaskID, name, desc, note string, tags []int) Task {
	// // Generate a unique ID for the task
	// id := GenerateID()

	return Task{
		Name:   name,
		Desc:   desc,
		Note:   note,
		Status: NotStarted,
		ID:     id,
		Tags:   tags,
	}
}
