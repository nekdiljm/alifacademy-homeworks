package task

type Task struct {
	Title string
	Done  bool
}

func NewTask(title string) Task {
	return Task{
		Title: title,
	}
}
