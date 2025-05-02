package taskmgr

type TaskStatus int
type TaskID int

const (
	// Task Status Constants
	NotStarted = iota
	InProgress
	Done
	Deferred
)

var StatusStr = map[TaskStatus]string{
	NotStarted: "Not Started",
	InProgress: "In Progress",
	Done:       "Done",
	Deferred:   "Deferred",
}

// Tags not implemented yet
type Task struct {
	Name   string
	Desc   string
	Notes  []string
	ID     TaskID
	Status TaskStatus
	Tags   []int
}

func MakeTask(id TaskID, name, desc string, notes []string, tags []int) Task {
	return Task{
		Name:   name,
		Desc:   desc,
		Notes:  notes,
		Status: NotStarted,
		ID:     id,
		Tags:   tags,
	}
}
