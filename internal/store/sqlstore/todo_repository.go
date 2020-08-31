package sqlstore

import "todo-rest-api/internal/model"

// TodoRepository ...
type TodoRepository struct {
	store *Store
}

// Create ...
func (r *TodoRepository) Create(t *model.Todo) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO todo (title, description, done) values ($1, $2, $3) RETURNING id",
		t.Title, t.Description, t.Done,
	).Scan(&t.ID); err != nil {
		return err
	}

	return nil
}

// Get all
func (r *TodoRepository) GetAll() []model.Todo, error {
	t := &model.Todo{}
	
	rows, err := r.store.db.Query(
		"SELECT title, description, done FROM todo",
	)

	if err != nil {
		return nil, err
	}

	return t
}