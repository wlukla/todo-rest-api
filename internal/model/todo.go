package model

// Todo ...
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"-"`
}

// BeforeCreate ...
func (t *Todo) BeforeCreate() {
	t.Done = false
}
