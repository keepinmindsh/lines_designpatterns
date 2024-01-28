package todo

import (
	"todolist/app/usecase"
	"todolist/domain"
)

type ToDo struct {
	domain.CommonToDo
	domain.CustomToDo
}

func MustNewTodo(doType domain.ToDoType) ToDo {
	switch doType {
	case domain.TaskTodo:
		return ToDo{
			usecase.NewCommonToDo(),
			usecase.NewTaskToDo(),
		}
	case domain.HealthToDo:
		return ToDo{
			usecase.NewCommonToDo(),
			usecase.NewHealthTodo(),
		}
	default:
		panic("error")
	}
}
