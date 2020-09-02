package store

import "todo-rest-api/internal/model"

// TodoRepository ...
type TodoRepository interface {
	Create(t *model.Todo) error
	GetAll() []model.Todo
}
