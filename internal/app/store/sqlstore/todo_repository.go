package sqlstore

import "todo-rest-api/internal/app/model"

// TodoRepository ...
type TodoRepository struct {
	Store *Store
}

// Create ...
func (r *TodoRepository) Create(t *model.Todo) error {
	if err := r.Store.db.QueryRow(
		"INSERT INTO todo (title, description, done) values ($1, $2, $3) RETURNING id",
		t.Title, t.Description, t.Done,
	).Scan(&t.ID); err != nil {
		return err
	}

	return nil
}

// GetAll ...
func (r *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	rows, err := r.Store.db.Query(
		"SELECT title, description, done FROM todo",
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := model.Todo{}

		if err := rows.Scan(&todo.Title, &todo.Description, &todo.Done); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
