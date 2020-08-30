package main

import (
	"fmt"
	"todo-rest-api/platform/todo"
)

func main() {
	todos := todo.New()

	fmt.Println(todos)

	todos.Add(todo.Todo{"Hello", "How are you?"})

	fmt.Println(todos)
}
