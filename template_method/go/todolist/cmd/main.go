package main

import (
	"todolist/app/todo"
	"todolist/domain"
)

func main() {
	healthTodo := todo.MustNewTodo(domain.HealthToDo)

	healthTodo.MakeTitle()
	healthTodo.MakeContent()
	healthTodo.MakePeriod()
	healthTodo.MakeConfirm()

	taskTodo := todo.MustNewTodo(domain.TaskTodo)

	taskTodo.MakeTitle()
	taskTodo.MakeContent()
	taskTodo.MakePeriod()
	taskTodo.MakeConfirm()
}
