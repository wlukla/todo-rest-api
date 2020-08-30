package todo

import "testing"

func TestAdd(t *testing.T) {
	todos := New()
	todo1 := Todo{"Hello", "World"}
	todos.Add(todo1)

	if len(todos.Todos) != 1 {
		t.Errorf("Todo not added")
	}
}

func TestGetAll(t *testing.T) {
	todos := New()
	todo1 := Todo{"Hello", "World"}
	todos.Add(todo1)

	results := todos.GetAll()

	if len(results) != 1 {
		t.Errorf("Todos not present")
	}
}
