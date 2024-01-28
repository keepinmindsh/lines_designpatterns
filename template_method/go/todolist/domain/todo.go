package domain

type ToDoType string

const (
	HealthToDo ToDoType = "Health"
	TaskTodo   ToDoType = "Task"
)

type ToDo interface {
	CustomToDo
	CommonToDo
}

type CustomToDo interface {
	MakeContent()
}

type CommonToDo interface {
	MakeTitle()
	MakePeriod()
	MakeConfirm()
}
