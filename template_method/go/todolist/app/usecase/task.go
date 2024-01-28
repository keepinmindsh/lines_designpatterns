package usecase

import (
	"fmt"
	"todolist/domain"
)

type TaskTodo struct {
}

func (t TaskTodo) MakeContent() {
	fmt.Println("오늘의 할일을 간단하게 작성합니다.")
}

func NewTaskToDo() domain.CustomToDo {
	return &TaskTodo{}
}
