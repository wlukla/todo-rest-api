package todo

// Todo ...
type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Repository ...
type Repository struct {
	Todos []Todo
}

// Getter ...
type Getter interface {
	GetAll()
}

// Adder ...
type Adder interface {
	Add()
}

// New ...
func New() *Repository {
	return &Repository{
		Todos: []Todo{},
	}
}

// Add ...
func (r *Repository) Add(todo Todo) {
	r.Todos = append(r.Todos, todo)
}

// GetAll ...
func (r *Repository) GetAll() []Todo {
	return r.Todos
}
