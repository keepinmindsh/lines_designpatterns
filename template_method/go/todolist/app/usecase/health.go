package usecase

import (
	"fmt"
	"todolist/domain"
)

type HealthTodo struct {
}

func (h HealthTodo) MakeContent() {
	fmt.Println("운동을 위한 상세 운동 방법을 입력합니다.")
}

func NewHealthTodo() domain.CustomToDo {
	return &HealthTodo{}
}
